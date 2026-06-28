# Project Structure

```
cidrmgr/
├── main.go                      # CLI entry point and command handling
├── merge.go                     # Merge command implementation
├── merge_test.go                # Unit tests (9 test functions + benchmarks)
├── go.mod                       # Go module definition
├── go.sum                       # Go module checksums (generated)
├── Makefile                     # Build and test automation
├── README.md                    # Complete documentation
├── LICENSE                      # MIT License
├── CONTRIBUTING.md              # Contributing guidelines
├── CHANGELOG.md                 # Version history
├── .gitignore                   # Git ignore rules
├── .github/
│   └── workflows/
        ├── release.yml         # GitHub Actions CI/CD workflow
        ├── test.yml            # GitHub Actions CI/CD workflow
│       └── ci.yml              # GitHub Actions CI/CD workflow
└── examples/
    ├── ranges.txt              # Example input file
    └── merged.txt              # Example output file
```

## File Descriptions

### Core Files
- **main.go**: CLI framework, command routing, help text, version handling
- **merge.go**: Core merge logic using seancfoley/ipaddress-go library
- **merge_test.go**: Comprehensive tests covering all functionality

### Configuration
- **go.mod**: Defines module and dependencies
- **Makefile**: Convenient targets for build, test, clean, run

### Documentation
- **README.md**: User guide, installation, usage examples
- **CONTRIBUTING.md**: Guidelines for contributors
- **CHANGELOG.md**: Release history and version information

### CI/CD
- **.github/workflows/ci.yml**: Multi-platform build and test pipeline

### Examples
- **examples/**: Sample input/output files for testing

## How to Build

```bash
# Using Go directly
go build -o cidrmgr

# Using Makefile
make build

# Install globally
go install

# Cross-compile for different platforms
GOOS=linux GOARCH=amd64 go build -o cidrmgr-linux
GOOS=darwin GOARCH=amd64 go build -o cidrmgr-macos
GOOS=windows GOARCH=amd64 go build -o cidrmgr.exe
```

## Dependencies

The project has minimal dependencies:
- **github.com/seancfoley/ipaddress-go**: IP address parsing and merging
- **github.com/seancfoley/bintree**: Internal dependency of ipaddress-go

All dependencies are managed through go.mod and are automatically downloaded with `go mod download`.
