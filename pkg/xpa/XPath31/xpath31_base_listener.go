// Code generated from /private/var/tmp/_bazel_graham/081e5e0bb1448a80684d6ec6cd93bc45/sandbox/darwin-sandbox/99/execroot/__main__/pkg/xpa/grammar/XPath31.g4 by ANTLR 4.8. DO NOT EDIT.

package XPath31 // XPath31
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseXPath31Listener is a complete listener for a parse tree produced by XPath31Parser.
type BaseXPath31Listener struct{}

var _ XPath31Listener = &BaseXPath31Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseXPath31Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseXPath31Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseXPath31Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseXPath31Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterXpath is called when production xpath is entered.
func (s *BaseXPath31Listener) EnterXpath(ctx *XpathContext) {}

// ExitXpath is called when production xpath is exited.
func (s *BaseXPath31Listener) ExitXpath(ctx *XpathContext) {}

// EnterParamlist is called when production paramlist is entered.
func (s *BaseXPath31Listener) EnterParamlist(ctx *ParamlistContext) {}

// ExitParamlist is called when production paramlist is exited.
func (s *BaseXPath31Listener) ExitParamlist(ctx *ParamlistContext) {}

// EnterParam is called when production param is entered.
func (s *BaseXPath31Listener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BaseXPath31Listener) ExitParam(ctx *ParamContext) {}

// EnterFunctionbody is called when production functionbody is entered.
func (s *BaseXPath31Listener) EnterFunctionbody(ctx *FunctionbodyContext) {}

// ExitFunctionbody is called when production functionbody is exited.
func (s *BaseXPath31Listener) ExitFunctionbody(ctx *FunctionbodyContext) {}

// EnterEnclosedexpr is called when production enclosedexpr is entered.
func (s *BaseXPath31Listener) EnterEnclosedexpr(ctx *EnclosedexprContext) {}

// ExitEnclosedexpr is called when production enclosedexpr is exited.
func (s *BaseXPath31Listener) ExitEnclosedexpr(ctx *EnclosedexprContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseXPath31Listener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseXPath31Listener) ExitExpr(ctx *ExprContext) {}

// EnterExprsingle is called when production exprsingle is entered.
func (s *BaseXPath31Listener) EnterExprsingle(ctx *ExprsingleContext) {}

// ExitExprsingle is called when production exprsingle is exited.
func (s *BaseXPath31Listener) ExitExprsingle(ctx *ExprsingleContext) {}

// EnterForexpr is called when production forexpr is entered.
func (s *BaseXPath31Listener) EnterForexpr(ctx *ForexprContext) {}

// ExitForexpr is called when production forexpr is exited.
func (s *BaseXPath31Listener) ExitForexpr(ctx *ForexprContext) {}

// EnterSimpleforclause is called when production simpleforclause is entered.
func (s *BaseXPath31Listener) EnterSimpleforclause(ctx *SimpleforclauseContext) {}

// ExitSimpleforclause is called when production simpleforclause is exited.
func (s *BaseXPath31Listener) ExitSimpleforclause(ctx *SimpleforclauseContext) {}

// EnterSimpleforbinding is called when production simpleforbinding is entered.
func (s *BaseXPath31Listener) EnterSimpleforbinding(ctx *SimpleforbindingContext) {}

// ExitSimpleforbinding is called when production simpleforbinding is exited.
func (s *BaseXPath31Listener) ExitSimpleforbinding(ctx *SimpleforbindingContext) {}

// EnterLetexpr is called when production letexpr is entered.
func (s *BaseXPath31Listener) EnterLetexpr(ctx *LetexprContext) {}

// ExitLetexpr is called when production letexpr is exited.
func (s *BaseXPath31Listener) ExitLetexpr(ctx *LetexprContext) {}

// EnterSimpleletclause is called when production simpleletclause is entered.
func (s *BaseXPath31Listener) EnterSimpleletclause(ctx *SimpleletclauseContext) {}

// ExitSimpleletclause is called when production simpleletclause is exited.
func (s *BaseXPath31Listener) ExitSimpleletclause(ctx *SimpleletclauseContext) {}

// EnterSimpleletbinding is called when production simpleletbinding is entered.
func (s *BaseXPath31Listener) EnterSimpleletbinding(ctx *SimpleletbindingContext) {}

// ExitSimpleletbinding is called when production simpleletbinding is exited.
func (s *BaseXPath31Listener) ExitSimpleletbinding(ctx *SimpleletbindingContext) {}

// EnterQuantifiedexpr is called when production quantifiedexpr is entered.
func (s *BaseXPath31Listener) EnterQuantifiedexpr(ctx *QuantifiedexprContext) {}

