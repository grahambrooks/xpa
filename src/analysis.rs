use serde::Serialize;

use crate::xpath::{Kw, Parsed, Tok};

/// The metrics for one expression. Field names and order are the JSON output.
#[derive(Clone, Debug, PartialEq, Eq, Serialize)]
pub struct Analysis {
    /// Every token including the end of input. Whitespace and comments are
    /// not tokens.
    pub terminal_count: usize,
    /// Times a general comparison (`= != < <= > >=`) differs from the one
    /// before it. The first one counts.
    pub comparison_change: usize,
    /// Times an `and`/`or` token differs from the one before it. The first one
    /// counts. These are tokens, so an element named `and` counts too.
    pub boolean_operator_change: usize,
    pub function_call_count: usize,
    /// The input, verbatim.
    pub pattern: String,
}

impl Analysis {
    pub(crate) fn new(text: &str, parsed: &Parsed) -> Self {
        let mut comparison_change = 0;
        let mut last_comparison = "";
        let mut boolean_operator_change = 0;
        let mut last_boolean = None;
        for token in &parsed.tokens {
            match token.kind {
                Tok::Eq | Tok::Ne | Tok::Lt | Tok::Le | Tok::Gt | Tok::Ge
                    if text[token.start..token.end] != *last_comparison =>
                {
                    comparison_change += 1;
                    last_comparison = &text[token.start..token.end];
                }
                Tok::Kw(k @ (Kw::And | Kw::Or)) if last_boolean != Some(k) => {
                    boolean_operator_change += 1;
                    last_boolean = Some(k);
                }
                _ => {}
            }
        }
        Analysis {
            terminal_count: parsed.tokens.len(),
            comparison_change,
            boolean_operator_change,
            function_call_count: parsed.function_calls,
            pattern: text.to_string(),
        }
    }

    pub fn cognitive_complexity(&self) -> usize {
        self.comparison_change + self.function_call_count * 2 + self.boolean_operator_change * 2
    }
}
