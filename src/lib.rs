//! XPath Analyser: complexity metrics for XPath 3.1 expressions.

mod analysis;
pub mod xpath;

pub use analysis::Analysis;
pub use xpath::SyntaxError;

/// Parses `text` as an XPath 3.1 expression and measures it.
pub fn analyse(text: &str) -> Result<Analysis, SyntaxError> {
    let parsed = xpath::parse(text)?;
    Ok(Analysis::new(text, &parsed))
}