// ExitQuantifiedexpr is called when production quantifiedexpr is exited.
func (s *BaseXPath31Listener) ExitQuantifiedexpr(ctx *QuantifiedexprContext) {}

// EnterIfexpr is called when production ifexpr is entered.
func (s *BaseXPath31Listener) EnterIfexpr(ctx *IfexprContext) {}

// ExitIfexpr is called when production ifexpr is exited.
func (s *BaseXPath31Listener) ExitIfexpr(ctx *IfexprContext) {}

// EnterOrexpr is called when production orexpr is entered.
func (s *BaseXPath31Listener) EnterOrexpr(ctx *OrexprContext) {}

// ExitOrexpr is called when production orexpr is exited.
func (s *BaseXPath31Listener) ExitOrexpr(ctx *OrexprContext) {}

// EnterAndexpr is called when production andexpr is entered.
func (s *BaseXPath31Listener) EnterAndexpr(ctx *AndexprContext) {}

// ExitAndexpr is called when production andexpr is exited.
func (s *BaseXPath31Listener) ExitAndexpr(ctx *AndexprContext) {}

// EnterComparisonexpr is called when production comparisonexpr is entered.
func (s *BaseXPath31Listener) EnterComparisonexpr(ctx *ComparisonexprContext) {}

// ExitComparisonexpr is called when production comparisonexpr is exited.
func (s *BaseXPath31Listener) ExitComparisonexpr(ctx *ComparisonexprContext) {}

// EnterStringconcatexpr is called when production stringconcatexpr is entered.
func (s *BaseXPath31Listener) EnterStringconcatexpr(ctx *StringconcatexprContext) {}

// ExitStringconcatexpr is called when production stringconcatexpr is exited.
func (s *BaseXPath31Listener) ExitStringconcatexpr(ctx *StringconcatexprContext) {}

// EnterRangeexpr is called when production rangeexpr is entered.
func (s *BaseXPath31Listener) EnterRangeexpr(ctx *RangeexprContext) {}

// ExitRangeexpr is called when production rangeexpr is exited.
func (s *BaseXPath31Listener) ExitRangeexpr(ctx *RangeexprContext) {}

// EnterAdditiveexpr is called when production additiveexpr is entered.
func (s *BaseXPath31Listener) EnterAdditiveexpr(ctx *AdditiveexprContext) {}

// ExitAdditiveexpr is called when production additiveexpr is exited.
func (s *BaseXPath31Listener) ExitAdditiveexpr(ctx *AdditiveexprContext) {}

// EnterMultiplicativeexpr is called when production multiplicativeexpr is entered.
func (s *BaseXPath31Listener) EnterMultiplicativeexpr(ctx *MultiplicativeexprContext) {}

// ExitMultiplicativeexpr is called when production multiplicativeexpr is exited.
func (s *BaseXPath31Listener) ExitMultiplicativeexpr(ctx *MultiplicativeexprContext) {}

// EnterUnionexpr is called when production unionexpr is entered.
func (s *BaseXPath31Listener) EnterUnionexpr(ctx *UnionexprContext) {}

// ExitUnionexpr is called when production unionexpr is exited.
func (s *BaseXPath31Listener) ExitUnionexpr(ctx *UnionexprContext) {}

// EnterIntersectexceptexpr is called when production intersectexceptexpr is entered.
func (s *BaseXPath31Listener) EnterIntersectexceptexpr(ctx *IntersectexceptexprContext) {}

// ExitIntersectexceptexpr is called when production intersectexceptexpr is exited.
func (s *BaseXPath31Listener) ExitIntersectexceptexpr(ctx *IntersectexceptexprContext) {}

// EnterInstanceofexpr is called when production instanceofexpr is entered.
func (s *BaseXPath31Listener) EnterInstanceofexpr(ctx *InstanceofexprContext) {}

// ExitInstanceofexpr is called when production instanceofexpr is exited.
func (s *BaseXPath31Listener) ExitInstanceofexpr(ctx *InstanceofexprContext) {}

// EnterTreatexpr is called when production treatexpr is entered.
func (s *BaseXPath31Listener) EnterTreatexpr(ctx *TreatexprContext) {}

// ExitTreatexpr is called when production treatexpr is exited.
func (s *BaseXPath31Listener) ExitTreatexpr(ctx *TreatexprContext) {}

// EnterCastableexpr is called when production castableexpr is entered.
func (s *BaseXPath31Listener) EnterCastableexpr(ctx *CastableexprContext) {}

// ExitCastableexpr is called when production castableexpr is exited.
func (s *BaseXPath31Listener) ExitCastableexpr(ctx *CastableexprContext) {}

// EnterCastexpr is called when production castexpr is entered.
func (s *BaseXPath31Listener) EnterCastexpr(ctx *CastexprContext) {}

