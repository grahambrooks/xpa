//! Recursive descent over the parser rules of `docs/XPath31.g4`.
//!
//! Almost every decision is settled by the next two tokens. Two are not, and
//! ANTLR's full-context prediction accepts both readings there, so this
//! parser does too:
//!
//! - a lone `/` followed by a keyword or `*`: `/ * 5` is `/` times 5, while
//!   `/*` is a path;
//! - an occurrence indicator `*` or `+` after a sequence type: in
//!   `$a instance of xs:int + 1`, the `+` is an addition.
//!
//! Each of these is a *choice point*. The first attempt takes the first
//! alternative everywhere. If the parse fails, the last choice that has an
//! alternative left is advanced and the whole input is parsed again, depth
//! first, until a parse succeeds or there are no choices left. Only keywords
//! and `*` create choice points, so real expressions rarely need a second
//! attempt.

use super::SyntaxError;
use super::lexer::{Kw, Tok, Token, tokenize};

/// The most parses one expression may take before it is reported as invalid.
const MAX_ATTEMPTS: usize = 1024;

#[derive(Debug)]
pub struct Parsed {
    /// Every token, ending with `Tok::Eof`.
    pub tokens: Vec<Token>,
    /// `functioncall` matches: a name, not a reserved one, applied to an
    /// argument list. Arrow, dynamic and named-reference calls are not counted.
    pub function_calls: usize,
}

pub fn parse(text: &str) -> Result<Parsed, SyntaxError> {
    let tokens = tokenize(text)?;
    let mut script: Vec<Choice> = Vec::new();
    let mut first_error = None;
    for _ in 0..MAX_ATTEMPTS {
        let mut parser = Parser {
            src: text,
            tokens: &tokens,
            pos: 0,
            script,
            cursor: 0,
            function_calls: 0,
        };
        match parser.xpath() {
            Ok(()) => {
                return Ok(Parsed {
                    function_calls: parser.function_calls,
                    tokens,
                });
            }
            Err(e) => {
                let error = first_error.get_or_insert(e);
                script = parser.script;
                script.truncate(parser.cursor);
                if !advance(&mut script) {
                    return Err(error.clone());
                }
            }
        }
    }
    Err(first_error.expect("at least one attempt"))
}

#[derive(Clone, Copy)]
struct Choice {
    taken: u8,
    alternatives: u8,
}

/// Steps `script` to the next untried combination of choices.
fn advance(script: &mut Vec<Choice>) -> bool {
    while let Some(choice) = script.pop() {
        if choice.taken + 1 < choice.alternatives {
            script.push(Choice {
                taken: choice.taken + 1,
                ..choice
            });
            return true;
        }
    }
    false
}

type R = Result<(), SyntaxError>;

struct Parser<'a> {
    src: &'a str,
    tokens: &'a [Token],
    pos: usize,
    script: Vec<Choice>,
    cursor: usize,
    function_calls: usize,
}

