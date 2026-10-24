# XPath Analyser

![GitHub repo size](https://img.shields.io/github/repo-size/grahambrooks/xpa) ![Last Updated](https://img.shields.io/github/last-commit/grahambrooks/xpa)

XPath Analyser (`xpa`) measures the complexity of [XPath 3.1](https://www.w3.org/TR/xpath-31/)
expressions, for people working with large numbers of XSLT stylesheets or with
complex ones. It reads one or more expressions and reports the number of
tokens, function calls, and changes of comparison and boolean operator, plus a
complexity score derived from them.

## Installing

```shell
brew trust --formula grahambrooks/xpa/xpa
brew tap grahambrooks/xpa https://github.com/grahambrooks/xpa
brew install grahambrooks/xpa/xpa
```

The formula lives in this repository rather than a separate `homebrew-xpa`
tap, so `brew tap` needs the explicit URL. Homebrew 6 refuses to load
formulae from an untrusted third-party tap, hence `brew trust` first. The
formula appears with the first release.

or from source, with a Rust toolchain:

```shell
git clone https://github.com/grahambrooks/xpa.git
cd xpa
make install            # builds and installs to /usr/local/bin (PREFIX=… to change)
```

Release binaries for macOS and Linux are attached to each
[GitHub release](https://github.com/grahambrooks/xpa/releases).

## Using xpa

Pass expressions as arguments:

```shell
xpa --json-output "/bookstore/book[price>35]/price"
```

pipe one in:

```shell
echo "/bookstore/book[price>35]/price" | xpa --json-output | jq
```

```json
{
  "terminal_count": 12,
  "comparison_change": 1,
  "boolean_operator_change": 0,
  "function_call_count": 0,
  "pattern": "/bookstore/book[price>35]/price\n"
}
```

or name files that each hold one expression. Without `--json-output`, a
terminal gets a table and a pipe gets one line per expression:

```
complexity:         1
function calls:     0
boolean changes:    0
comparison changes: 1
tokens:             12
pattern:            /bookstore/book[price>35]/price
```

`complexity` is comparison changes + 2 × function calls + 2 × boolean
changes. [DESIGN.md](DESIGN.md) defines each metric exactly.

An invalid expression is reported with its line and column, and `xpa` exits
non-zero. Put `--` before an expression that starts with `-` (other than a
negative number).

## Building

```shell
make check    # fmt-check + clippy + tests
make build    # release binary
```

The parser is hand-written Rust following the
[ANTLR XPath 3.1 grammar](https://github.com/antlr/grammars-v4/tree/master/xpath/xpath31)
in [`docs/XPath31.g4`](docs/XPath31.g4). `tests/golden` pins its metrics to
the original Go implementation.

## Roadmap

* Apply the analysis to XSLT files, to show the complexity of the XPath
  expressions in a stylesheet.

## Contact

Head over to the discussions page.

## License

[MIT](LICENSE)
