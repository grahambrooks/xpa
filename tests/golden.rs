//! Parity with the Go implementation. Every `tests/corpus/*.xpath` was run
//! through the Go build, and `tests/golden/<name>.json` is what it reported:
//! the analysis, or `{"rejected": true}` for a syntax error.

use std::fs;
use std::path::Path;

use serde_json::{Value, json};

#[test]
fn matches_go_on_every_corpus_expression() {
    let corpus = Path::new(env!("CARGO_MANIFEST_DIR")).join("tests/corpus");
    let golden = Path::new(env!("CARGO_MANIFEST_DIR")).join("tests/golden");
    let mut files: Vec<_> = fs::read_dir(&corpus)
        .unwrap()
        .map(|e| e.unwrap().path())
        .filter(|p| p.extension().is_some_and(|e| e == "xpath"))
        .collect();
    files.sort();
    assert!(!files.is_empty());

    let mut failures = Vec::new();
    for path in &files {
        let name = path.file_stem().unwrap().to_str().unwrap();
        let text = fs::read_to_string(path).unwrap();
        let expected: Value =
            serde_json::from_str(&fs::read_to_string(golden.join(format!("{name}.json"))).unwrap())
                .unwrap();
        let actual = match xpa::analyse(&text) {
            Ok(analysis) => serde_json::to_value(analysis).unwrap(),
            Err(_) => json!({ "rejected": true }),
        };
        if actual != expected {
            failures.push(format!("{name}:\n  go:   {expected}\n  rust: {actual}"));
        }
    }
    assert!(
        failures.is_empty(),
        "{} of {} differ from Go:\n{}",
        failures.len(),
        files.len(),
        failures.join("\n")
    );
}
