mod output;

use std::io::{self, Read};
use std::path::Path;
use std::process::ExitCode;

use clap::Parser;

use output::Format;

/// A static analyser for XPath expressions.
#[derive(Parser)]
#[command(name = "xpa", version, about = "XPath analyser")]
struct Cli {
    /// Output the analysis in JSON format, one object per line
    #[arg(short, long)]
    json_output: bool,

    /// XPath expressions, or files that each hold one. Reads stdin when none
    /// are given. Put `--` before an expression that starts with `-`.
    #[arg(allow_negative_numbers = true)]
    inputs: Vec<String>,
}

fn main() -> ExitCode {
    let cli = Cli::parse();
    let format = Format::select(cli.json_output);
    format.banner();

    let mut sources = Vec::new();
    if cli.inputs.is_empty() {
        let mut text = String::new();
        if let Err(e) = io::stdin().read_to_string(&mut text) {
            format.error(&format!("failed to read from stdin: {e}"));
            return ExitCode::FAILURE;
        }
        sources.push(("stdin".to_string(), Ok(text)));
    } else {
        for input in cli.inputs {
            if Path::new(&input).is_file() {
                let text = std::fs::read_to_string(&input);
                sources.push((input, text));
            } else {
                sources.push((format!("`{input}`"), Ok(input)));
            }
        }
    }

    let mut ok = true;
    let mut out = io::stdout().lock();
    for (name, text) in sources {
        let result = match text {
            Ok(text) => xpa::analyse(&text).map_err(|e| e.to_string()),
            Err(e) => Err(format!("could not read file: {e}")),
        };
        match result {
            Ok(analysis) => {
                if let Err(e) = format.analysis(&mut out, &analysis) {
                    if e.kind() == io::ErrorKind::BrokenPipe {
                        return ExitCode::SUCCESS;
                    }
                    format.error(&format!("failed to write output: {e}"));
                    return ExitCode::FAILURE;
                }
            }
            Err(e) => {
                ok = false;
                format.error(&format!("failed to analyse {name}: {e}"));
            }
        }
    }
    if ok {
        ExitCode::SUCCESS
    } else {
        ExitCode::FAILURE
    }
}
