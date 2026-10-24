# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

`xpa` reports complexity metrics for XPath 3.1 expressions: tokens, changes of
comparison and boolean operator, and named function calls.
[DESIGN.md](DESIGN.md) defines the metrics, the quirks kept for parity with
the original Go version, the deliberate changes, and what is out of scope.
Read it before changing behaviour, and record new scope decisions there.

## The constraints that shape everything

1. **The golden corpus is the contract.** `tests/golden/*.json` are the Go
   implementation's outputs for `tests/corpus/*.xpath`. Every expression Go
   accepted must give identical metrics. Never regenerate them from the Rust
   build. The Go build is gone, so they can't be regenerated at all.
2. **Pure Rust, no parser generator.** The lexer and parser are written by
   hand against `docs/XPath31.g4`, with rule names matching. No ANTLR (it
   needs Java), and no pest. That was decided on 2026-10-24.
3. **The parser recognises; it doesn't build a tree.** Metrics come from the
   token stream plus a function-call counter in `primaryexpr`. Add a tree
   only when a metric needs one.

## Architecture

```
text ──tokenize──> tokens ──parse──> Parsed{tokens, function_calls} ──> Analysis
```

Things worth knowing before editing the parser:

- **Choice points replay.** Two decisions need more than two tokens of
  lookahead: a lone `/` before a keyword or `*`, and an occurrence indicator
  `*`/`+`. `Parser::choice` records them. A failed parse advances the last
  one and re-parses from the start, depth first, up to `MAX_ATTEMPTS`. Only
  add a choice point where ANTLR's full-context prediction genuinely accepts
  both readings. Each one multiplies the work on invalid input.
- **Keywords are names too.** `eqname` takes every keyword except `to`. In
  operator position a keyword is an operator. In operand position it is a
  name unless the next token makes it syntax (`for $`, `if (`, `map {`,
  `child::`).
- **Lexing follows ANTLR's longest match, with keywords winning ties.**
  Comments follow the rule's non-greedy semantics exactly. See
  `comment_ends`.

## Commands

```bash
make check                 # fmt-check + clippy + test — the pre-commit gate
cargo test --test golden   # parity with the Go implementation
cargo run -- -j 'a and b'  # one expression, JSON
```

**Run `make check` before every commit**, and only commit once it prints
`All pre-commit checks passed.`

## Git and release

- **Trunk-based: commit and push directly to `main`.** PRs are not required.
- **CalVer** (`2026.10.24`). The tag is the version; release-kit v2 stamps
  it into `Cargo.toml` and updates `Formula/xpa.rb`. `version` stays `0.0.0`
  until the first release.
- **`publish = false`.** No crates.io.
- **MSRV tracks latest stable**, pinned in `rust-toolchain.toml`.

## Conventions vault

`~/knowledge-base` holds the portfolio-wide conventions. Start from the task
page: `wiki/when-starting-a-new-project.md`, `wiki/when-cutting-a-release.md`,
`wiki/when-setting-up-git-and-ci.md`, `wiki/when-you-need-to-trust-a-result.md`.
Never edit vault files directly; mutations go through the `claude-obsidian`
skills.
