//! Tokens for XPath 3.1, following the lexer rules of `docs/XPath31.g4`.
//!
//! The rules are matched the way ANTLR matches them: longest match wins, and a
//! keyword wins a tie with a name. So `and` is the `and` keyword and `and-x` is
//! a name. Whitespace and comments produce no tokens.

use std::collections::{HashMap, HashSet};
use std::fmt;

use super::SyntaxError;

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
pub enum Kw {
    Ancestor,
    AncestorOrSelf,
    And,
    Array,
    As,
    Attribute,
    Cast,
    Castable,
    Child,
    Comment,
    Descendant,
    DescendantOrSelf,
    Div,
    DocumentNode,
    Element,
    Else,
    EmptySequence,
    Eq,
    Every,
    Except,
    Following,
    FollowingSibling,
    For,
    Function,
    Ge,
    Gt,
    Idiv,
    If,
    In,
    Instance,
    Intersect,
    Is,
    Item,
    Le,
    Let,
    Lt,
    Map,
    Mod,
    Namespace,
    NamespaceNode,
    Ne,
    Node,
    Of,
    Or,
    Parent,
    Preceding,
    PrecedingSibling,
    ProcessingInstruction,
    Return,
    Satisfies,
    SchemaAttribute,
    SchemaElement,
    SelfAxis,
    Some,
    Text,
    Then,
    To,
    Treat,
    Union,
}

const KEYWORDS: &[(&str, Kw)] = &[
    ("ancestor", Kw::Ancestor),
    ("ancestor-or-self", Kw::AncestorOrSelf),
    ("and", Kw::And),
    ("array", Kw::Array),
    ("as", Kw::As),
    ("attribute", Kw::Attribute),
    ("cast", Kw::Cast),
    ("castable", Kw::Castable),
    ("child", Kw::Child),
    ("comment", Kw::Comment),
    ("descendant", Kw::Descendant),
    ("descendant-or-self", Kw::DescendantOrSelf),
    ("div", Kw::Div),
    ("document-node", Kw::DocumentNode),
    ("element", Kw::Element),
    ("else", Kw::Else),
    ("empty-sequence", Kw::EmptySequence),
    ("eq", Kw::Eq),
    ("every", Kw::Every),
    ("except", Kw::Except),
    ("following", Kw::Following),
    ("following-sibling", Kw::FollowingSibling),
    ("for", Kw::For),
    ("function", Kw::Function),
    ("ge", Kw::Ge),
    ("gt", Kw::Gt),
    ("idiv", Kw::Idiv),
    ("if", Kw::If),
    ("in", Kw::In),
    ("instance", Kw::Instance),
    ("intersect", Kw::Intersect),
    ("is", Kw::Is),
    ("item", Kw::Item),
    ("le", Kw::Le),
    ("let", Kw::Let),
    ("lt", Kw::Lt),
    ("map", Kw::Map),
    ("mod", Kw::Mod),
    ("namespace", Kw::Namespace),
    ("namespace-node", Kw::NamespaceNode),
    ("ne", Kw::Ne),
    ("node", Kw::Node),
    ("of", Kw::Of),
    ("or", Kw::Or),
    ("parent", Kw::Parent),
    ("preceding", Kw::Preceding),
    ("preceding-sibling", Kw::PrecedingSibling),
    ("processing-instruction", Kw::ProcessingInstruction),
    ("return", Kw::Return),
    ("satisfies", Kw::Satisfies),
    ("schema-attribute", Kw::SchemaAttribute),
    ("schema-element", Kw::SchemaElement),
    ("self", Kw::SelfAxis),
    ("some", Kw::Some),
    ("text", Kw::Text),
    ("then", Kw::Then),
    ("to", Kw::To),
    ("treat", Kw::Treat),
    ("union", Kw::Union),
];

impl Kw {
    fn lookup(text: &str) -> Option<Kw> {
        KEYWORDS.iter().find(|(s, _)| *s == text).map(|(_, k)| *k)
    }

