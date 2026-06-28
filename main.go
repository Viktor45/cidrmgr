package main

import (
	"flag"
	"fmt"
	"os"
)

const version = "1.0.0"

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "merge":
		mergeCMD()
	case "help", "-h", "--help":
		printUsage()
	case "version", "-v", "--version":
		fmt.Printf("cidrmgr version %s\n", version)
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", cmd)
		printUsage()
		os.Exit(1)
	}
}

func mergeCMD() {
	fs := flag.NewFlagSet("merge", flag.ExitOnError)
	inputFile := fs.String("i", "", "Input file (default: stdin)")
	outputFile := fs.String("o", "", "Output file (default: stdout)")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, `
Usage: cidrmgr merge [options]

Merge overlapping IPv4/IPv6 CIDR ranges from a file or stdin.
Input format: one CIDR range per line (e.g., 192.168.0.0/24, 2001:db8::/32)

Options:
  -i string    Input file (reads from stdin if not specified)
  -o string    Output file (writes to stdout if not specified)

Examples:
  cidrmgr merge -i ranges.txt -o merged.txt
  cat ranges.txt | cidrmgr merge
  cidrmgr merge -i ranges.txt > merged.txt

`)
	}

	if err := fs.Parse(os.Args[2:]); err != nil {
		os.Exit(1)
	}

	if err := merge(*inputFile, *outputFile); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Fprintf(os.Stderr, `cidrmgr - Fast CIDR range management tool

Usage:
  cidrmgr <command> [options]

Commands:
  merge       Merge overlapping CIDR ranges
  help        Show this help message
  version     Show version

Run 'cidrmgr <command> -h' for more information about a command.

`)
}