// ExitCastexpr is called when production castexpr is exited.
func (s *BaseXPath31Listener) ExitCastexpr(ctx *CastexprContext) {}

// EnterArrowexpr is called when production arrowexpr is entered.
func (s *BaseXPath31Listener) EnterArrowexpr(ctx *ArrowexprContext) {}

// ExitArrowexpr is called when production arrowexpr is exited.
func (s *BaseXPath31Listener) ExitArrowexpr(ctx *ArrowexprContext) {}

// EnterUnaryexpr is called when production unaryexpr is entered.
func (s *BaseXPath31Listener) EnterUnaryexpr(ctx *UnaryexprContext) {}

// ExitUnaryexpr is called when production unaryexpr is exited.
func (s *BaseXPath31Listener) ExitUnaryexpr(ctx *UnaryexprContext) {}

// EnterValueexpr is called when production valueexpr is entered.
func (s *BaseXPath31Listener) EnterValueexpr(ctx *ValueexprContext) {}

// ExitValueexpr is called when production valueexpr is exited.
func (s *BaseXPath31Listener) ExitValueexpr(ctx *ValueexprContext) {}

// EnterGeneralcomp is called when production generalcomp is entered.
func (s *BaseXPath31Listener) EnterGeneralcomp(ctx *GeneralcompContext) {}

// ExitGeneralcomp is called when production generalcomp is exited.
func (s *BaseXPath31Listener) ExitGeneralcomp(ctx *GeneralcompContext) {}

// EnterValuecomp is called when production valuecomp is entered.
func (s *BaseXPath31Listener) EnterValuecomp(ctx *ValuecompContext) {}

// ExitValuecomp is called when production valuecomp is exited.
func (s *BaseXPath31Listener) ExitValuecomp(ctx *ValuecompContext) {}

// EnterNodecomp is called when production nodecomp is entered.
func (s *BaseXPath31Listener) EnterNodecomp(ctx *NodecompContext) {}

// ExitNodecomp is called when production nodecomp is exited.
func (s *BaseXPath31Listener) ExitNodecomp(ctx *NodecompContext) {}

// EnterSimplemapexpr is called when production simplemapexpr is entered.
func (s *BaseXPath31Listener) EnterSimplemapexpr(ctx *SimplemapexprContext) {}

// ExitSimplemapexpr is called when production simplemapexpr is exited.
func (s *BaseXPath31Listener) ExitSimplemapexpr(ctx *SimplemapexprContext) {}

// EnterPathexpr is called when production pathexpr is entered.
func (s *BaseXPath31Listener) EnterPathexpr(ctx *PathexprContext) {}

// ExitPathexpr is called when production pathexpr is exited.
func (s *BaseXPath31Listener) ExitPathexpr(ctx *PathexprContext) {}

// EnterRelativepathexpr is called when production relativepathexpr is entered.
func (s *BaseXPath31Listener) EnterRelativepathexpr(ctx *RelativepathexprContext) {}

// ExitRelativepathexpr is called when production relativepathexpr is exited.
func (s *BaseXPath31Listener) ExitRelativepathexpr(ctx *RelativepathexprContext) {}

// EnterStepexpr is called when production stepexpr is entered.
func (s *BaseXPath31Listener) EnterStepexpr(ctx *StepexprContext) {}

// ExitStepexpr is called when production stepexpr is exited.
func (s *BaseXPath31Listener) ExitStepexpr(ctx *StepexprContext) {}

// EnterAxisstep is called when production axisstep is entered.
func (s *BaseXPath31Listener) EnterAxisstep(ctx *AxisstepContext) {}

// ExitAxisstep is called when production axisstep is exited.
func (s *BaseXPath31Listener) ExitAxisstep(ctx *AxisstepContext) {}

// EnterForwardstep is called when production forwardstep is entered.
func (s *BaseXPath31Listener) EnterForwardstep(ctx *ForwardstepContext) {}

// ExitForwardstep is called when production forwardstep is exited.
func (s *BaseXPath31Listener) ExitForwardstep(ctx *ForwardstepContext) {}

// EnterForwardaxis is called when production forwardaxis is entered.
func (s *BaseXPath31Listener) EnterForwardaxis(ctx *ForwardaxisContext) {}

// ExitForwardaxis is called when production forwardaxis is exited.
func (s *BaseXPath31Listener) ExitForwardaxis(ctx *ForwardaxisContext) {}

// EnterAbbrevforwardstep is called when production abbrevforwardstep is entered.
func (s *BaseXPath31Listener) EnterAbbrevforwardstep(ctx *AbbrevforwardstepContext) {}

// ExitAbbrevforwardstep is called when production abbrevforwardstep is exited.
func (s *BaseXPath31Listener) ExitAbbrevforwardstep(ctx *AbbrevforwardstepContext) {}

