# xpa — complexity metrics for XPath expressions.
#
# `make check` is the pre-push gate and runs exactly what the CI `check` job runs
# (.github/workflows/ci.yml); `make hooks` wires it into git. Releases come from
# release.mk (release-kit): `make release` tags origin/main and pushes the tag.
PREFIX ?= /usr/local
# Honours a target-dir set in ~/.cargo/config.toml, not just ./target.
TARGET_DIR = $(shell cargo metadata --format-version 1 --no-deps | sed -n 's/.*"target_directory":"\([^"]*\)".*/\1/p')
.DEFAULT_GOAL := help

.PHONY: help build test lint fmt fmt-check check pre-commit hooks install uninstall clean update outdated

help: ## List targets
	@grep -hE '^[a-zA-Z0-9_-]+:.*## ' $(MAKEFILE_LIST) | awk -F':.*## ' '{ printf "  %-16s %s\n", $$1, $$2 }'

check: fmt-check lint test ## The pre-push gate: the CI check job (fmt, clippy, test)
	@echo "All pre-commit checks passed."

pre-commit: check ## Alias for check

build: ## Build the release binary
	cargo build --release

test: ## Tests, as CI runs them
	cargo test --all-features --locked

lint: ## Clippy with warnings denied, as CI runs it
	cargo clippy --all-targets --all-features --locked -- -D warnings

fmt: ## Format the code
	cargo fmt --all

fmt-check: ## Fail if the code is not formatted
	cargo fmt --all --check

hooks: ## Install the prek git hooks (pre-commit and pre-push)
	prek install

install: build ## Build and install to $(PREFIX)/bin
	install -d $(PREFIX)/bin
	install -m 755 $(TARGET_DIR)/release/xpa $(PREFIX)/bin/xpa
	@echo "Installed $(PREFIX)/bin/xpa"

uninstall: ## Remove from $(PREFIX)/bin
	rm -f $(PREFIX)/bin/xpa

update: ## Update dependencies
	cargo update

outdated: ## Show outdated dependencies
	cargo outdated --root-deps-only

clean: ## Remove build artifacts
	cargo clean

include release.mk
