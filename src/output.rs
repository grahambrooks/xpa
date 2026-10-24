use std::io::{self, IsTerminal, Write};

use xpa::Analysis;

const GREEN: &str = "\x1b[32m";
const BOLD_RED: &str = "\x1b[1;31m";
const RESET: &str = "\x1b[0m";

/// How results are written: JSON lines, a coloured table for a terminal, or
/// one plain line per expression when piped.
pub enum Format {
    Json,
    Terminal { colour: bool },
    Plain,
}

impl Format {
    pub fn select(json: bool) -> Self {
        if json {
            Format::Json
        } else if io::stdout().is_terminal() {
            let colour = std::env::var_os("NO_COLOR").is_none_or(|v| v.is_empty());
            Format::Terminal { colour }
        } else {
            Format::Plain
        }
    }

    pub fn banner(&self) {
        match self {
            Format::Json => {}
            Format::Terminal { colour } => {
                eprintln!("{}", paint(*colour, GREEN, "XPath Analyser (xpa)"))
            }
            Format::Plain => eprintln!("XPath Analyser (xpa)"),
        }
    }

    pub fn error(&self, message: &str) {
        match self {
            Format::Terminal { colour } => eprintln!("{}", paint(*colour, BOLD_RED, message)),
            Format::Json | Format::Plain => eprintln!("{message}"),
        }
    }

    pub fn analysis(&self, out: &mut impl Write, a: &Analysis) -> io::Result<()> {
        match self {
            Format::Json => {
                serde_json::to_writer(&mut *out, a)?;
                writeln!(out)
            }
            Format::Terminal { colour } => {
                let rows = [
                    ("complexity:         ", a.cognitive_complexity().to_string()),
                    ("function calls:     ", a.function_call_count.to_string()),
                    (
                        "boolean changes:    ",
                        a.boolean_operator_change.to_string(),
                    ),
                    ("comparison changes: ", a.comparison_change.to_string()),
                    ("tokens:             ", a.terminal_count.to_string()),
                    ("pattern:            ", a.pattern.clone()),
                ];
                for (label, value) in rows {
                    writeln!(out, "{label}{}", paint(*colour, GREEN, &value))?;
                }
                Ok(())
            }
            Format::Plain => writeln!(
                out,
                "complexity: {}, function calls {}, boolean changes: {}, comparison changes {}, tokens {} pattern {}",
                a.cognitive_complexity(),
                a.function_call_count,
                a.boolean_operator_change,
                a.comparison_change,
                a.terminal_count,
                a.pattern
            ),
        }
    }
}

fn paint(colour: bool, code: &str, text: &str) -> String {
    if colour {
        format!("{code}{text}{RESET}")
    } else {
        text.to_string()
    }
}
