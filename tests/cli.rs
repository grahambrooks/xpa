use std::io::Write;
use std::process::{Command, Output, Stdio};

fn xpa(args: &[&str], stdin: &str) -> Output {
    let mut child = Command::new(env!("CARGO_BIN_EXE_xpa"))
        .args(args)
        .stdin(Stdio::piped())
        .stdout(Stdio::piped())
        .stderr(Stdio::piped())
        .spawn()
        .unwrap();
    child
        .stdin
        .take()
        .unwrap()
        .write_all(stdin.as_bytes())
        .unwrap();
    child.wait_with_output().unwrap()
}

fn stdout(o: &Output) -> String {
    String::from_utf8(o.stdout.clone()).unwrap()
}

#[test]
fn reads_stdin_and_keeps_the_trailing_newline() {
    let o = xpa(&["-j"], "/bookstore/book[price>35]/price\n");
    assert!(o.status.success());
    let v: serde_json::Value = serde_json::from_str(&stdout(&o)).unwrap();
    assert_eq!(v["terminal_count"], 12);
    assert_eq!(v["comparison_change"], 1);
    assert_eq!(v["pattern"], "/bookstore/book[price>35]/price\n");
}

#[test]
fn json_is_one_object_per_argument_in_field_order() {
    let o = xpa(&["--json-output", "a and b", "f(1)"], "");
    assert!(o.status.success());
    assert_eq!(
        stdout(&o),
        concat!(
            r#"{"terminal_count":4,"comparison_change":0,"boolean_operator_change":1,"function_call_count":0,"pattern":"a and b"}"#,
            "\n",
            r#"{"terminal_count":5,"comparison_change":0,"boolean_operator_change":0,"function_call_count":1,"pattern":"f(1)"}"#,
            "\n"
        )
    );
}

#[test]
fn analyses_the_contents_of_a_file_argument() {
    let path = concat!(
        env!("CARGO_MANIFEST_DIR"),
        "/tests/corpus/readme-bookstore.xpath"
    );
    let o = xpa(&["-j", path], "");
    assert!(o.status.success());
    assert!(stdout(&o).contains(r#""pattern":"/bookstore/book[price>35]/price""#));
}

#[test]
fn a_directory_argument_is_an_expression() {
    // `.` exists on disk, but it is a directory and a valid XPath expression.
    let o = xpa(&["-j", "."], "");
    assert!(o.status.success());
    assert!(stdout(&o).contains(r#""pattern":".""#));
}

#[test]
fn plain_output_when_piped() {
    let o = xpa(&["f(a) and b = c"], "");
    assert!(o.status.success());
    assert_eq!(
        stdout(&o),
        "complexity: 5, function calls 1, boolean changes: 1, comparison changes 1, tokens 9 pattern f(a) and b = c\n"
    );
    assert_eq!(String::from_utf8_lossy(&o.stderr), "XPath Analyser (xpa)\n");
}

#[test]
fn negative_number_is_an_expression() {
    let o = xpa(&["-j", "-1"], "");
    assert!(o.status.success());
    assert!(stdout(&o).contains(r#""pattern":"-1""#));
}

#[test]
fn invalid_input_reports_the_position_and_fails() {
    let o = xpa(&["-j", "a", "a b"], "");
    assert!(!o.status.success());
    assert_eq!(stdout(&o).lines().count(), 1);
    let stderr = String::from_utf8_lossy(&o.stderr);
    assert!(
        stderr.contains("failed to analyse `a b`: line 1 column 3: unexpected `b`"),
        "{stderr}"
    );
}