// EnterReversestep is called when production reversestep is entered.
func (s *BaseXPath31Listener) EnterReversestep(ctx *ReversestepContext) {}

// ExitReversestep is called when production reversestep is exited.
func (s *BaseXPath31Listener) ExitReversestep(ctx *ReversestepContext) {}

// EnterReverseaxis is called when production reverseaxis is entered.
func (s *BaseXPath31Listener) EnterReverseaxis(ctx *ReverseaxisContext) {}

// ExitReverseaxis is called when production reverseaxis is exited.
func (s *BaseXPath31Listener) ExitReverseaxis(ctx *ReverseaxisContext) {}

// EnterAbbrevreversestep is called when production abbrevreversestep is entered.
func (s *BaseXPath31Listener) EnterAbbrevreversestep(ctx *AbbrevreversestepContext) {}

// ExitAbbrevreversestep is called when production abbrevreversestep is exited.
func (s *BaseXPath31Listener) ExitAbbrevreversestep(ctx *AbbrevreversestepContext) {}

// EnterNodetest is called when production nodetest is entered.
func (s *BaseXPath31Listener) EnterNodetest(ctx *NodetestContext) {}

// ExitNodetest is called when production nodetest is exited.
func (s *BaseXPath31Listener) ExitNodetest(ctx *NodetestContext) {}

// EnterNametest is called when production nametest is entered.
func (s *BaseXPath31Listener) EnterNametest(ctx *NametestContext) {}

// ExitNametest is called when production nametest is exited.
func (s *BaseXPath31Listener) ExitNametest(ctx *NametestContext) {}

// EnterWildcard is called when production wildcard is entered.
func (s *BaseXPath31Listener) EnterWildcard(ctx *WildcardContext) {}

// ExitWildcard is called when production wildcard is exited.
func (s *BaseXPath31Listener) ExitWildcard(ctx *WildcardContext) {}

// EnterPostfixexpr is called when production postfixexpr is entered.
func (s *BaseXPath31Listener) EnterPostfixexpr(ctx *PostfixexprContext) {}

// ExitPostfixexpr is called when production postfixexpr is exited.
func (s *BaseXPath31Listener) ExitPostfixexpr(ctx *PostfixexprContext) {}

// EnterArgumentlist is called when production argumentlist is entered.
func (s *BaseXPath31Listener) EnterArgumentlist(ctx *ArgumentlistContext) {}

// ExitArgumentlist is called when production argumentlist is exited.
func (s *BaseXPath31Listener) ExitArgumentlist(ctx *ArgumentlistContext) {}

// EnterPredicatelist is called when production predicatelist is entered.
func (s *BaseXPath31Listener) EnterPredicatelist(ctx *PredicatelistContext) {}

// ExitPredicatelist is called when production predicatelist is exited.
func (s *BaseXPath31Listener) ExitPredicatelist(ctx *PredicatelistContext) {}

// EnterPredicate is called when production predicate is entered.
func (s *BaseXPath31Listener) EnterPredicate(ctx *PredicateContext) {}

// ExitPredicate is called when production predicate is exited.
func (s *BaseXPath31Listener) ExitPredicate(ctx *PredicateContext) {}

// EnterLookup is called when production lookup is entered.
func (s *BaseXPath31Listener) EnterLookup(ctx *LookupContext) {}

// ExitLookup is called when production lookup is exited.
func (s *BaseXPath31Listener) ExitLookup(ctx *LookupContext) {}

// EnterKeyspecifier is called when production keyspecifier is entered.
func (s *BaseXPath31Listener) EnterKeyspecifier(ctx *KeyspecifierContext) {}

// ExitKeyspecifier is called when production keyspecifier is exited.
func (s *BaseXPath31Listener) ExitKeyspecifier(ctx *KeyspecifierContext) {}

// EnterArrowfunctionspecifier is called when production arrowfunctionspecifier is entered.
func (s *BaseXPath31Listener) EnterArrowfunctionspecifier(ctx *ArrowfunctionspecifierContext) {}

// ExitArrowfunctionspecifier is called when production arrowfunctionspecifier is exited.
func (s *BaseXPath31Listener) ExitArrowfunctionspecifier(ctx *ArrowfunctionspecifierContext) {}

// EnterPrimaryexpr is called when production primaryexpr is entered.
func (s *BaseXPath31Listener) EnterPrimaryexpr(ctx *PrimaryexprContext) {}

// ExitPrimaryexpr is called when production primaryexpr is exited.
func (s *BaseXPath31Listener) ExitPrimaryexpr(ctx *PrimaryexprContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseXPath31Listener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseXPath31Listener) ExitLiteral(ctx *LiteralContext) {}

