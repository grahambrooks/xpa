# XPath Analyser

![GitHub repo size](https://img.shields.io/github/grahambrooks/xpa/README.md)
<!--
![GitHub contributors]()
![GitHub stars]()
![GitHub forks]()
![Twitter Follow]()
-->
[XPath](https://www.w3.org/2002/11/xquery-xpath-applets/xpath-bnf.html) Analyser (XPA) is a tool that analyses XPath Expressions for people working with large number or complex XML Stylesheets.
XPA reads one or more xpath expressions and calculates the number of tokens, function calls and expression changes in the xpath expression.

The grammar is a slightly modified version from the [ANTLR example Grammars](https://github.com/antlr/grammars-v4/tree/master/xpath/xpath31) modified to support the Go ANTLR lexer/parser generator.

## Prerequisites

To build XPA you will need:
* [bazel](https://bazel.build) yep that's it bazel will take care of downloading the tools or at least telling you what you need to install to build the project.
* `git clone git@github.com:grahambrooks/xpa.git` or `git clone https://github.com/grahambrooks/undata.git` or `gh repo clone grahambrooks/xpa`
Building

```
bazel build //...

```

Running the tests

```
bazel test //...

```

## Installing xpa

To install xpa, follow these steps:

Installing is a little more tricky because bazel builds everything in a sandbax (which is a good thing IMO) but it does make locating and installing the compiled binary a little trickier.

On my macOS machine the compiled binary can be found in `bazel-bin/cmd/xpa/xpa_/xpa` and then copied to a convenient location `/usr/local/bin`

Once installed check out the help

```
xpa --help

```

## Using xpa

You can pass in an Xpath expression on the command line

```
pa --json-output "/bookstore/book[price>35]/price" 

``` 

or pipe to the input. The following outputs json format pipped into jq for a nicer display.

```
echo "/bookstore/book[price>35]/price" | xpa --json-output | jq 
```

output:
```json
                      
{
  "terminal_count": 12,
  "comparison_change": 1,
  "boolean_operator_change": 0,
  "function_call_count": 0,
  "pattern": "/bookstore/book[price>35]/price\n"
}
```

## Roadmap

* Applying the xpath analysis to XSLT XML files to understand the complexity of the xpath expressions in the document.

## Contributing to <project_name>

To contribute to <project_name>, follow these steps:

1. Fork this repository.
2. Create a branch: `git checkout -b <branch_name>`.
3. Make your changes and commit them: `git commit -m '<commit_message>'`
4. Push to the original branch: `git push origin <project_name>/<location>`
5. Create the pull request.

Alternatively see the GitHub documentation on [creating a pull request](https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request).

## Contributors

Thanks to the following people who have contributed to this project:

* [@grahambrooks](https://github.com/grahambrooks) 📖

## Contact

Head over to the discussions page.

## License

This project uses [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)