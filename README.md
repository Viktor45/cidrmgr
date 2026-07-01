# cidrmgr

Fast and efficient CLI tool for managing IPv4/IPv6 CIDR ranges. Built with Go and optimized for use in GitHub Actions.

- [cidrmgr](#cidrmgr)
  - [Features](#features)
  - [Installation](#installation)
    - [Build from Source](#build-from-source)
    - [Using Make (Optional)](#using-make-optional)
  - [Usage](#usage)
    - [Basic Commands](#basic-commands)
    - [Merge Command](#merge-command)
      - [Usage](#usage-1)
      - [Options](#options)
      - [Examples](#examples)
  - [Input Format](#input-format)
    - [Example Input File](#example-input-file)
    - [Example Output](#example-output)
  - [GitHub Actions Integration](#github-actions-integration)
  - [Performance](#performance)
  - [Testing](#testing)
    - [Test Coverage](#test-coverage)
  - [Dependencies](#dependencies)
  - [Error Handling](#error-handling)
  - [Output Characteristics](#output-characteristics)
  - [Contributing](#contributing)
  - [License](#license)
  - [Acknowledgments](#acknowledgments)
  - [Troubleshooting](#troubleshooting)
    - ["invalid CIDR format" error](#invalid-cidr-format-error)
    - [Large file processing is slow](#large-file-processing-is-slow)
    - [Memory issues](#memory-issues)
  - [Quick Start](#quick-start)


## Features

- ⚡ **Fast**: Efficient CIDR range merging using the `seancfoley/ipaddress-go` library
- 🔄 **Flexible I/O**: Read from files or stdin, write to files or stdout
- 🌐 **IPv4 & IPv6 Support**: Handles both IPv4 and IPv6 CIDR ranges
- 🧹 **Clean Output**: Automatically deduplicates and sorts ranges
- 📝 **Comments Support**: Ignore lines starting with `#`
- 🚀 **CI/CD Ready**: Perfect for GitHub Actions workflows

## Installation

### Build from Source

```bash
git clone https://github.com/viktor45/cidrmgr.git
cd cidrmgr
go build -o cidrmgr
```

### Using Make (Optional)

Create a `Makefile` for easier building:

```makefile
.PHONY: build test clean

build:
	go build -o cidrmgr

test:
	go test -v ./...

test-coverage:
	go test -coverage ./...

clean:
	rm -f cidrmgr
	go clean -testcache
```

Then run:
```bash
make build
make test
```

## Usage

### Basic Commands

```bash
# Show help
cidrmgr help
cidrmgr -h

# Show version
cidrmgr version
cidrmgr -v
```

### Merge Command

The `merge` command combines overlapping IPv4/IPv6 CIDR ranges and removes duplicates.

#### Usage

```bash
cidrmgr merge [options]
```

#### Options

- `-i string` - Input file path (reads from stdin if not specified)
- `-o string` - Output file path (writes to stdout if not specified)

#### Examples

**Merge from file to file:**
```bash
cidrmgr merge -i ranges.txt -o merged.txt
```

**Merge from file to stdout:**
```bash
cidrmgr merge -i ranges.txt
```

**Merge from stdin to file:**
```bash
cat ranges.txt | cidrmgr merge -o merged.txt
```

**Pipe input and output:**
```bash
cat ranges.txt | cidrmgr merge > merged.txt
```

**Default stdin/stdout:**
```bash
cidrmgr merge
# (reads from stdin, writes to stdout)
```

## Input Format

Each line should contain a single CIDR range:
- IPv4: `192.168.0.0/24`
- IPv6: `2001:db8::/32`

Lines starting with `#` are treated as comments and ignored.
Empty lines are also ignored.

### Example Input File

```
# Private IPv4 ranges
192.168.0.0/24
192.168.1.0/24
192.168.0.128/25

# RFC1918 ranges
10.0.0.0/8
172.16.0.0/12

# IPv6 ranges
2001:db8::/32
fc00::/7

# Duplicate (will be removed)
192.168.0.0/24
```

### Example Output

```
10.0.0.0/8
172.16.0.0/12
192.168.0.0/23
2001:db8::/32
fc00::/7
```

## GitHub Actions Integration

Use `cidrmgr` in your GitHub Actions workflows:

```yaml
name: Merge CIDR Ranges

on: [push]

jobs:
  merge:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - name: Setup Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.25'
      
      - name: Build cidrmgr
        run: go build -o cidrmgr
      
      - name: Merge CIDR ranges
        run: ./cidrmgr merge -i allowlist.txt -o merged.txt
      
      - name: Upload merged ranges
        uses: actions/upload-artifact@v3
        with:
          name: merged-ranges
          path: merged.txt
```

## Performance

The tool is optimized for performance and can handle large files efficiently:

- Processes thousands of CIDR ranges quickly
- Memory efficient with streaming input/output
- Perfect for CI/CD pipelines with strict time limits

Run benchmarks:
```bash
go test -bench=. -benchmem ./...
```

## Testing

Run the comprehensive test suite:

```bash
# Run all tests
go test -v ./...

# Run specific test
go test -v -run TestMergeIPv4Basic ./...

# With coverage
go test -v -cover ./...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Test Coverage

The project includes comprehensive tests for:
- ✅ Basic IPv4 range merging
- ✅ IPv6 range merging
- ✅ Mixed IPv4/IPv6 handling
- ✅ Duplicate removal
- ✅ Comment and empty line handling
- ✅ Stdin/stdout operations
- ✅ File I/O
- ✅ Invalid input handling
- ✅ Performance benchmarks

## Dependencies

- Go 1.25 or later
- `github.com/seancfoley/ipaddress-go` - IP address manipulation library

## Error Handling

The tool provides clear error messages for common issues:

```bash
# Invalid CIDR format
$ cidrmgr merge -i bad.txt
Error: invalid CIDR format '192.168.0.256/24': ...

# File not found
$ cidrmgr merge -i nonexistent.txt
Error: failed to read input: open nonexistent.txt: no such file or directory

# Missing input file permissions
$ cidrmgr merge -i restricted.txt
Error: failed to read input: permission denied
```

## Output Characteristics

- **Sorted**: Output is alphabetically sorted
- **Deduplicated**: Duplicate ranges are removed
- **Optimized**: Overlapping ranges are merged into larger blocks
- **One per line**: Each range is on its own line
- **No trailing whitespace**: Clean output format

## Contributing

Contributions are welcome! Please:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

MIT License - see LICENSE file for details

## Acknowledgments

- Built with [seancfoley/ipaddress-go](https://github.com/seancfoley/ipaddress-go) library
- Optimized for GitHub Actions and CI/CD pipelines

## Troubleshooting

### "invalid CIDR format" error

Make sure your CIDR ranges are properly formatted:
- IPv4: `10.0.0.0/8` (not `10.0.0.0/8.0`)
- IPv6: `2001:db8::/32` (not `2001:db8`)

### Large file processing is slow

Verify your system has sufficient RAM and try:
- Breaking the file into smaller chunks
- Using a faster disk (SSD preferred)
- Running on a machine with more CPU cores

### Memory issues

If processing very large files:
- Increase available memory
- Process the file in chunks
- Use a machine with more RAM

## Quick Start

```bash
# 1. Clone and build
git clone https://github.com/viktor45/cidrmgr.git
cd cidrmgr
go build -o cidrmgr

# 2. Create a test file
cat > test.txt << EOF
192.168.0.0/24
192.168.1.0/24
10.0.0.0/8
2001:db8::/32
EOF

# 3. Merge ranges
./cidrmgr merge -i test.txt -o output.txt

# 4. View results
cat output.txt
```

---

**Made with ❤️ for the DevOps community**
 