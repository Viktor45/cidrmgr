package main

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func TestMergeIPv4Basic(t *testing.T) {
	// Create temp input file
	tmpInput, err := os.CreateTemp("", "cidrmgr-input-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpInput.Name())

	tmpOutput, err := os.CreateTemp("", "cidrmgr-output-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpOutput.Name())

	// Write overlapping IPv4 ranges
	inputData := `192.168.0.0/24
192.168.1.0/24
192.168.0.128/25
10.0.0.0/8
10.1.0.0/16
`
	if _, err := tmpInput.WriteString(inputData); err != nil {
		t.Fatalf("Failed to write input: %v", err)
	}
	tmpInput.Close()

	// Run merge
	if err := merge(tmpInput.Name(), tmpOutput.Name()); err != nil {
		t.Fatalf("merge() failed: %v", err)
	}

	// Read output
	output, err := os.ReadFile(tmpOutput.Name())
	if err != nil {
		t.Fatalf("Failed to read output: %v", err)
	}

	result := strings.TrimSpace(string(output))
	lines := strings.Split(result, "\n")

	// Should have merged some ranges
	if len(lines) >= 5 {
		t.Errorf("Expected ranges to be merged, got %d lines: %v", len(lines), lines)
	}

	// Check that output contains valid CIDR notation
	for _, line := range lines {
		if !strings.Contains(line, "/") {
			t.Errorf("Invalid output format: %s", line)
		}
	}
}

func TestMergeIPv6Basic(t *testing.T) {
	tmpInput, err := os.CreateTemp("", "cidrmgr-ipv6-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpInput.Name())

	tmpOutput, err := os.CreateTemp("", "cidrmgr-output-ipv6-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpOutput.Name())

	// Write overlapping IPv6 ranges
	inputData := `2001:db8::/32
2001:db8:1::/48
2001:db8:2::/48
fc00::/7
`
	if _, err := tmpInput.WriteString(inputData); err != nil {
		t.Fatalf("Failed to write input: %v", err)
	}
	tmpInput.Close()

	if err := merge(tmpInput.Name(), tmpOutput.Name()); err != nil {
		t.Fatalf("merge() failed: %v", err)
	}

	output, err := os.ReadFile(tmpOutput.Name())
	if err != nil {
		t.Fatalf("Failed to read output: %v", err)
	}

	result := strings.TrimSpace(string(output))
	lines := strings.Split(result, "\n")

	if len(lines) == 0 {
		t.Errorf("Expected output, got empty result")
	}

	for _, line := range lines {
		if !strings.Contains(line, ":") {
			t.Errorf("Invalid IPv6 format: %s", line)
		}
	}
}

func TestMergeMixed(t *testing.T) {
	tmpInput, err := os.CreateTemp("", "cidrmgr-mixed-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpInput.Name())

	tmpOutput, err := os.CreateTemp("", "cidrmgr-output-mixed-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpOutput.Name())

	// Mix IPv4 and IPv6
	inputData := `192.168.0.0/24
2001:db8::/32
192.168.1.0/24
2001:db8:1::/48
# This is a comment
10.0.0.0/8

172.16.0.0/12
`
	if _, err := tmpInput.WriteString(inputData); err != nil {
		t.Fatalf("Failed to write input: %v", err)
	}
	tmpInput.Close()

	if err := merge(tmpInput.Name(), tmpOutput.Name()); err != nil {
		t.Fatalf("merge() failed: %v", err)
	}

	output, err := os.ReadFile(tmpOutput.Name())
	if err != nil {
		t.Fatalf("Failed to read output: %v", err)
	}

	result := strings.TrimSpace(string(output))
	if result == "" {
		t.Errorf("Expected output, got empty result")
	}

	// Should have both IPv4 and IPv6
	hasIPv4 := strings.Contains(result, "192") || strings.Contains(result, "10") || strings.Contains(result, "172")
	hasIPv6 := strings.Contains(result, "2001")

	if !hasIPv4 || !hasIPv6 {
		t.Errorf("Expected mixed IPv4/IPv6 output, got:\n%s", result)
	}
}

func TestMergeDuplicates(t *testing.T) {
	tmpInput, err := os.CreateTemp("", "cidrmgr-dup-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpInput.Name())

	tmpOutput, err := os.CreateTemp("", "cidrmgr-output-dup-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpOutput.Name())

	// Duplicate entries
	inputData := `192.168.0.0/24
192.168.0.0/24
192.168.0.0/24
10.0.0.0/8
10.0.0.0/8
`
	if _, err := tmpInput.WriteString(inputData); err != nil {
		t.Fatalf("Failed to write input: %v", err)
	}
	tmpInput.Close()

	if err := merge(tmpInput.Name(), tmpOutput.Name()); err != nil {
		t.Fatalf("merge() failed: %v", err)
	}

	output, err := os.ReadFile(tmpOutput.Name())
	if err != nil {
		t.Fatalf("Failed to read output: %v", err)
	}

	result := strings.TrimSpace(string(output))
	lines := strings.Split(result, "\n")

	if len(lines) != 2 {
		t.Errorf("Expected 2 lines after deduplication, got %d: %v", len(lines), lines)
	}
}

func TestMergeStdin(t *testing.T) {
	tmpOutput, err := os.CreateTemp("", "cidrmgr-stdout-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpOutput.Name())

	// Create a pipe to simulate stdin
	r, w, _ := os.Pipe()
	oldStdin := os.Stdin
	os.Stdin = r

	go func() {
		w.WriteString("192.168.0.0/24\n192.168.1.0/24\n")
		w.Close()
	}()

	if err := merge("", tmpOutput.Name()); err != nil {
		t.Fatalf("merge() failed: %v", err)
	}

	os.Stdin = oldStdin

	output, err := os.ReadFile(tmpOutput.Name())
	if err != nil {
		t.Fatalf("Failed to read output: %v", err)
	}

	result := strings.TrimSpace(string(output))
	if result == "" {
		t.Errorf("Expected output from stdin, got empty result")
	}
}

func TestMergeEmptyFile(t *testing.T) {
	tmpInput, err := os.CreateTemp("", "cidrmgr-empty-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpInput.Name())

	tmpOutput, err := os.CreateTemp("", "cidrmgr-output-empty-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpOutput.Name())

	tmpInput.Close()

	if err := merge(tmpInput.Name(), tmpOutput.Name()); err != nil {
		t.Fatalf("merge() failed: %v", err)
	}

	output, err := os.ReadFile(tmpOutput.Name())
	if err != nil {
		t.Fatalf("Failed to read output: %v", err)
	}

	result := strings.TrimSpace(string(output))
	if result != "" {
		t.Errorf("Expected empty output, got: %s", result)
	}
}

func TestMergeInvalidInput(t *testing.T) {
	tmpInput, err := os.CreateTemp("", "cidrmgr-invalid-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpInput.Name())

	tmpOutput, err := os.CreateTemp("", "cidrmgr-output-invalid-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpOutput.Name())

	// Invalid CIDR format
	inputData := `not-a-cidr-range
invalid
`
	if _, err := tmpInput.WriteString(inputData); err != nil {
		t.Fatalf("Failed to write input: %v", err)
	}
	tmpInput.Close()

	if err := merge(tmpInput.Name(), tmpOutput.Name()); err == nil {
		t.Errorf("Expected error for invalid input, got nil")
	}
}

func BenchmarkMergeLarge(b *testing.B) {
	tmpInput, err := os.CreateTemp("", "cidrmgr-bench-*.txt")
	if err != nil {
		b.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpInput.Name())

	tmpOutput, err := os.CreateTemp("", "cidrmgr-bench-out-*.txt")
	if err != nil {
		b.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpOutput.Name())

	// Generate large input with many ranges
	writer := bufio.NewWriter(tmpInput)
	for i := range 1000 {
		writer.WriteString(strings.Repeat("0", i%3) + strings.Repeat("1", i%3) + ".0.0.0/" + string(rune(8+(i%17))) + "\n")
	}
	writer.Flush()
	tmpInput.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		merge(tmpInput.Name(), tmpOutput.Name())
	}
}

func TestReadInputFile(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "cidrmgr-read-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	content := "192.168.0.0/24\n10.0.0.0/8\n"
	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	lines, err := readInput(tmpFile.Name())
	if err != nil {
		t.Fatalf("readInput() failed: %v", err)
	}

	if len(lines) != 2 {
		t.Errorf("Expected 2 lines, got %d", len(lines))
	}

	if lines[0] != "192.168.0.0/24" {
		t.Errorf("Expected first line to be '192.168.0.0/24', got '%s'", lines[0])
	}
}

func TestWriteOutputFile(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "cidrmgr-write-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	data := []string{"192.168.0.0/24", "10.0.0.0/8"}
	if err := writeOutput(tmpFile.Name(), data); err != nil {
		t.Fatalf("writeOutput() failed: %v", err)
	}

	content, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	expected := "192.168.0.0/24\n10.0.0.0/8\n"
	if string(content) != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, string(content))
	}
}
