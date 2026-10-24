.PHONY: help build test lint fmt fmt-check check pre-commit install uninstall clean update outdated release release-dry-run

VERSION := $(shell date +%Y.%-m.%-d)
TAG := v$(VERSION)
PREFIX ?= /usr/local
# Honours a target-dir set in ~/.cargo/config.toml, not just ./target.
TARGET_DIR := $(shell cargo metadata --format-version 1 --no-deps | sed -n 's/.*"target_directory":"\([^"]*\)".*/\1/p')

all: help

help:
	@echo "xpa — complexity metrics for XPath expressions"
	@echo ""
	@echo "Usage: make [target]"
	@echo ""
	@echo "Development:"
	@echo "  check            fmt-check + clippy + test (the pre-commit gate)"
	@echo "  pre-commit       Alias for check; run before every commit"
	@echo "  build            Build the release binary"
	@echo "  test             Run all tests"
	@echo "  lint             Run clippy with warnings denied"
	@echo "  fmt              Format the code"
	@echo ""
	@echo "Installation:"
	@echo "  install          Build and install to $(PREFIX)/bin"
	@echo "  uninstall        Remove from $(PREFIX)/bin"
	@echo ""
	@echo "Release:"
	@echo "  release          Tag and push a CalVer release ($(TAG))"
	@echo "  release-dry-run  Show what a release would do"
	@echo ""
	@echo "Maintenance:"
	@echo "  update           Update dependencies"
	@echo "  outdated         Show outdated dependencies"
	@echo "  clean            Remove build artifacts"

build:
	cargo build --release

test:
	cargo test

lint:
	cargo clippy --all-targets -- -D warnings

fmt:
	cargo fmt --all

fmt-check:
	cargo fmt --all -- --check

check: fmt-check lint test
	@echo "All pre-commit checks passed."

pre-commit: check

install: build
	install -d $(PREFIX)/bin
	install -m 755 $(TARGET_DIR)/release/xpa $(PREFIX)/bin/xpa
	@echo "Installed $(PREFIX)/bin/xpa"

uninstall:
	rm -f $(PREFIX)/bin/xpa

update:
	cargo update

outdated:
	cargo outdated --root-deps-only

clean:
	cargo clean

release-dry-run:
	@echo "Would release $(TAG)"
	@git status --porcelain

release: check
	@test -z "$$(git status --porcelain)" || (echo "working tree is dirty" && exit 1)
	sed -i '' 's/^version = ".*"/version = "$(VERSION)"/' Cargo.toml
	cargo check --quiet
	git add Cargo.toml Cargo.lock
	git commit -m "Release $(VERSION)"
	git tag -a $(TAG) -m "Release $(VERSION)"
	git push origin main
	git push origin $(TAG)