// EnterNumericliteral is called when production numericliteral is entered.
func (s *BaseXPath31Listener) EnterNumericliteral(ctx *NumericliteralContext) {}

// ExitNumericliteral is called when production numericliteral is exited.
func (s *BaseXPath31Listener) ExitNumericliteral(ctx *NumericliteralContext) {}

// EnterVarref is called when production varref is entered.
func (s *BaseXPath31Listener) EnterVarref(ctx *VarrefContext) {}

// ExitVarref is called when production varref is exited.
func (s *BaseXPath31Listener) ExitVarref(ctx *VarrefContext) {}

// EnterVarname is called when production varname is entered.
func (s *BaseXPath31Listener) EnterVarname(ctx *VarnameContext) {}

// ExitVarname is called when production varname is exited.
func (s *BaseXPath31Listener) ExitVarname(ctx *VarnameContext) {}

// EnterParenthesizedexpr is called when production parenthesizedexpr is entered.
func (s *BaseXPath31Listener) EnterParenthesizedexpr(ctx *ParenthesizedexprContext) {}

// ExitParenthesizedexpr is called when production parenthesizedexpr is exited.
func (s *BaseXPath31Listener) ExitParenthesizedexpr(ctx *ParenthesizedexprContext) {}

// EnterContextitemexpr is called when production contextitemexpr is entered.
func (s *BaseXPath31Listener) EnterContextitemexpr(ctx *ContextitemexprContext) {}

// ExitContextitemexpr is called when production contextitemexpr is exited.
func (s *BaseXPath31Listener) ExitContextitemexpr(ctx *ContextitemexprContext) {}

// EnterFunctioncall is called when production functioncall is entered.
func (s *BaseXPath31Listener) EnterFunctioncall(ctx *FunctioncallContext) {}

// ExitFunctioncall is called when production functioncall is exited.
func (s *BaseXPath31Listener) ExitFunctioncall(ctx *FunctioncallContext) {}

// EnterArgument is called when production argument is entered.
func (s *BaseXPath31Listener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BaseXPath31Listener) ExitArgument(ctx *ArgumentContext) {}

// EnterArgumentplaceholder is called when production argumentplaceholder is entered.
func (s *BaseXPath31Listener) EnterArgumentplaceholder(ctx *ArgumentplaceholderContext) {}

// ExitArgumentplaceholder is called when production argumentplaceholder is exited.
func (s *BaseXPath31Listener) ExitArgumentplaceholder(ctx *ArgumentplaceholderContext) {}

// EnterFunctionitemexpr is called when production functionitemexpr is entered.
func (s *BaseXPath31Listener) EnterFunctionitemexpr(ctx *FunctionitemexprContext) {}

// ExitFunctionitemexpr is called when production functionitemexpr is exited.
func (s *BaseXPath31Listener) ExitFunctionitemexpr(ctx *FunctionitemexprContext) {}

// EnterNamedfunctionref is called when production namedfunctionref is entered.
func (s *BaseXPath31Listener) EnterNamedfunctionref(ctx *NamedfunctionrefContext) {}

// ExitNamedfunctionref is called when production namedfunctionref is exited.
func (s *BaseXPath31Listener) ExitNamedfunctionref(ctx *NamedfunctionrefContext) {}

// EnterInlinefunctionexpr is called when production inlinefunctionexpr is entered.
func (s *BaseXPath31Listener) EnterInlinefunctionexpr(ctx *InlinefunctionexprContext) {}

// ExitInlinefunctionexpr is called when production inlinefunctionexpr is exited.
func (s *BaseXPath31Listener) ExitInlinefunctionexpr(ctx *InlinefunctionexprContext) {}

// EnterMapconstructor is called when production mapconstructor is entered.
func (s *BaseXPath31Listener) EnterMapconstructor(ctx *MapconstructorContext) {}

// ExitMapconstructor is called when production mapconstructor is exited.
func (s *BaseXPath31Listener) ExitMapconstructor(ctx *MapconstructorContext) {}

// EnterMapconstructorentry is called when production mapconstructorentry is entered.
func (s *BaseXPath31Listener) EnterMapconstructorentry(ctx *MapconstructorentryContext) {}

// ExitMapconstructorentry is called when production mapconstructorentry is exited.
func (s *BaseXPath31Listener) ExitMapconstructorentry(ctx *MapconstructorentryContext) {}

// EnterMapkeyexpr is called when production mapkeyexpr is entered.
func (s *BaseXPath31Listener) EnterMapkeyexpr(ctx *MapkeyexprContext) {}

// ExitMapkeyexpr is called when production mapkeyexpr is exited.
func (s *BaseXPath31Listener) ExitMapkeyexpr(ctx *MapkeyexprContext) {}