impl Parser<'_> {
    // ---- token helpers ----

    fn peek(&self) -> Tok {
        self.peek_n(0)
    }

    fn peek_n(&self, n: usize) -> Tok {
        let i = (self.pos + n).min(self.tokens.len() - 1);
        self.tokens[i].kind
    }

    fn bump(&mut self) {
        if self.pos < self.tokens.len() - 1 {
            self.pos += 1;
        }
    }

    fn eat(&mut self, tok: Tok) -> bool {
        if self.peek() == tok {
            self.bump();
            true
        } else {
            false
        }
    }

    fn expect(&mut self, tok: Tok) -> R {
        if self.eat(tok) {
            Ok(())
        } else {
            Err(self.unexpected(&tok.to_string()))
        }
    }

    fn unexpected(&self, expected: &str) -> SyntaxError {
        let token = &self.tokens[self.pos];
        let found = match token.kind {
            Tok::Eof => "end of input".to_string(),
            _ => format!("`{}`", &self.src[token.start..token.end]),
        };
        SyntaxError {
            line: token.line,
            column: token.column,
            message: format!("unexpected {found}, expected {expected}"),
        }
    }

    /// Takes the next alternative at a choice point: from the script on a
    /// replay, otherwise the first one, recorded for the next attempt.
    fn choice(&mut self, alternatives: u8) -> u8 {
        let taken = match self.script.get(self.cursor) {
            Some(choice) => choice.taken,
            None => {
                self.script.push(Choice {
                    taken: 0,
                    alternatives,
                });
                0
            }
        };
        self.cursor += 1;
        taken
    }

    /// An unprefixed name or any keyword, which is what `NCName` matches.
    fn is_ncname(&self, n: usize) -> bool {
        let i = (self.pos + n).min(self.tokens.len() - 1);
        let token = &self.tokens[i];
        match token.kind {
            Tok::Kw(_) => true,
            Tok::QName => !self.src[token.start..token.end].contains(':'),
            _ => false,
        }
    }

    fn eqname(&mut self) -> R {
        if is_eqname(self.peek()) {
            self.bump();
            Ok(())
        } else {
            Err(self.unexpected("a name"))
        }
    }

    // ---- [1]–[15] expressions ----

    fn xpath(&mut self) -> R {
        self.expr()?;
        if self.peek() == Tok::Eof {
            Ok(())
        } else {
            Err(self.unexpected("an operator or end of input"))
        }
    }

    fn expr(&mut self) -> R {
        self.exprsingle()?;
        while self.eat(Tok::Comma) {
            self.exprsingle()?;
        }
        Ok(())
    }

    fn exprsingle(&mut self) -> R {
        match (self.peek(), self.peek_n(1)) {
            (Tok::Kw(Kw::For), Tok::Dollar) => self.forexpr(),
            (Tok::Kw(Kw::Let), Tok::Dollar) => self.letexpr(),
            (Tok::Kw(Kw::Some | Kw::Every), Tok::Dollar) => self.quantifiedexpr(),
            (Tok::Kw(Kw::If), Tok::OpenParen) => self.ifexpr(),
            _ => self.orexpr(),
        }
    }

    fn forexpr(&mut self) -> R {
        self.bump();
        loop {
            self.expect(Tok::Dollar)?;
            self.eqname()?;
            self.expect(Tok::Kw(Kw::In))?;
            self.exprsingle()?;
            if !self.eat(Tok::Comma) {
                break;
            }
        }
        self.expect(Tok::Kw(Kw::Return))?;
        self.exprsingle()
    }

    fn letexpr(&mut self) -> R {
        self.bump();
        loop {
            self.expect(Tok::Dollar)?;
            self.eqname()?;
            self.expect(Tok::ColonEq)?;
            self.exprsingle()?;
            if !self.eat(Tok::Comma) {
                break;
            }
        }
        self.expect(Tok::Kw(Kw::Return))?;
        self.exprsingle()
    }

    fn quantifiedexpr(&mut self) -> R {
        self.bump();
        loop {
            self.expect(Tok::Dollar)?;
            self.eqname()?;
            self.expect(Tok::Kw(Kw::In))?;
            self.exprsingle()?;
            if !self.eat(Tok::Comma) {
                break;
            }
        }
        self.expect(Tok::Kw(Kw::Satisfies))?;
        self.exprsingle()
    }

    fn ifexpr(&mut self) -> R {
        self.bump();
        self.expect(Tok::OpenParen)?;
        self.expr()?;
        self.expect(Tok::CloseParen)?;
        self.expect(Tok::Kw(Kw::Then))?;
        self.exprsingle()?;
        self.expect(Tok::Kw(Kw::Else))?;
        self.exprsingle()
    }

    /// `operand (op operand)*`
    fn chain(&mut self, operand: fn(&mut Self) -> R, is_op: fn(Tok) -> bool) -> R {
        operand(self)?;
        while is_op(self.peek()) {
            self.bump();
            operand(self)?;
        }
        Ok(())
    }

    fn orexpr(&mut self) -> R {
        self.chain(Self::andexpr, |t| t == Tok::Kw(Kw::Or))
    }

    fn andexpr(&mut self) -> R {
        self.chain(Self::comparisonexpr, |t| t == Tok::Kw(Kw::And))
    }

    // ---- [18]–[35] operators ----

    fn comparisonexpr(&mut self) -> R {
        self.stringconcatexpr()?;
        if is_comparison(self.peek()) {
            self.bump();
            self.stringconcatexpr()?;
        }
        Ok(())
    }

    fn stringconcatexpr(&mut self) -> R {
        self.chain(Self::rangeexpr, |t| t == Tok::PipePipe)
    }

    fn rangeexpr(&mut self) -> R {
        self.additiveexpr()?;
        if self.eat(Tok::Kw(Kw::To)) {
            self.additiveexpr()?;
        }
        Ok(())
    }

    fn additiveexpr(&mut self) -> R {
        self.chain(Self::multiplicativeexpr, |t| {
            matches!(t, Tok::Plus | Tok::Minus)
        })
    }

    fn multiplicativeexpr(&mut self) -> R {
        self.chain(Self::unionexpr, |t| {
            matches!(t, Tok::Star | Tok::Kw(Kw::Div | Kw::Idiv | Kw::Mod))
        })
    }

    fn unionexpr(&mut self) -> R {
        self.chain(Self::intersectexceptexpr, |t| {
            matches!(t, Tok::Pipe | Tok::Kw(Kw::Union))
        })
    }

    fn intersectexceptexpr(&mut self) -> R {
        self.chain(Self::instanceofexpr, |t| {
            matches!(t, Tok::Kw(Kw::Intersect | Kw::Except))
        })
    }

    fn instanceofexpr(&mut self) -> R {
        self.treatexpr()?;
        if self.eat(Tok::Kw(Kw::Instance)) {
            self.expect(Tok::Kw(Kw::Of))?;
            self.sequencetype()?;
        }
        Ok(())
    }

    fn treatexpr(&mut self) -> R {
        self.castableexpr()?;
        if self.eat(Tok::Kw(Kw::Treat)) {
            self.expect(Tok::Kw(Kw::As))?;
            self.sequencetype()?;
        }
        Ok(())
    }

    fn castableexpr(&mut self) -> R {
        self.castexpr()?;
        if self.eat(Tok::Kw(Kw::Castable)) {
            self.expect(Tok::Kw(Kw::As))?;
            self.singletype()?;
        }
        Ok(())
    }

    fn castexpr(&mut self) -> R {
        self.arrowexpr()?;
        if self.eat(Tok::Kw(Kw::Cast)) {
            self.expect(Tok::Kw(Kw::As))?;
            self.singletype()?;
        }
        Ok(())
    }

    fn arrowexpr(&mut self) -> R {
        self.unaryexpr()?;
        while self.eat(Tok::Arrow) {
            // arrowfunctionspecifier: the name here is not a functioncall.
            match self.peek() {
                Tok::Dollar => self.varref()?,
                Tok::OpenParen => self.parenthesizedexpr()?,
                _ => self.eqname()?,
            }
            self.argumentlist()?;
        }
        Ok(())
    }

    fn unaryexpr(&mut self) -> R {
        while matches!(self.peek(), Tok::Plus | Tok::Minus) {
            self.bump();
        }
        self.simplemapexpr()
    }

    // ---- [35]–[55] paths ----

    fn simplemapexpr(&mut self) -> R {
        self.chain(Self::pathexpr, |t| t == Tok::Bang)
    }

    fn pathexpr(&mut self) -> R {
        match self.peek() {
            Tok::Slash => {
                self.bump();
                if self.can_start_step() {
                    let ambiguous = matches!(self.peek(), Tok::Kw(_) | Tok::Star);
                    if !ambiguous || self.choice(2) == 0 {
                        self.relativepathexpr()?;
                    }
                }
                Ok(())
            }
            Tok::SlashSlash => {
                self.bump();
                self.relativepathexpr()
            }
            _ => self.relativepathexpr(),
        }
    }

    fn can_start_step(&self) -> bool {
        let t = self.peek();
        is_eqname(t)
            || matches!(
                t,
                Tok::Integer
                    | Tok::Decimal
                    | Tok::Double
                    | Tok::String
                    | Tok::Dollar
                    | Tok::OpenParen
                    | Tok::Dot
                    | Tok::DotDot
                    | Tok::OpenBracket
                    | Tok::Question
                    | Tok::At
                    | Tok::Star
                    | Tok::StarColon
                    | Tok::BracedUri
            )
    }

    fn relativepathexpr(&mut self) -> R {
        self.stepexpr()?;
        while matches!(self.peek(), Tok::Slash | Tok::SlashSlash) {
            self.bump();
            self.stepexpr()?;
        }
        Ok(())
    }

    fn stepexpr(&mut self) -> R {
        let (t, next) = (self.peek(), self.peek_n(1));
        let postfix = match t {
            Tok::Integer
            | Tok::Decimal
            | Tok::Double
            | Tok::String
            | Tok::Dollar
            | Tok::OpenParen
            | Tok::Dot
            | Tok::OpenBracket
            | Tok::Question => true,
            Tok::Kw(Kw::Function) if next == Tok::OpenParen => true,
            Tok::Kw(Kw::Map | Kw::Array) if next == Tok::OpenBrace => true,
            _ if is_eqname(t) && next == Tok::OpenParen => !is_reserved_function_name(t),
            _ if is_eqname(t) && next == Tok::Pound => true,
            _ => false,
        };
        if postfix {
            self.postfixexpr()
        } else {
            self.axisstep()
        }
    }

    fn axisstep(&mut self) -> R {
        match (self.peek(), self.peek_n(1)) {
            (Tok::Kw(k), Tok::ColonColon) if is_axis(k) => {
                self.bump();
                self.bump();
                self.nodetest()?;
            }
            (Tok::DotDot, _) => self.bump(),
            _ => {
                self.eat(Tok::At);
                self.nodetest()?;
            }
        }
        while self.peek() == Tok::OpenBracket {
            self.predicate()?;
        }
        Ok(())
    }

    fn nodetest(&mut self) -> R {
        match (self.peek(), self.peek_n(1)) {
            (Tok::Kw(k), Tok::OpenParen) if is_kind_test(k) => self.kindtest(k),
            _ => self.nametest(),
        }
    }

    fn nametest(&mut self) -> R {
        match self.peek() {
            Tok::Star => self.bump(),
            Tok::StarColon => {
                self.bump();
                self.ncname()?;
            }
            Tok::BracedUri => {
                self.bump();
                self.expect(Tok::Star)?;
            }
            _ if self.is_ncname(0) && self.peek_n(1) == Tok::ColonStar => {
                self.bump();
                self.bump();
            }
            _ if is_eqname(self.peek()) => self.bump(),
            _ => return Err(self.unexpected("an expression")),
        }
        Ok(())
    }

    fn ncname(&mut self) -> R {
        if self.is_ncname(0) {
            self.bump();
            Ok(())
        } else {
            Err(self.unexpected("an unprefixed name"))
        }
    }

    fn postfixexpr(&mut self) -> R {
        self.primaryexpr()?;
        loop {
            match self.peek() {
                Tok::OpenBracket => self.predicate()?,
                Tok::OpenParen => self.argumentlist()?,
                Tok::Question if self.can_start_keyspecifier(1) => {
                    self.bump();
                    self.keyspecifier()?;
                }
                _ => return Ok(()),
            }
        }
    }

    fn argumentlist(&mut self) -> R {
        self.expect(Tok::OpenParen)?;
        if self.eat(Tok::CloseParen) {
            return Ok(());
        }
        loop {
            // argumentplaceholder
            if self.peek() == Tok::Question
                && matches!(self.peek_n(1), Tok::Comma | Tok::CloseParen)
            {
                self.bump();
            } else {
                self.exprsingle()?;
            }
            if !self.eat(Tok::Comma) {
                break;
            }
        }
        self.expect(Tok::CloseParen)
    }

    fn predicate(&mut self) -> R {
        self.expect(Tok::OpenBracket)?;
        self.expr()?;
        self.expect(Tok::CloseBracket)
    }

    fn can_start_keyspecifier(&self, n: usize) -> bool {
        self.is_ncname(n) || matches!(self.peek_n(n), Tok::Integer | Tok::OpenParen | Tok::Star)
    }

    fn keyspecifier(&mut self) -> R {
        match self.peek() {
            Tok::Integer | Tok::Star => self.bump(),
            Tok::OpenParen => self.parenthesizedexpr()?,
            _ => self.ncname()?,
        }
        Ok(())
    }

    // ---- [56]–[77] primaries ----

    fn primaryexpr(&mut self) -> R {
        let (t, next) = (self.peek(), self.peek_n(1));
        match t {
            Tok::Integer | Tok::Decimal | Tok::Double | Tok::String | Tok::Dot => {
                self.bump();
                Ok(())
            }
            Tok::Dollar => self.varref(),
            Tok::OpenParen => self.parenthesizedexpr(),
            Tok::OpenBracket => self.squarearrayconstructor(),
            Tok::Question => {
                self.bump();
                self.keyspecifier()
            }
            Tok::Kw(Kw::Function) if next == Tok::OpenParen => self.inlinefunctionexpr(),
            Tok::Kw(Kw::Map) if next == Tok::OpenBrace => self.mapconstructor(),
            Tok::Kw(Kw::Array) if next == Tok::OpenBrace => {
                self.bump();
                self.enclosedexpr()
            }
            _ if is_eqname(t) && next == Tok::Pound => {
                self.bump();
                self.bump();
                self.expect(Tok::Integer)
            }
            _ if is_eqname(t) && next == Tok::OpenParen && !is_reserved_function_name(t) => {
                self.function_calls += 1;
                self.bump();
                self.argumentlist()
            }
            _ => Err(self.unexpected("an expression")),
        }
    }

    fn varref(&mut self) -> R {
        self.expect(Tok::Dollar)?;
        self.eqname()
    }

    fn parenthesizedexpr(&mut self) -> R {
        self.expect(Tok::OpenParen)?;
        if self.eat(Tok::CloseParen) {
            return Ok(());
        }
        self.expr()?;
        self.expect(Tok::CloseParen)
    }

    /// `{ expr? }`, the body of a function and of `array { … }`.
    fn enclosedexpr(&mut self) -> R {
        self.expect(Tok::OpenBrace)?;
        if self.eat(Tok::CloseBrace) {
            return Ok(());
        }
        self.expr()?;
        self.expect(Tok::CloseBrace)
    }

    fn inlinefunctionexpr(&mut self) -> R {
        self.bump();
        self.expect(Tok::OpenParen)?;
        if self.peek() == Tok::Dollar {
            loop {
                self.varref()?;
                if self.eat(Tok::Kw(Kw::As)) {
                    self.sequencetype()?;
                }
                if !self.eat(Tok::Comma) {
                    break;
                }
            }
        }
        self.expect(Tok::CloseParen)?;
        if self.eat(Tok::Kw(Kw::As)) {
            self.sequencetype()?;
        }
        self.enclosedexpr()
    }

    fn mapconstructor(&mut self) -> R {
        self.bump();
        self.expect(Tok::OpenBrace)?;
        if self.eat(Tok::CloseBrace) {
            return Ok(());
        }
        loop {
            self.exprsingle()?;
            self.expect(Tok::Colon)?;
            self.exprsingle()?;
            if !self.eat(Tok::Comma) {
                break;
            }
        }
        self.expect(Tok::CloseBrace)
    }

    fn squarearrayconstructor(&mut self) -> R {
        self.expect(Tok::OpenBracket)?;
        if self.eat(Tok::CloseBracket) {
            return Ok(());
        }
        loop {
            self.exprsingle()?;
            if !self.eat(Tok::Comma) {
                break;
            }
        }
        self.expect(Tok::CloseBracket)
    }

    // ---- [77]–[111] types ----

    fn singletype(&mut self) -> R {
        self.eqname()?;
        self.eat(Tok::Question);
        Ok(())
    }

    fn sequencetype(&mut self) -> R {
        if self.peek() == Tok::Kw(Kw::EmptySequence) && self.peek_n(1) == Tok::OpenParen {
            self.bump();
            self.bump();
            return self.expect(Tok::CloseParen);
        }
        self.itemtype()?;
        match self.peek() {
            Tok::Question => self.bump(),
            Tok::Star | Tok::Plus if self.choice(2) == 0 => self.bump(),
            _ => {}
        }
        Ok(())
    }

    fn itemtype(&mut self) -> R {
        match (self.peek(), self.peek_n(1)) {
            (Tok::Kw(k), Tok::OpenParen) if is_kind_test(k) => self.kindtest(k),
            (Tok::Kw(Kw::Item), Tok::OpenParen) => {
                self.bump();
                self.bump();
                self.expect(Tok::CloseParen)
            }
            (Tok::Kw(Kw::Function), Tok::OpenParen) => {
                self.bump();
                self.bump();
                if self.peek() == Tok::Star {
                    self.bump();
                    return self.expect(Tok::CloseParen);
                }
                if !self.eat(Tok::CloseParen) {
                    loop {
                        self.sequencetype()?;
                        if !self.eat(Tok::Comma) {
                            break;
                        }
                    }
                    self.expect(Tok::CloseParen)?;
                }
                self.expect(Tok::Kw(Kw::As))?;
                self.sequencetype()
            }
            (Tok::Kw(Kw::Map), Tok::OpenParen) => {
                self.bump();
                self.bump();
                if !self.eat(Tok::Star) {
                    self.eqname()?;
                    self.expect(Tok::Comma)?;
                    self.sequencetype()?;
                }
                self.expect(Tok::CloseParen)
            }
            (Tok::Kw(Kw::Array), Tok::OpenParen) => {
                self.bump();
                self.bump();
                if !self.eat(Tok::Star) {
                    self.sequencetype()?;
                }
                self.expect(Tok::CloseParen)
            }
            (Tok::OpenParen, _) => {
                self.bump();
                self.itemtype()?;
                self.expect(Tok::CloseParen)
            }
            _ => self.eqname(),
        }
    }

    fn kindtest(&mut self, kind: Kw) -> R {
        self.bump();
        self.expect(Tok::OpenParen)?;
        match kind {
            Kw::DocumentNode => {
                if let (Tok::Kw(k @ (Kw::Element | Kw::SchemaElement)), Tok::OpenParen) =
                    (self.peek(), self.peek_n(1))
                {
                    self.kindtest(k)?;
                }
            }
            Kw::Element | Kw::Attribute if self.peek() != Tok::CloseParen => {
                if !self.eat(Tok::Star) {
                    self.eqname()?;
                }
                if self.eat(Tok::Comma) {
                    self.eqname()?;
                    if kind == Kw::Element {
                        self.eat(Tok::Question);
                    }
                }
            }
            Kw::SchemaElement | Kw::SchemaAttribute => self.eqname()?,
            Kw::ProcessingInstruction if self.is_ncname(0) || self.peek() == Tok::String => {
                self.bump();
            }
            _ => {}
        }
        self.expect(Tok::CloseParen)
    }
}

