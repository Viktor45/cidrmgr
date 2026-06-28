# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2024-06-28

### Added
- Initial release of cidrmgr
- `merge` command for combining overlapping IPv4/IPv6 CIDR ranges
- Support for reading from files or stdin
- Support for writing to files or stdout
- Duplicate detection and removal
- Comment line support (lines starting with #)
- Empty line handling
- Comprehensive unit tests with >80% coverage
- GitHub Actions CI/CD workflow
- Makefile for common tasks
- Full documentation and examples
- Support for mixed IPv4/IPv6 ranges
- Automatic sorting of output ranges
- Optimized for GitHub Actions and CI/CD pipelines

### Features
- Fast and efficient CIDR range merging
- IPv4 and IPv6 support
- Clean error messages
- Flexible input/output handling
- Perfect for GitHub Actions workflows

[1.0.0]: https://github.com/viktor45/cidrmgr/releases/tag/v1.0.0