// EnterMapvalueexpr is called when production mapvalueexpr is entered.
func (s *BaseXPath31Listener) EnterMapvalueexpr(ctx *MapvalueexprContext) {}

// ExitMapvalueexpr is called when production mapvalueexpr is exited.
func (s *BaseXPath31Listener) ExitMapvalueexpr(ctx *MapvalueexprContext) {}

// EnterArrayconstructor is called when production arrayconstructor is entered.
func (s *BaseXPath31Listener) EnterArrayconstructor(ctx *ArrayconstructorContext) {}

// ExitArrayconstructor is called when production arrayconstructor is exited.
func (s *BaseXPath31Listener) ExitArrayconstructor(ctx *ArrayconstructorContext) {}

// EnterSquarearrayconstructor is called when production squarearrayconstructor is entered.
func (s *BaseXPath31Listener) EnterSquarearrayconstructor(ctx *SquarearrayconstructorContext) {}

// ExitSquarearrayconstructor is called when production squarearrayconstructor is exited.
func (s *BaseXPath31Listener) ExitSquarearrayconstructor(ctx *SquarearrayconstructorContext) {}

// EnterCurlyarrayconstructor is called when production curlyarrayconstructor is entered.
func (s *BaseXPath31Listener) EnterCurlyarrayconstructor(ctx *CurlyarrayconstructorContext) {}

// ExitCurlyarrayconstructor is called when production curlyarrayconstructor is exited.
func (s *BaseXPath31Listener) ExitCurlyarrayconstructor(ctx *CurlyarrayconstructorContext) {}

// EnterUnarylookup is called when production unarylookup is entered.
func (s *BaseXPath31Listener) EnterUnarylookup(ctx *UnarylookupContext) {}

// ExitUnarylookup is called when production unarylookup is exited.
func (s *BaseXPath31Listener) ExitUnarylookup(ctx *UnarylookupContext) {}

// EnterSingletype is called when production singletype is entered.
func (s *BaseXPath31Listener) EnterSingletype(ctx *SingletypeContext) {}

// ExitSingletype is called when production singletype is exited.
func (s *BaseXPath31Listener) ExitSingletype(ctx *SingletypeContext) {}

// EnterTypedeclaration is called when production typedeclaration is entered.
func (s *BaseXPath31Listener) EnterTypedeclaration(ctx *TypedeclarationContext) {}

// ExitTypedeclaration is called when production typedeclaration is exited.
func (s *BaseXPath31Listener) ExitTypedeclaration(ctx *TypedeclarationContext) {}

// EnterSequencetype is called when production sequencetype is entered.
func (s *BaseXPath31Listener) EnterSequencetype(ctx *SequencetypeContext) {}

// ExitSequencetype is called when production sequencetype is exited.
func (s *BaseXPath31Listener) ExitSequencetype(ctx *SequencetypeContext) {}

// EnterOccurrenceindicator is called when production occurrenceindicator is entered.
func (s *BaseXPath31Listener) EnterOccurrenceindicator(ctx *OccurrenceindicatorContext) {}

// ExitOccurrenceindicator is called when production occurrenceindicator is exited.
func (s *BaseXPath31Listener) ExitOccurrenceindicator(ctx *OccurrenceindicatorContext) {}

// EnterItemtype is called when production itemtype is entered.
func (s *BaseXPath31Listener) EnterItemtype(ctx *ItemtypeContext) {}

// ExitItemtype is called when production itemtype is exited.
func (s *BaseXPath31Listener) ExitItemtype(ctx *ItemtypeContext) {}

// EnterAtomicoruniontype is called when production atomicoruniontype is entered.
func (s *BaseXPath31Listener) EnterAtomicoruniontype(ctx *AtomicoruniontypeContext) {}

// ExitAtomicoruniontype is called when production atomicoruniontype is exited.
func (s *BaseXPath31Listener) ExitAtomicoruniontype(ctx *AtomicoruniontypeContext) {}

// EnterKindtest is called when production kindtest is entered.
func (s *BaseXPath31Listener) EnterKindtest(ctx *KindtestContext) {}

// ExitKindtest is called when production kindtest is exited.
func (s *BaseXPath31Listener) ExitKindtest(ctx *KindtestContext) {}

// EnterAnykindtest is called when production anykindtest is entered.
func (s *BaseXPath31Listener) EnterAnykindtest(ctx *AnykindtestContext) {}

// ExitAnykindtest is called when production anykindtest is exited.
func (s *BaseXPath31Listener) ExitAnykindtest(ctx *AnykindtestContext) {}

// EnterDocumenttest is called when production documenttest is entered.
func (s *BaseXPath31Listener) EnterDocumenttest(ctx *DocumenttestContext) {}