/// `eqname` accepts every keyword except `to`.
fn is_eqname(t: Tok) -> bool {
    match t {
        Tok::QName | Tok::UriQualifiedName => true,
        Tok::Kw(k) => k != Kw::To,
        _ => false,
    }
}

fn is_comparison(t: Tok) -> bool {
    matches!(
        t,
        Tok::Eq
            | Tok::Ne
            | Tok::Lt
            | Tok::Le
            | Tok::Gt
            | Tok::Ge
            | Tok::LtLt
            | Tok::GtGt
            | Tok::Kw(Kw::Eq | Kw::Ne | Kw::Lt | Kw::Le | Kw::Gt | Kw::Ge | Kw::Is)
    )
}

fn is_axis(k: Kw) -> bool {
    matches!(
        k,
        Kw::Child
            | Kw::Descendant
            | Kw::Attribute
            | Kw::SelfAxis
            | Kw::DescendantOrSelf
            | Kw::FollowingSibling
            | Kw::Following
            | Kw::Namespace
            | Kw::Parent
            | Kw::Ancestor
            | Kw::PrecedingSibling
            | Kw::Preceding
            | Kw::AncestorOrSelf
    )
}

fn is_kind_test(k: Kw) -> bool {
    matches!(
        k,
        Kw::DocumentNode
            | Kw::Element
            | Kw::Attribute
            | Kw::SchemaElement
            | Kw::SchemaAttribute
            | Kw::ProcessingInstruction
            | Kw::Comment
            | Kw::Text
            | Kw::NamespaceNode
            | Kw::Node
    )
}

