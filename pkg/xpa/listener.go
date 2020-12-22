package xpa

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/grahambrooks/xpa/pkg/xpa/XPath31"
)

type xpathListener struct {
	*XPath31.BaseXPath31Listener

	TerminalCount         int
	LastGeneralComp       string
	GeneralCompChange     int
	LastBooleanOperator   int
	BooleanOperatorChange int
	FunctionCallCount     int
}

func (s *xpathListener) VisitTerminal(node antlr.TerminalNode) {
	s.TerminalCount += 1
	symbol := node.GetSymbol()

	if symbol.GetTokenType() == XPath31.XPath31LexerKW_AND || symbol.GetTokenType() ==  XPath31.XPath31LexerKW_OR {
		if s.LastBooleanOperator != symbol.GetTokenType() {
			s.BooleanOperatorChange += 1
			s.LastBooleanOperator = symbol.GetTokenType()
		}
	}
	trace("Symbol type: %d, line: %d, col: %d, txt: %s index: %d\n", symbol.GetTokenType(), symbol.GetLine(), symbol.GetColumn(), symbol.GetText(), symbol.GetTokenIndex())
}

func (s *xpathListener) VisitErrorNode(node antlr.ErrorNode) {
	trace("Node %s\n", node.GetText())
	if symbol := node.GetSymbol(); symbol != nil {
		trace("Symbol type: %d, line: %d, col: %d, txt: %s index: %d\n", symbol.GetTokenType(), symbol.GetLine(), symbol.GetColumn(), symbol.GetText(), symbol.GetTokenIndex())
	} else {
		trace("No symbol in node")
	}
}

func (s *xpathListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	trace("Node %s\n", ctx.GetText())
}


func (s *xpathListener) EnterGeneralcomp(ctx *XPath31.GeneralcompContext) {
	if s.LastGeneralComp != ctx.GetText() {
		s.GeneralCompChange += 1
		s.LastGeneralComp = ctx.GetText()
	}

	trace("******* Node %s\n", ctx.GetText())
}


func (s *xpathListener) EnterValuecomp(ctx *XPath31.ValuecompContext) { trace("Node %s\n", ctx.GetText()) }


func (s *xpathListener) EnterNodecomp(ctx *XPath31.NodecompContext) { trace("Node %s\n", ctx.GetText()) }


func (s *xpathListener) EnterFunctioncall(ctx *XPath31.FunctioncallContext) {
	s.FunctionCallCount += 1
	trace("Node %s\n", ctx.GetText())
}


func (s *xpathListener) EnterEqname(ctx *XPath31.EqnameContext) { trace("Node %s\n", ctx.GetText()) }

