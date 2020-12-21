package xpath

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type xpathListener struct {
	*BaseXPath31Listener

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

	if symbol.GetTokenType() == XPath31LexerKW_AND || symbol.GetTokenType() == XPath31LexerKW_OR {
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

//func (s *xpathListener) ExitEveryRule(ctx antlr.ParserRuleContext) { trace("Node %s\n", ctx.GetText()) }

// EnterXpath is called when production xpath is entered.
//func (s *xpathListener) EnterXpath(ctx *XpathContext) { trace("Node %s\n", ctx.GetText()) }
//
//func (s *xpathListener) ExitXpath(ctx *XpathContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterParamlist is called when production paramlist is entered.
//func (s *xpathListener) EnterParamlist(ctx *ParamlistContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitParamlist is called when production paramlist is exited.
//func (s *xpathListener) ExitParamlist(ctx *ParamlistContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterParam is called when production param is entered.
//func (s *xpathListener) EnterParam(ctx *ParamContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitParam is called when production param is exited.
//func (s *xpathListener) ExitParam(ctx *ParamContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterFunctionbody is called when production functionbody is entered.
//func (s *xpathListener) EnterFunctionbody(ctx *FunctionbodyContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitFunctionbody is called when production functionbody is exited.
//func (s *xpathListener) ExitFunctionbody(ctx *FunctionbodyContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterEnclosedexpr is called when production enclosedexpr is entered.
//func (s *xpathListener) EnterEnclosedexpr(ctx *EnclosedexprContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitEnclosedexpr is called when production enclosedexpr is exited.
//func (s *xpathListener) ExitEnclosedexpr(ctx *EnclosedexprContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterExpr is called when production expr is entered.
//func (s *xpathListener) EnterExpr(ctx *ExprContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitExpr is called when production expr is exited.
//func (s *xpathListener) ExitExpr(ctx *ExprContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterExprsingle is called when production exprsingle is entered.
//func (s *xpathListener) EnterExprsingle(ctx *ExprsingleContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitExprsingle is called when production exprsingle is exited.
//func (s *xpathListener) ExitExprsingle(ctx *ExprsingleContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterForexpr is called when production forexpr is entered.
//func (s *xpathListener) EnterForexpr(ctx *ForexprContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitForexpr is called when production forexpr is exited.
//func (s *xpathListener) ExitForexpr(ctx *ForexprContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterSimpleforclause is called when production simpleforclause is entered.
//func (s *xpathListener) EnterSimpleforclause(ctx *SimpleforclauseContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitSimpleforclause is called when production simpleforclause is exited.
//func (s *xpathListener) ExitSimpleforclause(ctx *SimpleforclauseContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterSimpleforbinding is called when production simpleforbinding is entered.
//func (s *xpathListener) EnterSimpleforbinding(ctx *SimpleforbindingContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitSimpleforbinding is called when production simpleforbinding is exited.
//func (s *xpathListener) ExitSimpleforbinding(ctx *SimpleforbindingContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterLetexpr is called when production letexpr is entered.
//func (s *xpathListener) EnterLetexpr(ctx *LetexprContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitLetexpr is called when production letexpr is exited.
//func (s *xpathListener) ExitLetexpr(ctx *LetexprContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterSimpleletclause is called when production simpleletclause is entered.
//func (s *xpathListener) EnterSimpleletclause(ctx *SimpleletclauseContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitSimpleletclause is called when production simpleletclause is exited.
//func (s *xpathListener) ExitSimpleletclause(ctx *SimpleletclauseContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterSimpleletbinding is called when production simpleletbinding is entered.
//func (s *xpathListener) EnterSimpleletbinding(ctx *SimpleletbindingContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitSimpleletbinding is called when production simpleletbinding is exited.
//func (s *xpathListener) ExitSimpleletbinding(ctx *SimpleletbindingContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterQuantifiedexpr is called when production quantifiedexpr is entered.
//func (s *xpathListener) EnterQuantifiedexpr(ctx *QuantifiedexprContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitQuantifiedexpr is called when production quantifiedexpr is exited.
//func (s *xpathListener) ExitQuantifiedexpr(ctx *QuantifiedexprContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterIfexpr is called when production ifexpr is entered.
//func (s *xpathListener) EnterIfexpr(ctx *IfexprContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitIfexpr is called when production ifexpr is exited.
//func (s *xpathListener) ExitIfexpr(ctx *IfexprContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterOrexpr is called when production orexpr is entered.
//func (s *xpathListener) EnterOrexpr(ctx *OrexprContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitOrexpr is called when production orexpr is exited.
//func (s *xpathListener) ExitOrexpr(ctx *OrexprContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterAndexpr is called when production andexpr is entered.
//func (s *xpathListener) EnterAndexpr(ctx *AndexprContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitAndexpr is called when production andexpr is exited.
//func (s *xpathListener) ExitAndexpr(ctx *AndexprContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterComparisonexpr is called when production comparisonexpr is entered.
//func (s *xpathListener) EnterComparisonexpr(ctx *ComparisonexprContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitComparisonexpr is called when production comparisonexpr is exited.
//func (s *xpathListener) ExitComparisonexpr(ctx *ComparisonexprContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterStringconcatexpr is called when production stringconcatexpr is entered.
//func (s *xpathListener) EnterStringconcatexpr(ctx *StringconcatexprContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitStringconcatexpr is called when production stringconcatexpr is exited.
//func (s *xpathListener) ExitStringconcatexpr(ctx *StringconcatexprContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterRangeexpr is called when production rangeexpr is entered.
//func (s *xpathListener) EnterRangeexpr(ctx *RangeexprContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitRangeexpr is called when production rangeexpr is exited.
//func (s *xpathListener) ExitRangeexpr(ctx *RangeexprContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterAdditiveexpr is called when production additiveexpr is entered.
//func (s *xpathListener) EnterAdditiveexpr(ctx *AdditiveexprContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitAdditiveexpr is called when production additiveexpr is exited.
//func (s *xpathListener) ExitAdditiveexpr(ctx *AdditiveexprContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterMultiplicativeexpr is called when production multiplicativeexpr is entered.
//func (s *xpathListener) EnterMultiplicativeexpr(ctx *MultiplicativeexprContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitMultiplicativeexpr is called when production multiplicativeexpr is exited.
//func (s *xpathListener) ExitMultiplicativeexpr(ctx *MultiplicativeexprContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterUnionexpr is called when production unionexpr is entered.
//func (s *xpathListener) EnterUnionexpr(ctx *UnionexprContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitUnionexpr is called when production unionexpr is exited.
//func (s *xpathListener) ExitUnionexpr(ctx *UnionexprContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterIntersectexceptexpr is called when production intersectexceptexpr is entered.
//func (s *xpathListener) EnterIntersectexceptexpr(ctx *IntersectexceptexprContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitIntersectexceptexpr is called when production intersectexceptexpr is exited.
//func (s *xpathListener) ExitIntersectexceptexpr(ctx *IntersectexceptexprContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterInstanceofexpr is called when production instanceofexpr is entered.
//func (s *xpathListener) EnterInstanceofexpr(ctx *InstanceofexprContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitInstanceofexpr is called when production instanceofexpr is exited.
//func (s *xpathListener) ExitInstanceofexpr(ctx *InstanceofexprContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterTreatexpr is called when production treatexpr is entered.
//func (s *xpathListener) EnterTreatexpr(ctx *TreatexprContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitTreatexpr is called when production treatexpr is exited.
//func (s *xpathListener) ExitTreatexpr(ctx *TreatexprContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterCastableexpr is called when production castableexpr is entered.
//func (s *xpathListener) EnterCastableexpr(ctx *CastableexprContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitCastableexpr is called when production castableexpr is exited.
//func (s *xpathListener) ExitCastableexpr(ctx *CastableexprContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterCastexpr is called when production castexpr is entered.
//func (s *xpathListener) EnterCastexpr(ctx *CastexprContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitCastexpr is called when production castexpr is exited.
//func (s *xpathListener) ExitCastexpr(ctx *CastexprContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterArrowexpr is called when production arrowexpr is entered.
//func (s *xpathListener) EnterArrowexpr(ctx *ArrowexprContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitArrowexpr is called when production arrowexpr is exited.
//func (s *xpathListener) ExitArrowexpr(ctx *ArrowexprContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterUnaryexpr is called when production unaryexpr is entered.
//func (s *xpathListener) EnterUnaryexpr(ctx *UnaryexprContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitUnaryexpr is called when production unaryexpr is exited.
//func (s *xpathListener) ExitUnaryexpr(ctx *UnaryexprContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterValueexpr is called when production valueexpr is entered.
//func (s *xpathListener) EnterValueexpr(ctx *ValueexprContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitValueexpr is called when production valueexpr is exited.
//func (s *xpathListener) ExitValueexpr(ctx *ValueexprContext) { trace("Node %s\n", ctx.GetText()) }
//

func (s *xpathListener) EnterGeneralcomp(ctx *GeneralcompContext) {
	if s.LastGeneralComp != ctx.GetText() {
		s.GeneralCompChange += 1
		s.LastGeneralComp = ctx.GetText()
	}

	trace("******* Node %s\n", ctx.GetText())
}

//// ExitGeneralcomp is called when production generalcomp is exited.
//func (s *xpathListener) ExitGeneralcomp(ctx *GeneralcompContext) { trace("Node %s\n", ctx.GetText()) }
//
// EnterValuecomp is called when production valuecomp is entered.
func (s *xpathListener) EnterValuecomp(ctx *ValuecompContext) { trace("Node %s\n", ctx.GetText()) }

//// ExitValuecomp is called when production valuecomp is exited.
//func (s *xpathListener) ExitValuecomp(ctx *ValuecompContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterNodecomp is called when production nodecomp is entered.
func (s *xpathListener) EnterNodecomp(ctx *NodecompContext) { trace("Node %s\n", ctx.GetText()) }

//// ExitNodecomp is called when production nodecomp is exited.
//func (s *xpathListener) ExitNodecomp(ctx *NodecompContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterSimplemapexpr is called when production simplemapexpr is entered.
//func (s *xpathListener) EnterSimplemapexpr(ctx *SimplemapexprContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitSimplemapexpr is called when production simplemapexpr is exited.
//func (s *xpathListener) ExitSimplemapexpr(ctx *SimplemapexprContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterPathexpr is called when production pathexpr is entered.
//func (s *xpathListener) EnterPathexpr(ctx *PathexprContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitPathexpr is called when production pathexpr is exited.
//func (s *xpathListener) ExitPathexpr(ctx *PathexprContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterRelativepathexpr is called when production relativepathexpr is entered.
//func (s *xpathListener) EnterRelativepathexpr(ctx *RelativepathexprContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitRelativepathexpr is called when production relativepathexpr is exited.
//func (s *xpathListener) ExitRelativepathexpr(ctx *RelativepathexprContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterStepexpr is called when production stepexpr is entered.
//func (s *xpathListener) EnterStepexpr(ctx *StepexprContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitStepexpr is called when production stepexpr is exited.
//func (s *xpathListener) ExitStepexpr(ctx *StepexprContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterAxisstep is called when production axisstep is entered.
//func (s *xpathListener) EnterAxisstep(ctx *AxisstepContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitAxisstep is called when production axisstep is exited.
//func (s *xpathListener) ExitAxisstep(ctx *AxisstepContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterForwardstep is called when production forwardstep is entered.
//func (s *xpathListener) EnterForwardstep(ctx *ForwardstepContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitForwardstep is called when production forwardstep is exited.
//func (s *xpathListener) ExitForwardstep(ctx *ForwardstepContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterForwardaxis is called when production forwardaxis is entered.
//func (s *xpathListener) EnterForwardaxis(ctx *ForwardaxisContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitForwardaxis is called when production forwardaxis is exited.
//func (s *xpathListener) ExitForwardaxis(ctx *ForwardaxisContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterAbbrevforwardstep is called when production abbrevforwardstep is entered.
//func (s *xpathListener) EnterAbbrevforwardstep(ctx *AbbrevforwardstepContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitAbbrevforwardstep is called when production abbrevforwardstep is exited.
//func (s *xpathListener) ExitAbbrevforwardstep(ctx *AbbrevforwardstepContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterReversestep is called when production reversestep is entered.
//func (s *xpathListener) EnterReversestep(ctx *ReversestepContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitReversestep is called when production reversestep is exited.
//func (s *xpathListener) ExitReversestep(ctx *ReversestepContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterReverseaxis is called when production reverseaxis is entered.
//func (s *xpathListener) EnterReverseaxis(ctx *ReverseaxisContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitReverseaxis is called when production reverseaxis is exited.
//func (s *xpathListener) ExitReverseaxis(ctx *ReverseaxisContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterAbbrevreversestep is called when production abbrevreversestep is entered.
//func (s *xpathListener) EnterAbbrevreversestep(ctx *AbbrevreversestepContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitAbbrevreversestep is called when production abbrevreversestep is exited.
//func (s *xpathListener) ExitAbbrevreversestep(ctx *AbbrevreversestepContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterNodetest is called when production nodetest is entered.
//func (s *xpathListener) EnterNodetest(ctx *NodetestContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitNodetest is called when production nodetest is exited.
//func (s *xpathListener) ExitNodetest(ctx *NodetestContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterNametest is called when production nametest is entered.
//func (s *xpathListener) EnterNametest(ctx *NametestContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitNametest is called when production nametest is exited.
//func (s *xpathListener) ExitNametest(ctx *NametestContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterWildcard is called when production wildcard is entered.
//func (s *xpathListener) EnterWildcard(ctx *WildcardContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitWildcard is called when production wildcard is exited.
//func (s *xpathListener) ExitWildcard(ctx *WildcardContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterPostfixexpr is called when production postfixexpr is entered.
//func (s *xpathListener) EnterPostfixexpr(ctx *PostfixexprContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitPostfixexpr is called when production postfixexpr is exited.
//func (s *xpathListener) ExitPostfixexpr(ctx *PostfixexprContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterArgumentlist is called when production argumentlist is entered.
//func (s *xpathListener) EnterArgumentlist(ctx *ArgumentlistContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitArgumentlist is called when production argumentlist is exited.
//func (s *xpathListener) ExitArgumentlist(ctx *ArgumentlistContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterPredicatelist is called when production predicatelist is entered.
//func (s *xpathListener) EnterPredicatelist(ctx *PredicatelistContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitPredicatelist is called when production predicatelist is exited.
//func (s *xpathListener) ExitPredicatelist(ctx *PredicatelistContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterPredicate is called when production predicate is entered.
//func (s *xpathListener) EnterPredicate(ctx *PredicateContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitPredicate is called when production predicate is exited.
//func (s *xpathListener) ExitPredicate(ctx *PredicateContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterLookup is called when production lookup is entered.
//func (s *xpathListener) EnterLookup(ctx *LookupContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitLookup is called when production lookup is exited.
//func (s *xpathListener) ExitLookup(ctx *LookupContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterKeyspecifier is called when production keyspecifier is entered.
//func (s *xpathListener) EnterKeyspecifier(ctx *KeyspecifierContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitKeyspecifier is called when production keyspecifier is exited.
//func (s *xpathListener) ExitKeyspecifier(ctx *KeyspecifierContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterArrowfunctionspecifier is called when production arrowfunctionspecifier is entered.
//func (s *xpathListener) EnterArrowfunctionspecifier(ctx *ArrowfunctionspecifierContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitArrowfunctionspecifier is called when production arrowfunctionspecifier is exited.
//func (s *xpathListener) ExitArrowfunctionspecifier(ctx *ArrowfunctionspecifierContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterPrimaryexpr is called when production primaryexpr is entered.
//func (s *xpathListener) EnterPrimaryexpr(ctx *PrimaryexprContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitPrimaryexpr is called when production primaryexpr is exited.
//func (s *xpathListener) ExitPrimaryexpr(ctx *PrimaryexprContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterLiteral is called when production literal is entered.
//func (s *xpathListener) EnterLiteral(ctx *LiteralContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitLiteral is called when production literal is exited.
//func (s *xpathListener) ExitLiteral(ctx *LiteralContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterNumericliteral is called when production numericliteral is entered.
//func (s *xpathListener) EnterNumericliteral(ctx *NumericliteralContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitNumericliteral is called when production numericliteral is exited.
//func (s *xpathListener) ExitNumericliteral(ctx *NumericliteralContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterVarref is called when production varref is entered.
//func (s *xpathListener) EnterVarref(ctx *VarrefContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitVarref is called when production varref is exited.
//func (s *xpathListener) ExitVarref(ctx *VarrefContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterVarname is called when production varname is entered.
//func (s *xpathListener) EnterVarname(ctx *VarnameContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitVarname is called when production varname is exited.
//func (s *xpathListener) ExitVarname(ctx *VarnameContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterParenthesizedexpr is called when production parenthesizedexpr is entered.
//func (s *xpathListener) EnterParenthesizedexpr(ctx *ParenthesizedexprContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitParenthesizedexpr is called when production parenthesizedexpr is exited.
//func (s *xpathListener) ExitParenthesizedexpr(ctx *ParenthesizedexprContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterContextitemexpr is called when production contextitemexpr is entered.
//func (s *xpathListener) EnterContextitemexpr(ctx *ContextitemexprContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitContextitemexpr is called when production contextitemexpr is exited.
//func (s *xpathListener) ExitContextitemexpr(ctx *ContextitemexprContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//

func (s *xpathListener) EnterFunctioncall(ctx *FunctioncallContext) {
	s.FunctionCallCount += 1
	trace("Node %s\n", ctx.GetText())
}

//// ExitFunctioncall is called when production functioncall is exited.
//func (s *xpathListener) ExitFunctioncall(ctx *FunctioncallContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterArgument is called when production argument is entered.
//func (s *xpathListener) EnterArgument(ctx *ArgumentContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitArgument is called when production argument is exited.
//func (s *xpathListener) ExitArgument(ctx *ArgumentContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterArgumentplaceholder is called when production argumentplaceholder is entered.
//func (s *xpathListener) EnterArgumentplaceholder(ctx *ArgumentplaceholderContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitArgumentplaceholder is called when production argumentplaceholder is exited.
//func (s *xpathListener) ExitArgumentplaceholder(ctx *ArgumentplaceholderContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterFunctionitemexpr is called when production functionitemexpr is entered.
//func (s *xpathListener) EnterFunctionitemexpr(ctx *FunctionitemexprContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitFunctionitemexpr is called when production functionitemexpr is exited.
//func (s *xpathListener) ExitFunctionitemexpr(ctx *FunctionitemexprContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterNamedfunctionref is called when production namedfunctionref is entered.
//func (s *xpathListener) EnterNamedfunctionref(ctx *NamedfunctionrefContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitNamedfunctionref is called when production namedfunctionref is exited.
//func (s *xpathListener) ExitNamedfunctionref(ctx *NamedfunctionrefContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterInlinefunctionexpr is called when production inlinefunctionexpr is entered.
//func (s *xpathListener) EnterInlinefunctionexpr(ctx *InlinefunctionexprContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitInlinefunctionexpr is called when production inlinefunctionexpr is exited.
//func (s *xpathListener) ExitInlinefunctionexpr(ctx *InlinefunctionexprContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterMapconstructor is called when production mapconstructor is entered.
//func (s *xpathListener) EnterMapconstructor(ctx *MapconstructorContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitMapconstructor is called when production mapconstructor is exited.
//func (s *xpathListener) ExitMapconstructor(ctx *MapconstructorContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterMapconstructorentry is called when production mapconstructorentry is entered.
//func (s *xpathListener) EnterMapconstructorentry(ctx *MapconstructorentryContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitMapconstructorentry is called when production mapconstructorentry is exited.
//func (s *xpathListener) ExitMapconstructorentry(ctx *MapconstructorentryContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterMapkeyexpr is called when production mapkeyexpr is entered.
//func (s *xpathListener) EnterMapkeyexpr(ctx *MapkeyexprContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitMapkeyexpr is called when production mapkeyexpr is exited.
//func (s *xpathListener) ExitMapkeyexpr(ctx *MapkeyexprContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterMapvalueexpr is called when production mapvalueexpr is entered.
//func (s *xpathListener) EnterMapvalueexpr(ctx *MapvalueexprContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitMapvalueexpr is called when production mapvalueexpr is exited.
//func (s *xpathListener) ExitMapvalueexpr(ctx *MapvalueexprContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterArrayconstructor is called when production arrayconstructor is entered.
//func (s *xpathListener) EnterArrayconstructor(ctx *ArrayconstructorContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitArrayconstructor is called when production arrayconstructor is exited.
//func (s *xpathListener) ExitArrayconstructor(ctx *ArrayconstructorContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterSquarearrayconstructor is called when production squarearrayconstructor is entered.
//func (s *xpathListener) EnterSquarearrayconstructor(ctx *SquarearrayconstructorContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitSquarearrayconstructor is called when production squarearrayconstructor is exited.
//func (s *xpathListener) ExitSquarearrayconstructor(ctx *SquarearrayconstructorContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterCurlyarrayconstructor is called when production curlyarrayconstructor is entered.
//func (s *xpathListener) EnterCurlyarrayconstructor(ctx *CurlyarrayconstructorContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitCurlyarrayconstructor is called when production curlyarrayconstructor is exited.
//func (s *xpathListener) ExitCurlyarrayconstructor(ctx *CurlyarrayconstructorContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterUnarylookup is called when production unarylookup is entered.
//func (s *xpathListener) EnterUnarylookup(ctx *UnarylookupContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitUnarylookup is called when production unarylookup is exited.
//func (s *xpathListener) ExitUnarylookup(ctx *UnarylookupContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterSingletype is called when production singletype is entered.
//func (s *xpathListener) EnterSingletype(ctx *SingletypeContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitSingletype is called when production singletype is exited.
//func (s *xpathListener) ExitSingletype(ctx *SingletypeContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterTypedeclaration is called when production typedeclaration is entered.
//func (s *xpathListener) EnterTypedeclaration(ctx *TypedeclarationContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitTypedeclaration is called when production typedeclaration is exited.
//func (s *xpathListener) ExitTypedeclaration(ctx *TypedeclarationContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterSequencetype is called when production sequencetype is entered.
//func (s *xpathListener) EnterSequencetype(ctx *SequencetypeContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitSequencetype is called when production sequencetype is exited.
//func (s *xpathListener) ExitSequencetype(ctx *SequencetypeContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterOccurrenceindicator is called when production occurrenceindicator is entered.
//func (s *xpathListener) EnterOccurrenceindicator(ctx *OccurrenceindicatorContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitOccurrenceindicator is called when production occurrenceindicator is exited.
//func (s *xpathListener) ExitOccurrenceindicator(ctx *OccurrenceindicatorContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterItemtype is called when production itemtype is entered.
//func (s *xpathListener) EnterItemtype(ctx *ItemtypeContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitItemtype is called when production itemtype is exited.
//func (s *xpathListener) ExitItemtype(ctx *ItemtypeContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterAtomicoruniontype is called when production atomicoruniontype is entered.
//func (s *xpathListener) EnterAtomicoruniontype(ctx *AtomicoruniontypeContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitAtomicoruniontype is called when production atomicoruniontype is exited.
//func (s *xpathListener) ExitAtomicoruniontype(ctx *AtomicoruniontypeContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterKindtest is called when production kindtest is entered.
//func (s *xpathListener) EnterKindtest(ctx *KindtestContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitKindtest is called when production kindtest is exited.
//func (s *xpathListener) ExitKindtest(ctx *KindtestContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterAnykindtest is called when production anykindtest is entered.
//func (s *xpathListener) EnterAnykindtest(ctx *AnykindtestContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitAnykindtest is called when production anykindtest is exited.
//func (s *xpathListener) ExitAnykindtest(ctx *AnykindtestContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterDocumenttest is called when production documenttest is entered.
//func (s *xpathListener) EnterDocumenttest(ctx *DocumenttestContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitDocumenttest is called when production documenttest is exited.
//func (s *xpathListener) ExitDocumenttest(ctx *DocumenttestContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterTexttest is called when production texttest is entered.
//func (s *xpathListener) EnterTexttest(ctx *TexttestContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitTexttest is called when production texttest is exited.
//func (s *xpathListener) ExitTexttest(ctx *TexttestContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterCommenttest is called when production commenttest is entered.
//func (s *xpathListener) EnterCommenttest(ctx *CommenttestContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitCommenttest is called when production commenttest is exited.
//func (s *xpathListener) ExitCommenttest(ctx *CommenttestContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterNamespacenodetest is called when production namespacenodetest is entered.
//func (s *xpathListener) EnterNamespacenodetest(ctx *NamespacenodetestContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitNamespacenodetest is called when production namespacenodetest is exited.
//func (s *xpathListener) ExitNamespacenodetest(ctx *NamespacenodetestContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterPitest is called when production pitest is entered.
//func (s *xpathListener) EnterPitest(ctx *PitestContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitPitest is called when production pitest is exited.
//func (s *xpathListener) ExitPitest(ctx *PitestContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterAttributetest is called when production attributetest is entered.
//func (s *xpathListener) EnterAttributetest(ctx *AttributetestContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitAttributetest is called when production attributetest is exited.
//func (s *xpathListener) ExitAttributetest(ctx *AttributetestContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterAttribnameorwildcard is called when production attribnameorwildcard is entered.
//func (s *xpathListener) EnterAttribnameorwildcard(ctx *AttribnameorwildcardContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitAttribnameorwildcard is called when production attribnameorwildcard is exited.
//func (s *xpathListener) ExitAttribnameorwildcard(ctx *AttribnameorwildcardContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterSchemaattributetest is called when production schemaattributetest is entered.
//func (s *xpathListener) EnterSchemaattributetest(ctx *SchemaattributetestContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitSchemaattributetest is called when production schemaattributetest is exited.
//func (s *xpathListener) ExitSchemaattributetest(ctx *SchemaattributetestContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterAttributedeclaration is called when production attributedeclaration is entered.
//func (s *xpathListener) EnterAttributedeclaration(ctx *AttributedeclarationContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitAttributedeclaration is called when production attributedeclaration is exited.
//func (s *xpathListener) ExitAttributedeclaration(ctx *AttributedeclarationContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterElementtest is called when production elementtest is entered.
//func (s *xpathListener) EnterElementtest(ctx *ElementtestContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitElementtest is called when production elementtest is exited.
//func (s *xpathListener) ExitElementtest(ctx *ElementtestContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterElementnameorwildcard is called when production elementnameorwildcard is entered.
//func (s *xpathListener) EnterElementnameorwildcard(ctx *ElementnameorwildcardContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitElementnameorwildcard is called when production elementnameorwildcard is exited.
//func (s *xpathListener) ExitElementnameorwildcard(ctx *ElementnameorwildcardContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterSchemaelementtest is called when production schemaelementtest is entered.
//func (s *xpathListener) EnterSchemaelementtest(ctx *SchemaelementtestContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitSchemaelementtest is called when production schemaelementtest is exited.
//func (s *xpathListener) ExitSchemaelementtest(ctx *SchemaelementtestContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterElementdeclaration is called when production elementdeclaration is entered.
//func (s *xpathListener) EnterElementdeclaration(ctx *ElementdeclarationContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitElementdeclaration is called when production elementdeclaration is exited.
//func (s *xpathListener) ExitElementdeclaration(ctx *ElementdeclarationContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterAttributename is called when production attributename is entered.
//func (s *xpathListener) EnterAttributename(ctx *AttributenameContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitAttributename is called when production attributename is exited.
//func (s *xpathListener) ExitAttributename(ctx *AttributenameContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterElementname is called when production elementname is entered.
//func (s *xpathListener) EnterElementname(ctx *ElementnameContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitElementname is called when production elementname is exited.
//func (s *xpathListener) ExitElementname(ctx *ElementnameContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterSimpletypename is called when production simpletypename is entered.
//func (s *xpathListener) EnterSimpletypename(ctx *SimpletypenameContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitSimpletypename is called when production simpletypename is exited.
//func (s *xpathListener) ExitSimpletypename(ctx *SimpletypenameContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterTypename is called when production typename is entered.
//func (s *xpathListener) EnterTypename(ctx *TypenameContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitTypename is called when production typename is exited.
//func (s *xpathListener) ExitTypename(ctx *TypenameContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterFunctiontest is called when production functiontest is entered.
//func (s *xpathListener) EnterFunctiontest(ctx *FunctiontestContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitFunctiontest is called when production functiontest is exited.
//func (s *xpathListener) ExitFunctiontest(ctx *FunctiontestContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterAnyfunctiontest is called when production anyfunctiontest is entered.
//func (s *xpathListener) EnterAnyfunctiontest(ctx *AnyfunctiontestContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitAnyfunctiontest is called when production anyfunctiontest is exited.
//func (s *xpathListener) ExitAnyfunctiontest(ctx *AnyfunctiontestContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterTypedfunctiontest is called when production typedfunctiontest is entered.
//func (s *xpathListener) EnterTypedfunctiontest(ctx *TypedfunctiontestContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitTypedfunctiontest is called when production typedfunctiontest is exited.
//func (s *xpathListener) ExitTypedfunctiontest(ctx *TypedfunctiontestContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterMaptest is called when production maptest is entered.
//func (s *xpathListener) EnterMaptest(ctx *MaptestContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitMaptest is called when production maptest is exited.
//func (s *xpathListener) ExitMaptest(ctx *MaptestContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterAnymaptest is called when production anymaptest is entered.
//func (s *xpathListener) EnterAnymaptest(ctx *AnymaptestContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitAnymaptest is called when production anymaptest is exited.
//func (s *xpathListener) ExitAnymaptest(ctx *AnymaptestContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterTypedmaptest is called when production typedmaptest is entered.
//func (s *xpathListener) EnterTypedmaptest(ctx *TypedmaptestContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitTypedmaptest is called when production typedmaptest is exited.
//func (s *xpathListener) ExitTypedmaptest(ctx *TypedmaptestContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterArraytest is called when production arraytest is entered.
//func (s *xpathListener) EnterArraytest(ctx *ArraytestContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitArraytest is called when production arraytest is exited.
//func (s *xpathListener) ExitArraytest(ctx *ArraytestContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterAnyarraytest is called when production anyarraytest is entered.
//func (s *xpathListener) EnterAnyarraytest(ctx *AnyarraytestContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitAnyarraytest is called when production anyarraytest is exited.
//func (s *xpathListener) ExitAnyarraytest(ctx *AnyarraytestContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterTypedarraytest is called when production typedarraytest is entered.
//func (s *xpathListener) EnterTypedarraytest(ctx *TypedarraytestContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitTypedarraytest is called when production typedarraytest is exited.
//func (s *xpathListener) ExitTypedarraytest(ctx *TypedarraytestContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterParenthesizeditemtype is called when production parenthesizeditemtype is entered.
//func (s *xpathListener) EnterParenthesizeditemtype(ctx *ParenthesizeditemtypeContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// ExitParenthesizeditemtype is called when production parenthesizeditemtype is exited.
//func (s *xpathListener) ExitParenthesizeditemtype(ctx *ParenthesizeditemtypeContext) {
//	trace("Node %s\n", ctx.GetText())
//}
//
//// EnterEqname is called when production eqname is entered.
func (s *xpathListener) EnterEqname(ctx *EqnameContext) { trace("Node %s\n", ctx.GetText()) }

//// ExitEqname is called when production eqname is exited.
//func (s *xpathListener) ExitEqname(ctx *EqnameContext) { trace("Node %s\n", ctx.GetText()) }
//
//// EnterAuxilary is called when production auxilary is entered.
//func (s *xpathListener) EnterAuxilary(ctx *AuxilaryContext) { trace("Node %s\n", ctx.GetText()) }
//
//// ExitAuxilary is called when production auxilary is exited.
//func (s *xpathListener) ExitAuxilary(ctx *AuxilaryContext) { trace("Node %s\n", ctx.GetText()) }