// ExitDocumenttest is called when production documenttest is exited.
func (s *BaseXPath31Listener) ExitDocumenttest(ctx *DocumenttestContext) {}

// EnterTexttest is called when production texttest is entered.
func (s *BaseXPath31Listener) EnterTexttest(ctx *TexttestContext) {}

// ExitTexttest is called when production texttest is exited.
func (s *BaseXPath31Listener) ExitTexttest(ctx *TexttestContext) {}

// EnterCommenttest is called when production commenttest is entered.
func (s *BaseXPath31Listener) EnterCommenttest(ctx *CommenttestContext) {}

// ExitCommenttest is called when production commenttest is exited.
func (s *BaseXPath31Listener) ExitCommenttest(ctx *CommenttestContext) {}

// EnterNamespacenodetest is called when production namespacenodetest is entered.
func (s *BaseXPath31Listener) EnterNamespacenodetest(ctx *NamespacenodetestContext) {}

// ExitNamespacenodetest is called when production namespacenodetest is exited.
func (s *BaseXPath31Listener) ExitNamespacenodetest(ctx *NamespacenodetestContext) {}

// EnterPitest is called when production pitest is entered.
func (s *BaseXPath31Listener) EnterPitest(ctx *PitestContext) {}

// ExitPitest is called when production pitest is exited.
func (s *BaseXPath31Listener) ExitPitest(ctx *PitestContext) {}

// EnterAttributetest is called when production attributetest is entered.
func (s *BaseXPath31Listener) EnterAttributetest(ctx *AttributetestContext) {}

// ExitAttributetest is called when production attributetest is exited.
func (s *BaseXPath31Listener) ExitAttributetest(ctx *AttributetestContext) {}

// EnterAttribnameorwildcard is called when production attribnameorwildcard is entered.
func (s *BaseXPath31Listener) EnterAttribnameorwildcard(ctx *AttribnameorwildcardContext) {}

// ExitAttribnameorwildcard is called when production attribnameorwildcard is exited.
func (s *BaseXPath31Listener) ExitAttribnameorwildcard(ctx *AttribnameorwildcardContext) {}

// EnterSchemaattributetest is called when production schemaattributetest is entered.
func (s *BaseXPath31Listener) EnterSchemaattributetest(ctx *SchemaattributetestContext) {}

// ExitSchemaattributetest is called when production schemaattributetest is exited.
func (s *BaseXPath31Listener) ExitSchemaattributetest(ctx *SchemaattributetestContext) {}

// EnterAttributedeclaration is called when production attributedeclaration is entered.
func (s *BaseXPath31Listener) EnterAttributedeclaration(ctx *AttributedeclarationContext) {}

// ExitAttributedeclaration is called when production attributedeclaration is exited.
func (s *BaseXPath31Listener) ExitAttributedeclaration(ctx *AttributedeclarationContext) {}

// EnterElementtest is called when production elementtest is entered.
func (s *BaseXPath31Listener) EnterElementtest(ctx *ElementtestContext) {}

// ExitElementtest is called when production elementtest is exited.
func (s *BaseXPath31Listener) ExitElementtest(ctx *ElementtestContext) {}

// EnterElementnameorwildcard is called when production elementnameorwildcard is entered.
func (s *BaseXPath31Listener) EnterElementnameorwildcard(ctx *ElementnameorwildcardContext) {}

// ExitElementnameorwildcard is called when production elementnameorwildcard is exited.
func (s *BaseXPath31Listener) ExitElementnameorwildcard(ctx *ElementnameorwildcardContext) {}

// EnterSchemaelementtest is called when production schemaelementtest is entered.
func (s *BaseXPath31Listener) EnterSchemaelementtest(ctx *SchemaelementtestContext) {}

// ExitSchemaelementtest is called when production schemaelementtest is exited.
func (s *BaseXPath31Listener) ExitSchemaelementtest(ctx *SchemaelementtestContext) {}

// EnterElementdeclaration is called when production elementdeclaration is entered.
func (s *BaseXPath31Listener) EnterElementdeclaration(ctx *ElementdeclarationContext) {}

// ExitElementdeclaration is called when production elementdeclaration is exited.
func (s *BaseXPath31Listener) ExitElementdeclaration(ctx *ElementdeclarationContext) {}

// EnterAttributename is called when production attributename is entered.
func (s *BaseXPath31Listener) EnterAttributename(ctx *AttributenameContext) {}

// ExitAttributename is called when production attributename is exited.
func (s *BaseXPath31Listener) ExitAttributename(ctx *AttributenameContext) {}

// EnterElementname is called when production elementname is entered.
func (s *BaseXPath31Listener) EnterElementname(ctx *ElementnameContext) {}

// ExitElementname is called when production elementname is exited.
func (s *BaseXPath31Listener) ExitElementname(ctx *ElementnameContext) {}

