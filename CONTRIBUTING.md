# Contributing to cidrmgr

Thank you for your interest in contributing to cidrmgr! We welcome contributions of all kinds.

## Getting Started

1. Fork the repository
2. Clone your fork: `git clone https://github.com/viktor45/cidrmgr.git`
3. Create a feature branch: `git checkout -b feature/my-feature`
4. Make your changes
5. Test your changes: `make test`
6. Commit: `git commit -am 'Add feature'`
7. Push: `git push origin feature/my-feature`
8. Submit a Pull Request

## Development Setup

```bash
# Clone the repo
git clone https://github.com/viktor45/cidrmgr.git
cd cidrmgr

# Install dependencies
go mod download

# Build
make build

# Run tests
make test

# Run tests with coverage
make test-coverage
```

## Code Style

- Follow Go's standard code style (use `gofmt`)
- Run `go vet ./...` before submitting
- Add tests for new features
- Keep commits atomic and well-described

## Testing

- Add unit tests for all new features
- Ensure all tests pass: `go test -v ./...`
- Aim for >80% code coverage for new code
- Run benchmarks for performance-critical changes: `go test -bench=.`

## Pull Request Process

1. Update README.md with any new features or changes to usage
2. Update CHANGELOG.md if applicable
3. Ensure all tests pass locally
4. Provide a clear description of your changes
5. Reference any related issues

## Reporting Issues

- Check existing issues before opening a new one
- Provide a clear description of the problem
- Include minimal reproducible example
- Specify your environment (OS, Go version, etc.)

## Feature Requests

- Search existing issues/discussions before requesting
- Explain the use case and benefits
- Provide examples if possible

## Code of Conduct

Be respectful and professional. We're committed to providing a welcoming environment.

## Questions?

Feel free to open a discussion or issue if you have questions about contributing.

Thank you for helping make cidrmgr better!
