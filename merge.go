package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/seancfoley/ipaddress-go/ipaddr"
)

func merge(inputFile, outputFile string) error {
	// Read input
	lines, err := readInput(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	// Parse and separate IPv4 and IPv6
	var ipv4Ranges, ipv6Ranges []*ipaddr.IPAddress
	duplicates := make(map[string]bool)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Check for duplicates
		if duplicates[line] {
			continue
		}
		duplicates[line] = true

		// Try to parse as CIDR
		cidr := ipaddr.NewIPAddressString(line)
		if cidr == nil {
			return fmt.Errorf("invalid CIDR format '%s'", line)
		}

		addr := cidr.GetAddress()
		if addr == nil {
			return fmt.Errorf("failed to parse address '%s'", line)
		}

		if addr.IsIPv4() {
			ipv4Ranges = append(ipv4Ranges, addr)
		} else {
			ipv6Ranges = append(ipv6Ranges, addr)
		}
	}

	// Merge ranges
	mergedIPv4 := mergeRanges(ipv4Ranges)
	mergedIPv6 := mergeRanges(ipv6Ranges)

	// Combine and sort
	var merged []string
	for _, addr := range mergedIPv4 {
		merged = append(merged, addr.String())
	}
	for _, addr := range mergedIPv6 {
		merged = append(merged, addr.String())
	}
	sort.Strings(merged)

	// Write output
	return writeOutput(outputFile, merged)
}

// mergeRanges merges overlapping IP address ranges
func mergeRanges(ranges []*ipaddr.IPAddress) []*ipaddr.IPAddress {
	if len(ranges) == 0 {
		return ranges
	}

	if len(ranges) == 1 {
		return ranges
	}

	// Use the library's merge functionality
	// MergeToPrefixBlocks merges the first address with remaining addresses
	first := ranges[0]
	remaining := ranges[1:]

	return first.MergeToPrefixBlocks(remaining...)
}

func readInput(filename string) ([]string, error) {
	var file *os.File
	var err error

	if filename == "" {
		file = os.Stdin
	} else {
		file, err = os.Open(filename)
		if err != nil {
			return nil, err
		}
		defer file.Close()
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func writeOutput(filename string, data []string) error {
	var file *os.File
	var err error

	if filename == "" {
		file = os.Stdout
	} else {
		file, err = os.Create(filename)
		if err != nil {
			return err
		}
		defer file.Close()
	}

	writer := bufio.NewWriter(file)
	for _, line := range data {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}