    pub fn as_str(self) -> &'static str {
        KEYWORDS
            .iter()
            .find(|(_, k)| *k == self)
            .map_or("", |(s, _)| s)
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
pub enum Tok {
    At,
    Bang,
    CloseBracket,
    CloseBrace,
    ColonEq,
    Colon,
    ColonColon,
    Comma,
    CloseParen,
    ColonStar,
    Dot,
    DotDot,
    Dollar,
    Arrow,
    Eq,
    Ge,
    GtGt,
    Gt,
    Le,
    LtLt,
    Lt,
    Minus,
    Ne,
    OpenBracket,
    OpenBrace,
    OpenParen,
    Pipe,
    Plus,
    Pound,
    PipePipe,
    Question,
    StarColon,
    Slash,
    SlashSlash,
    Star,
    Semicolon,
    Kw(Kw),
    Integer,
    Decimal,
    Double,
    String,
    UriQualifiedName,
    BracedUri,
    /// A name, prefixed or not. The grammar's separate `NCName` token can never
    /// win a tie with `QName`, so this is the only name token.
    QName,
    Eof,
}

impl fmt::Display for Tok {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        let s = match self {
            Tok::At => "`@`",
            Tok::Bang => "`!`",
            Tok::CloseBracket => "`]`",
            Tok::CloseBrace => "`}`",
            Tok::ColonEq => "`:=`",
            Tok::Colon => "`:`",
            Tok::ColonColon => "`::`",
            Tok::Comma => "`,`",
            Tok::CloseParen => "`)`",
            Tok::ColonStar => "`:*`",
            Tok::Dot => "`.`",
            Tok::DotDot => "`..`",
            Tok::Dollar => "`$`",
            Tok::Arrow => "`=>`",
            Tok::Eq => "`=`",
            Tok::Ge => "`>=`",
            Tok::GtGt => "`>>`",
            Tok::Gt => "`>`",
            Tok::Le => "`<=`",
            Tok::LtLt => "`<<`",
            Tok::Lt => "`<`",
            Tok::Minus => "`-`",
            Tok::Ne => "`!=`",
            Tok::OpenBracket => "`[`",
            Tok::OpenBrace => "`{`",
            Tok::OpenParen => "`(`",
            Tok::Pipe => "`|`",
            Tok::Plus => "`+`",
            Tok::Pound => "`#`",
            Tok::PipePipe => "`||`",
            Tok::Question => "`?`",
            Tok::StarColon => "`*:`",
            Tok::Slash => "`/`",
            Tok::SlashSlash => "`//`",
            Tok::Star => "`*`",
            Tok::Semicolon => "`;`",
            Tok::Kw(k) => return write!(f, "`{}`", k.as_str()),
            Tok::Integer | Tok::Decimal | Tok::Double => "a number",
            Tok::String => "a string",
            Tok::UriQualifiedName | Tok::QName => "a name",
            Tok::BracedUri => "a braced URI",
            Tok::Eof => "end of input",
        };
        f.write_str(s)
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
pub struct Token {
    pub kind: Tok,
    pub start: usize,
    pub end: usize,
    pub line: usize,
    pub column: usize,
}

/// Tokenises `text`. The result always ends with a single `Tok::Eof`.
pub fn tokenize(text: &str) -> Result<Vec<Token>, SyntaxError> {
    let mut lexer = Lexer {
        src: text,
        pos: 0,
        line: 1,
        line_start: 0,
        comments: HashMap::new(),
    };
    let mut tokens = Vec::new();
    loop {
        lexer.skip_trivia();
        let (line, column) = lexer.location(lexer.pos);
        let start = lexer.pos;
        let kind = lexer.next_kind()?;
        tokens.push(Token {
            kind,
            start,
            end: lexer.pos,
            line,
            column,
        });
        if kind == Tok::Eof {
            return Ok(tokens);
        }
    }
}

struct Lexer<'a> {
    src: &'a str,
    pos: usize,
    line: usize,
    line_start: usize,
    /// `comment_ends` results by start offset, shared across the whole input.
    comments: HashMap<usize, Vec<usize>>,
}

impl Lexer<'_> {
    fn rest(&self) -> &str {
        &self.src[self.pos..]
    }

    fn peek(&self) -> Option<char> {
        self.rest().chars().next()
    }

    fn peek_at(&self, n: usize) -> Option<char> {
        self.rest().chars().nth(n)
    }

    fn location(&self, pos: usize) -> (usize, usize) {
        (
            self.line,
            self.src[self.line_start..pos].chars().count() + 1,
        )
    }

    fn advance_to(&mut self, end: usize) {
        for (i, c) in self.src[self.pos..end].char_indices() {
            if c == '\n' {
                self.line += 1;
                self.line_start = self.pos + i + 1;
            }
        }
        self.pos = end;
    }