/// The names the `functioncall` predicate in the grammar refuses
/// (XPath 3.1 A.3 reserved function names).
fn is_reserved_function_name(t: Tok) -> bool {
    matches!(
        t,
        Tok::Kw(
            Kw::Array
                | Kw::Attribute
                | Kw::Comment
                | Kw::DocumentNode
                | Kw::Element
                | Kw::EmptySequence
                | Kw::Function
                | Kw::If
                | Kw::Item
                | Kw::Map
                | Kw::NamespaceNode
                | Kw::Node
                | Kw::ProcessingInstruction
                | Kw::SchemaAttribute
                | Kw::SchemaElement
                | Kw::Text
        )
    )
}

#[cfg(test)]
mod tests {
    use super::*;

    fn calls(text: &str) -> usize {
        parse(text).unwrap().function_calls
    }

    #[test]
    fn counts_named_calls_only() {
        assert_eq!(calls("f()"), 1);
        assert_eq!(calls("fn:count(a)"), 1);
        assert_eq!(calls("foo(1)(2)"), 1);
        assert_eq!(calls("and(1)"), 1);
        assert_eq!(calls("1 => f()"), 0);
        assert_eq!(calls("$f(1)"), 0);
        assert_eq!(calls("count#1"), 0);
        assert_eq!(calls("text()"), 0);
    }

    #[test]
    fn lone_slash_backtracks() {
        assert!(parse("/ * 5").is_ok());
        assert!(parse("/*").is_ok());
        assert!(parse("/ div 5").is_ok());
    }

    #[test]
    fn occurrence_indicator_backtracks() {
        assert!(parse("$a instance of xs:int + 1").is_ok());
        assert!(parse(". treat as item()*").is_ok());
    }

    #[test]
    fn rejects_reserved_names_as_calls() {
        for text in ["if(1)", "map(1)", "array(1)", "item()"] {
            assert!(parse(text).is_err(), "{text}");
        }
    }

    #[test]
    fn accepts_spec_constructs_the_antlr_lexer_could_not() {
        for text in ["a:*", "*:a", "?name", "$m?key", "processing-instruction(x)"] {
            assert!(parse(text).is_ok(), "{text}");
        }
    }

    #[test]
    fn error_names_the_token_and_position() {
        let err = parse("a b").unwrap_err();
        assert_eq!((err.line, err.column), (1, 3));
        assert_eq!(
            err.message,
            "unexpected `b`, expected an operator or end of input"
        );
    }
}
