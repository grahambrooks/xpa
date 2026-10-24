//! A hand-written XPath 3.1 lexer and recognising parser.
//!
//! The grammar is `docs/XPath31.g4`; rule names in the parser match it.
//! The parser builds no tree, because nothing needs one: it validates the
//! expression and counts the constructs the analysis reports.

mod lexer;
mod parser;

use std::fmt;

pub use lexer::{Kw, Tok, Token, tokenize};
pub use parser::{Parsed, parse};

#[derive(Clone, Debug, PartialEq, Eq)]
pub struct SyntaxError {
    pub line: usize,
    pub column: usize,
    pub message: String,
}

impl fmt::Display for SyntaxError {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(
            f,
            "line {} column {}: {}",
            self.line, self.column, self.message
        )
    }
}

impl std::error::Error for SyntaxError {}