    fn error(&self, message: String) -> SyntaxError {
        let (line, column) = self.location(self.pos);
        SyntaxError {
            line,
            column,
            message,
        }
    }

    fn skip_trivia(&mut self) {
        loop {
            let ws = self
                .rest()
                .find(|c| !matches!(c, '\r' | '\n' | ' ' | '\t'))
                .unwrap_or(self.rest().len());
            if ws > 0 {
                self.advance_to(self.pos + ws);
                continue;
            }
            // An unterminated comment is not a comment: ANTLR falls back to
            // lexing `(` and `:`, which the parser then rejects.
            match comment_ends(self.src, self.pos, &mut self.comments)
                .into_iter()
                .max()
            {
                Some(end) => self.advance_to(end),
                None => return,
            }
        }
    }

    fn next_kind(&mut self) -> Result<Tok, SyntaxError> {
        let Some(c) = self.peek() else {
            return Ok(Tok::Eof);
        };
        let next = self.peek_at(1);

        if c.is_ascii_digit() || (c == '.' && next.is_some_and(|d| d.is_ascii_digit())) {
            return Ok(self.number());
        }
        if c == '"' || c == '\'' {
            return self.string(c);
        }
        if c == 'Q'
            && next == Some('{')
            && let Some(kind) = self.braced_uri()
        {
            return Ok(kind);
        }
        if is_name_start(c) {
            return Ok(self.name());
        }

        let (kind, len) = match (c, next) {
            (':', Some(':')) => (Tok::ColonColon, 2),
            (':', Some('=')) => (Tok::ColonEq, 2),
            (':', Some('*')) => (Tok::ColonStar, 2),
            (':', _) => (Tok::Colon, 1),
            ('.', Some('.')) => (Tok::DotDot, 2),
            ('.', _) => (Tok::Dot, 1),
            ('!', Some('=')) => (Tok::Ne, 2),
            ('!', _) => (Tok::Bang, 1),
            ('=', Some('>')) => (Tok::Arrow, 2),
            ('=', _) => (Tok::Eq, 1),
            ('>', Some('=')) => (Tok::Ge, 2),
            ('>', Some('>')) => (Tok::GtGt, 2),
            ('>', _) => (Tok::Gt, 1),
            ('<', Some('=')) => (Tok::Le, 2),
            ('<', Some('<')) => (Tok::LtLt, 2),
            ('<', _) => (Tok::Lt, 1),
            ('|', Some('|')) => (Tok::PipePipe, 2),
            ('|', _) => (Tok::Pipe, 1),
            ('*', Some(':')) => (Tok::StarColon, 2),
            ('*', _) => (Tok::Star, 1),
            ('/', Some('/')) => (Tok::SlashSlash, 2),
            ('/', _) => (Tok::Slash, 1),
            ('@', _) => (Tok::At, 1),
            (']', _) => (Tok::CloseBracket, 1),
            ('}', _) => (Tok::CloseBrace, 1),
            (',', _) => (Tok::Comma, 1),
            (')', _) => (Tok::CloseParen, 1),
            ('$', _) => (Tok::Dollar, 1),
            ('-', _) => (Tok::Minus, 1),
            ('[', _) => (Tok::OpenBracket, 1),
            ('{', _) => (Tok::OpenBrace, 1),
            ('(', _) => (Tok::OpenParen, 1),
            ('+', _) => (Tok::Plus, 1),
            ('#', _) => (Tok::Pound, 1),
            ('?', _) => (Tok::Question, 1),
            (';', _) => (Tok::Semicolon, 1),
            _ => return Err(self.error(format!("unexpected character `{c}`"))),
        };
        self.advance_to(self.pos + len);
        Ok(kind)
    }

    /// `IntegerLiteral`, `DecimalLiteral` or `DoubleLiteral`, whichever is longest.
    fn number(&mut self) -> Tok {
        let bytes = self.rest().as_bytes();
        let digits = |from: usize| {
            bytes[from..]
                .iter()
                .take_while(|b| b.is_ascii_digit())
                .count()
        };
        let mut len = digits(0);
        let mut kind = Tok::Integer;
        if bytes.get(len) == Some(&b'.') {
            len += 1 + digits(len + 1);
            kind = Tok::Decimal;
        }
        if matches!(bytes.get(len), Some(b'e' | b'E')) {
            let mut exp = len + 1;
            if matches!(bytes.get(exp), Some(b'+' | b'-')) {
                exp += 1;
            }
            let exp_digits = digits(exp);
            if exp_digits > 0 {
                len = exp + exp_digits;
                kind = Tok::Double;
            }
        }
        self.advance_to(self.pos + len);
        kind
    }