// EnterSimpletypename is called when production simpletypename is entered.
func (s *BaseXPath31Listener) EnterSimpletypename(ctx *SimpletypenameContext) {}

// ExitSimpletypename is called when production simpletypename is exited.
func (s *BaseXPath31Listener) ExitSimpletypename(ctx *SimpletypenameContext) {}

// EnterTypename is called when production typename is entered.
func (s *BaseXPath31Listener) EnterTypename(ctx *TypenameContext) {}

// ExitTypename is called when production typename is exited.
func (s *BaseXPath31Listener) ExitTypename(ctx *TypenameContext) {}

// EnterFunctiontest is called when production functiontest is entered.
func (s *BaseXPath31Listener) EnterFunctiontest(ctx *FunctiontestContext) {}

// ExitFunctiontest is called when production functiontest is exited.
func (s *BaseXPath31Listener) ExitFunctiontest(ctx *FunctiontestContext) {}

// EnterAnyfunctiontest is called when production anyfunctiontest is entered.
func (s *BaseXPath31Listener) EnterAnyfunctiontest(ctx *AnyfunctiontestContext) {}

// ExitAnyfunctiontest is called when production anyfunctiontest is exited.
func (s *BaseXPath31Listener) ExitAnyfunctiontest(ctx *AnyfunctiontestContext) {}

// EnterTypedfunctiontest is called when production typedfunctiontest is entered.
func (s *BaseXPath31Listener) EnterTypedfunctiontest(ctx *TypedfunctiontestContext) {}

// ExitTypedfunctiontest is called when production typedfunctiontest is exited.
func (s *BaseXPath31Listener) ExitTypedfunctiontest(ctx *TypedfunctiontestContext) {}

// EnterMaptest is called when production maptest is entered.
func (s *BaseXPath31Listener) EnterMaptest(ctx *MaptestContext) {}

// ExitMaptest is called when production maptest is exited.
func (s *BaseXPath31Listener) ExitMaptest(ctx *MaptestContext) {}

// EnterAnymaptest is called when production anymaptest is entered.
func (s *BaseXPath31Listener) EnterAnymaptest(ctx *AnymaptestContext) {}

// ExitAnymaptest is called when production anymaptest is exited.
func (s *BaseXPath31Listener) ExitAnymaptest(ctx *AnymaptestContext) {}

// EnterTypedmaptest is called when production typedmaptest is entered.
func (s *BaseXPath31Listener) EnterTypedmaptest(ctx *TypedmaptestContext) {}

// ExitTypedmaptest is called when production typedmaptest is exited.
func (s *BaseXPath31Listener) ExitTypedmaptest(ctx *TypedmaptestContext) {}

// EnterArraytest is called when production arraytest is entered.
func (s *BaseXPath31Listener) EnterArraytest(ctx *ArraytestContext) {}

// ExitArraytest is called when production arraytest is exited.
func (s *BaseXPath31Listener) ExitArraytest(ctx *ArraytestContext) {}

// EnterAnyarraytest is called when production anyarraytest is entered.
func (s *BaseXPath31Listener) EnterAnyarraytest(ctx *AnyarraytestContext) {}

// ExitAnyarraytest is called when production anyarraytest is exited.
func (s *BaseXPath31Listener) ExitAnyarraytest(ctx *AnyarraytestContext) {}

// EnterTypedarraytest is called when production typedarraytest is entered.
func (s *BaseXPath31Listener) EnterTypedarraytest(ctx *TypedarraytestContext) {}

// ExitTypedarraytest is called when production typedarraytest is exited.
func (s *BaseXPath31Listener) ExitTypedarraytest(ctx *TypedarraytestContext) {}

// EnterParenthesizeditemtype is called when production parenthesizeditemtype is entered.
func (s *BaseXPath31Listener) EnterParenthesizeditemtype(ctx *ParenthesizeditemtypeContext) {}

// ExitParenthesizeditemtype is called when production parenthesizeditemtype is exited.
func (s *BaseXPath31Listener) ExitParenthesizeditemtype(ctx *ParenthesizeditemtypeContext) {}

// EnterEqname is called when production eqname is entered.
func (s *BaseXPath31Listener) EnterEqname(ctx *EqnameContext) {}

// ExitEqname is called when production eqname is exited.
func (s *BaseXPath31Listener) ExitEqname(ctx *EqnameContext) {}

// EnterAuxilary is called when production auxilary is entered.
func (s *BaseXPath31Listener) EnterAuxilary(ctx *AuxilaryContext) {}

// ExitAuxilary is called when production auxilary is exited.
func (s *BaseXPath31Listener) ExitAuxilary(ctx *AuxilaryContext) {}
