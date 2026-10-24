# xpa design

`xpa` measures how hard an XPath 3.1 expression is to read. It was a Go tool
with an ANTLR-generated parser built by Bazel. It is now a single Rust crate
with a hand-written lexer and recursive-descent parser. This file records what
the metrics mean, where the port deliberately differs from Go, and what is
out of scope.

## Parity contract

`tests/corpus/*.xpath` were run through the Go build, and
`tests/golden/*.json` are its outputs. `tests/golden.rs` holds the Rust
implementation to them. **For every expression the Go version accepted, the
metrics are identical.** That was also checked by differential fuzzing against
the Go binary: 25,000 generated expressions, about 8,400 of them accepted by
Go, with no differences.

Don't regenerate the golden files from the Rust build. They're a record of the
Go behaviour. A new metric or a behaviour change needs new corpus cases with
hand-checked expectations, and an entry under "Deliberate changes" below.

## Metrics

| Field | Meaning |
|---|---|
| `terminal_count` | Tokens, **including end of input**. Whitespace and `(: comments :)` are not tokens. `/a` is 3. |
| `comparison_change` | Times a general comparison (`= != < <= > >=`) differs from the previous one. Starts from "none", so the first counts. `eq`/`is`/`<<` never count. |
| `boolean_operator_change` | Times an `and`/`or` **token** differs from the previous one. Starts from "none". |
| `function_call_count` | Named function calls: an EQName that isn't a reserved function name, followed by an argument list. Arrow calls (`=> f()`), dynamic calls (`$f()`), named references (`f#1`) and kind tests (`text()`) don't count. |
| `pattern` | The input verbatim, so stdin keeps its trailing newline. |

`complexity = comparison_change + 2 × function_call_count + 2 × boolean_operator_change`.

### Quirks kept for parity

- `boolean_operator_change` counts tokens, not operators. In `/and/or`, the
  element names `and` and `or` count as two changes.
- A lone `/` followed by a keyword or `*` is accepted either way, as ANTLR's
  full-context prediction does: `/ * 5` is `/` times 5. The XPath spec says
  `/*` is always a path there. The parser handles this with replayable choice
  points (`src/xpath/parser.rs`), along with `instance of xs:int + 1`, where
  `+` could be an occurrence indicator.
- `map{a:b}` is a syntax error: `a:b` lexes as one prefixed name. Write
  `map{a : b}`. The spec agrees.
- Comments follow the ANTLR rule exactly. An inner `(:` nests only if that
  reading can close. `(: a : (: b :)` is one comment.

## Deliberate changes from Go

These only turn Go errors into accepted expressions, turn Go acceptances of
malformed input into errors, or change output that isn't a metric.

**Valid XPath the Go grammar couldn't parse, now accepted:**

- `prefix:*` and `*:local` wildcards, and `?name` lookups. The grammar's
  `NCName` token could never win a tie with `QName`.
- `'it''s'` and `"a""b"`: doubled-quote escapes. The ANTLR rule was
  non-greedy and split them in two.
- Braced URIs with content: `Q{http://x}local`. The ANTLR character set
  `[^{}]` matched only `^`, `{` and `}`. That also made Go accept `Q{{}*`,
  which is now an error.
- `processing-instruction(name)` with an unquoted name.

**Errors are reported, not hidden:**

- An unrecognised character (`1 ^+ 2`) or an unterminated string is a syntax
  error. The Go lexer printed a warning, dropped the character and carried on.
- A syntax error prints `failed to analyse <input>: line L column C: …` to
  stderr in every output mode. Go printed a bare `failed to analyse` (or
  nothing in JSON mode). Columns are 1-based.
- **The exit status is non-zero if any input failed.** Go always exited 0,
  which hid failures in pipelines.

**CLI:**

- An argument is read as a file only if it is a regular file. Go `stat`ed it,
  so `.`, `..` and `/` (directories, and valid XPath) failed.
- Negative numbers (`xpa -1`) are expressions, not flags. Any other expression
  that starts with `-` needs `--` first.
- JSON output uses serde's escaping. Go escaped `<`, `>` and `&` as `<`
  and so on. The parsed JSON is identical.
- The terminal banner goes to stderr only. Go also wrote a stray blank line to
  stdout.

## Out of scope

- **XSLT scanning** (pulling expressions out of `.xsl` files) is still a
  roadmap item. It isn't part of the port.
- **New metrics.** The four counts and the complexity formula are the
  contract above.
- **A parse tree.** The parser only recognises: it validates and counts.
  Building an AST is warranted only when a metric needs one.

## Layout

```
src/main.rs          CLI (clap): args, stdin, files, exit status
src/output.rs        JSON lines / terminal table / plain line
src/lib.rs           analyse(text) -> Result<Analysis, SyntaxError>
src/analysis.rs      the metrics, computed from tokens and parse counts
src/xpath/lexer.rs   tokens, mirroring the lexer rules of docs/XPath31.g4
src/xpath/parser.rs  recognising recursive descent over its parser rules
docs/XPath31.g4      the grammar the Go version generated its parser from
```