    /// A string literal, with a doubled quote as the escape for the quote.
    fn string(&mut self, quote: char) -> Result<Tok, SyntaxError> {
        let rest = self.rest();
        let mut chars = rest.char_indices().skip(1).peekable();
        while let Some((i, c)) = chars.next() {
            if c == quote {
                if chars.peek().is_some_and(|&(_, n)| n == quote) {
                    chars.next();
                    continue;
                }
                self.advance_to(self.pos + i + 1);
                return Ok(Tok::String);
            }
        }
        Err(self.error("unterminated string literal".to_string()))
    }

    /// `Q{uri}` alone (`BracedURILiteral`) or followed by a local name
    /// (`URIQualifiedName`). `None` means there's no closing brace, so `Q` is
    /// just a name.
    fn braced_uri(&mut self) -> Option<Tok> {
        let rest = self.rest();
        let close = rest[2..].find(['{', '}'])? + 2;
        if rest.as_bytes()[close] != b'}' {
            return None;
        }
        let after = close + 1;
        let local = ncname_len(&rest[after..]);
        self.advance_to(self.pos + after + local);
        Some(if local > 0 {
            Tok::UriQualifiedName
        } else {
            Tok::BracedUri
        })
    }

    /// `QName`: an NCName, optionally followed by `:` and another NCName. An
    /// unprefixed name that spells a keyword is that keyword.
    fn name(&mut self) -> Tok {
        let rest = self.rest();
        let mut len = ncname_len(rest);
        let mut prefixed = false;
        if rest[len..].starts_with(':') {
            let local = ncname_len(&rest[len + 1..]);
            if local > 0 {
                len += 1 + local;
                prefixed = true;
            }
        }
        let kind = match (prefixed, Kw::lookup(&rest[..len])) {
            (false, Some(kw)) => Tok::Kw(kw),
            _ => Tok::QName,
        };
        self.advance_to(self.pos + len);
        kind
    }
}

/// The ends of a `(: … :)` comment at `start` in `s`; empty if there is no
/// complete one.
///
/// This is the ANTLR rule `'(:' (Comment | CommentContents)*? ':)'`. An inner
/// `(:` may open a nested comment or be plain characters, and along either
/// reading the non-greedy loop stops at the first `:)` it reaches. The lexer
/// then keeps the longest comment any reading completes. So
/// `(: a (: b :) :)` nests, but in `(: a : (: b :)` the nested reading never
/// closes, and the comment ends at the only `:)`.
///
/// Returns every offset at which a comment starting at `start` can end; the
/// caller takes the largest. Results are memoised by start offset.
fn comment_ends(s: &str, start: usize, memo: &mut HashMap<usize, Vec<usize>>) -> Vec<usize> {
    if !s[start..].starts_with("(:") {
        return Vec::new();
    }
    if let Some(ends) = memo.get(&start) {
        return ends.clone();
    }
    // A comment's ends depend only on comments that start after it, so fill
    // those in from the last one back. Then the scan never recurses, however
    // deeply `(:` nests.
    let later: Vec<usize> = s[start + 1..]
        .match_indices("(:")
        .map(|(i, _)| start + 1 + i)
        .filter(|p| !memo.contains_key(p))
        .collect();
    for &p in later.iter().rev().chain([start].iter()) {
        let ends = scan_comment(s, p, memo);
        memo.insert(p, ends);
    }
    memo[&start].clone()
}

/// The ends of the comment at `start`, given the ends of every later one.
fn scan_comment(s: &str, start: usize, memo: &HashMap<usize, Vec<usize>>) -> Vec<usize> {
    let mut ends = Vec::new();
    if !s[start + 2..].contains(":)") {
        return ends;
    }
    let mut seen = HashSet::new();
    let mut pending = vec![start + 2];
    while let Some(i) = pending.pop() {
        if !seen.insert(i) {
            continue;
        }
        let rest = &s[i..];
        let Some(c) = rest.chars().next() else {
            continue;
        };
        if rest.starts_with(":)") {
            ends.push(i + 2);
            continue;
        }
        pending.push(i + c.len_utf8());
        if let Some(nested) = memo.get(&i) {
            pending.extend(nested);
        }
    }
    ends
}

