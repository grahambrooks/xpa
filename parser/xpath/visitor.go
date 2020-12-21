package xpath

import (
	"fmt"
	"os"
	"runtime"
)

type Visitor struct {
	BaseXPath31Visitor
}

var trace = _trace

func _trace(format string, args ...interface{}) {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		_, _ = fmt.Fprintf(os.Stderr, format, args...)
	}

	fn := runtime.FuncForPC(pc)
	_, _ = fmt.Fprint(os.Stderr, fn.Name())
	_, _ = fmt.Fprintf(os.Stderr, format, args...)
}

func NewXpathVisitor() *Visitor {
	visitor := new(Visitor)
	//visitor.SetSuper(visitor.BaseXPath31Visitor)
	return visitor
}

func (v *Visitor) VisitNametest(ctx *NametestContext) interface{} {
	trace("VisitNameTest %s", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitXpath(ctx *XpathContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitParamlist(ctx *ParamlistContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitParam(ctx *ParamContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitFunctionbody(ctx *FunctionbodyContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitEnclosedexpr(ctx *EnclosedexprContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitExpr(ctx *ExprContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitExprsingle(ctx *ExprsingleContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitForexpr(ctx *ForexprContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitSimpleforclause(ctx *SimpleforclauseContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitSimpleforbinding(ctx *SimpleforbindingContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitLetexpr(ctx *LetexprContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitSimpleletclause(ctx *SimpleletclauseContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitSimpleletbinding(ctx *SimpleletbindingContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitQuantifiedexpr(ctx *QuantifiedexprContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitIfexpr(ctx *IfexprContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitOrexpr(ctx *OrexprContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAndexpr(ctx *AndexprContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitComparisonexpr(ctx *ComparisonexprContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitStringconcatexpr(ctx *StringconcatexprContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitRangeexpr(ctx *RangeexprContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAdditiveexpr(ctx *AdditiveexprContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitMultiplicativeexpr(ctx *MultiplicativeexprContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitUnionexpr(ctx *UnionexprContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitIntersectexceptexpr(ctx *IntersectexceptexprContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitInstanceofexpr(ctx *InstanceofexprContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitTreatexpr(ctx *TreatexprContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitCastableexpr(ctx *CastableexprContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitCastexpr(ctx *CastexprContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitArrowexpr(ctx *ArrowexprContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitUnaryexpr(ctx *UnaryexprContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitValueexpr(ctx *ValueexprContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitGeneralcomp(ctx *GeneralcompContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitValuecomp(ctx *ValuecompContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitNodecomp(ctx *NodecompContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitSimplemapexpr(ctx *SimplemapexprContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPathexpr(ctx *PathexprContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitRelativepathexpr(ctx *RelativepathexprContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitStepexpr(ctx *StepexprContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAxisstep(ctx *AxisstepContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitForwardstep(ctx *ForwardstepContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitForwardaxis(ctx *ForwardaxisContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAbbrevforwardstep(ctx *AbbrevforwardstepContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitReversestep(ctx *ReversestepContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitReverseaxis(ctx *ReverseaxisContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAbbrevreversestep(ctx *AbbrevreversestepContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitNodetest(ctx *NodetestContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitWildcard(ctx *WildcardContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPostfixexpr(ctx *PostfixexprContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitArgumentlist(ctx *ArgumentlistContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPredicatelist(ctx *PredicatelistContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPredicate(ctx *PredicateContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitLookup(ctx *LookupContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitKeyspecifier(ctx *KeyspecifierContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitArrowfunctionspecifier(ctx *ArrowfunctionspecifierContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPrimaryexpr(ctx *PrimaryexprContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitLiteral(ctx *LiteralContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitNumericliteral(ctx *NumericliteralContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitVarref(ctx *VarrefContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitVarname(ctx *VarnameContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitParenthesizedexpr(ctx *ParenthesizedexprContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitContextitemexpr(ctx *ContextitemexprContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitFunctioncall(ctx *FunctioncallContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitArgument(ctx *ArgumentContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitArgumentplaceholder(ctx *ArgumentplaceholderContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitFunctionitemexpr(ctx *FunctionitemexprContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitNamedfunctionref(ctx *NamedfunctionrefContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitInlinefunctionexpr(ctx *InlinefunctionexprContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitMapconstructor(ctx *MapconstructorContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitMapconstructorentry(ctx *MapconstructorentryContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitMapkeyexpr(ctx *MapkeyexprContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitMapvalueexpr(ctx *MapvalueexprContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitArrayconstructor(ctx *ArrayconstructorContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitSquarearrayconstructor(ctx *SquarearrayconstructorContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitCurlyarrayconstructor(ctx *CurlyarrayconstructorContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitUnarylookup(ctx *UnarylookupContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitSingletype(ctx *SingletypeContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitTypedeclaration(ctx *TypedeclarationContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitSequencetype(ctx *SequencetypeContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitOccurrenceindicator(ctx *OccurrenceindicatorContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitItemtype(ctx *ItemtypeContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAtomicoruniontype(ctx *AtomicoruniontypeContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitKindtest(ctx *KindtestContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAnykindtest(ctx *AnykindtestContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitDocumenttest(ctx *DocumenttestContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitTexttest(ctx *TexttestContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitCommenttest(ctx *CommenttestContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitNamespacenodetest(ctx *NamespacenodetestContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPitest(ctx *PitestContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAttributetest(ctx *AttributetestContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAttribnameorwildcard(ctx *AttribnameorwildcardContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitSchemaattributetest(ctx *SchemaattributetestContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAttributedeclaration(ctx *AttributedeclarationContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitElementtest(ctx *ElementtestContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitElementnameorwildcard(ctx *ElementnameorwildcardContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitSchemaelementtest(ctx *SchemaelementtestContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitElementdeclaration(ctx *ElementdeclarationContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAttributename(ctx *AttributenameContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitElementname(ctx *ElementnameContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitSimpletypename(ctx *SimpletypenameContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitTypename(ctx *TypenameContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitFunctiontest(ctx *FunctiontestContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAnyfunctiontest(ctx *AnyfunctiontestContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitTypedfunctiontest(ctx *TypedfunctiontestContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitMaptest(ctx *MaptestContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAnymaptest(ctx *AnymaptestContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitTypedmaptest(ctx *TypedmaptestContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitArraytest(ctx *ArraytestContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAnyarraytest(ctx *AnyarraytestContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitTypedarraytest(ctx *TypedarraytestContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitParenthesizeditemtype(ctx *ParenthesizeditemtypeContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitEqname(ctx *EqnameContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAuxilary(ctx *AuxilaryContext) interface{} {
	trace(" '%s' ", ctx.GetText())
	return v.VisitChildren(ctx)
}