fn ncname_len(s: &str) -> usize {
    let mut chars = s.char_indices();
    match chars.next() {
        Some((_, c)) if is_name_start(c) => {}
        _ => return 0,
    }
    chars
        .find(|&(_, c)| !is_name_char(c))
        .map_or(s.len(), |(i, _)| i)
}

fn is_name_start(c: char) -> bool {
    matches!(c,
        'A'..='Z'
        | '_'
        | 'a'..='z'
        | '\u{C0}'..='\u{D6}'
        | '\u{D8}'..='\u{F6}'
        | '\u{F8}'..='\u{2FF}'
        | '\u{370}'..='\u{37D}'
        | '\u{37F}'..='\u{1FFF}'
        | '\u{200C}'..='\u{200D}'
        | '\u{2070}'..='\u{218F}'
        | '\u{2C00}'..='\u{2FEF}'
        | '\u{3001}'..='\u{D7FF}'
        | '\u{F900}'..='\u{FDCF}'
        | '\u{FDF0}'..='\u{FFFD}'
        | '\u{10000}'..='\u{EFFFF}')
}

fn is_name_char(c: char) -> bool {
    is_name_start(c)
        || matches!(c,
            '-' | '.' | '0'..='9' | '\u{B7}' | '\u{300}'..='\u{36F}' | '\u{203F}'..='\u{2040}')
}

#[cfg(test)]
mod tests {
    use super::*;

    fn kinds(text: &str) -> Vec<Tok> {
        tokenize(text)
            .unwrap()
            .into_iter()
            .map(|t| t.kind)
            .collect()
    }

    #[test]
    fn keyword_wins_a_tie_but_not_a_longer_name() {
        assert_eq!(kinds("and"), [Tok::Kw(Kw::And), Tok::Eof]);
        assert_eq!(kinds("and-x"), [Tok::QName, Tok::Eof]);
        assert_eq!(kinds("fn:and"), [Tok::QName, Tok::Eof]);
    }

    #[test]
    fn axis_is_not_a_prefix() {
        assert_eq!(
            kinds("child::a"),
            [Tok::Kw(Kw::Child), Tok::ColonColon, Tok::QName, Tok::Eof]
        );
    }

    #[test]
    fn numbers_take_the_longest_form() {
        assert_eq!(kinds("1"), [Tok::Integer, Tok::Eof]);
        assert_eq!(kinds("5."), [Tok::Decimal, Tok::Eof]);
        assert_eq!(kinds(".5"), [Tok::Decimal, Tok::Eof]);
        assert_eq!(kinds("2.e1"), [Tok::Double, Tok::Eof]);
        assert_eq!(kinds("1e"), [Tok::Integer, Tok::QName, Tok::Eof]);
    }

    #[test]
    fn doubled_quote_escapes_the_quote() {
        assert_eq!(kinds("'it''s'"), [Tok::String, Tok::Eof]);
        assert_eq!(kinds(r#""a""b""#), [Tok::String, Tok::Eof]);
    }

    #[test]
    fn braced_uri_takes_any_uri() {
        assert_eq!(kinds("Q{http://x}a"), [Tok::UriQualifiedName, Tok::Eof]);
        assert_eq!(kinds("Q{}*"), [Tok::BracedUri, Tok::Star, Tok::Eof]);
        assert_eq!(kinds("Q"), [Tok::QName, Tok::Eof]);
    }

    #[test]
    fn comments_nest_and_are_skipped() {
        assert_eq!(kinds("1 (: a (: b :) c :) 2").len(), 3);
        assert_eq!(kinds("(: open")[..2], [Tok::OpenParen, Tok::Colon]);
        assert_eq!(kinds("(: c (:) 1"), [Tok::Integer, Tok::Eof]);
        assert_eq!(kinds("(: c : (: c :) .5"), [Tok::Decimal, Tok::Eof]);
        let unclosed = "(: ".repeat(2000);
        assert!(tokenize(&unclosed).is_ok_and(|t| t.len() > 1));
        // One `:)` closes the outermost comment, with every inner `(` as text.
        let closed = format!("{} :) 1", "(: ".repeat(200));
        assert_eq!(kinds(&closed), [Tok::Integer, Tok::Eof]);
    }

    #[test]
    fn unknown_character_is_an_error() {
        let err = tokenize("1 ^ 2").unwrap_err();
        assert_eq!((err.line, err.column), (1, 3));
    }
}
