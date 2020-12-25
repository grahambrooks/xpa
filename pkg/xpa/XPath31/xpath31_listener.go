// Code generated from /private/var/tmp/_bazel_graham/081e5e0bb1448a80684d6ec6cd93bc45/sandbox/darwin-sandbox/99/execroot/__main__/pkg/xpa/grammar/XPath31.g4 by ANTLR 4.8. DO NOT EDIT.

package XPath31 // XPath31
import "github.com/antlr/antlr4/runtime/Go/antlr"

// XPath31Listener is a complete listener for a parse tree produced by XPath31Parser.
type XPath31Listener interface {
	antlr.ParseTreeListener

	// EnterXpath is called when entering the xpath production.
	EnterXpath(c *XpathContext)

	// EnterParamlist is called when entering the paramlist production.
	EnterParamlist(c *ParamlistContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterFunctionbody is called when entering the functionbody production.
	EnterFunctionbody(c *FunctionbodyContext)

	// EnterEnclosedexpr is called when entering the enclosedexpr production.
	EnterEnclosedexpr(c *EnclosedexprContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterExprsingle is called when entering the exprsingle production.
	EnterExprsingle(c *ExprsingleContext)

	// EnterForexpr is called when entering the forexpr production.
	EnterForexpr(c *ForexprContext)

	// EnterSimpleforclause is called when entering the simpleforclause production.
	EnterSimpleforclause(c *SimpleforclauseContext)

	// EnterSimpleforbinding is called when entering the simpleforbinding production.
	EnterSimpleforbinding(c *SimpleforbindingContext)

	// EnterLetexpr is called when entering the letexpr production.
	EnterLetexpr(c *LetexprContext)

	// EnterSimpleletclause is called when entering the simpleletclause production.
	EnterSimpleletclause(c *SimpleletclauseContext)

	// EnterSimpleletbinding is called when entering the simpleletbinding production.
	EnterSimpleletbinding(c *SimpleletbindingContext)

	// EnterQuantifiedexpr is called when entering the quantifiedexpr production.
	EnterQuantifiedexpr(c *QuantifiedexprContext)

	// EnterIfexpr is called when entering the ifexpr production.
	EnterIfexpr(c *IfexprContext)

	// EnterOrexpr is called when entering the orexpr production.
	EnterOrexpr(c *OrexprContext)

	// EnterAndexpr is called when entering the andexpr production.
	EnterAndexpr(c *AndexprContext)

	// EnterComparisonexpr is called when entering the comparisonexpr production.
	EnterComparisonexpr(c *ComparisonexprContext)

	// EnterStringconcatexpr is called when entering the stringconcatexpr production.
	EnterStringconcatexpr(c *StringconcatexprContext)

	// EnterRangeexpr is called when entering the rangeexpr production.
	EnterRangeexpr(c *RangeexprContext)

	// EnterAdditiveexpr is called when entering the additiveexpr production.
	EnterAdditiveexpr(c *AdditiveexprContext)

	// EnterMultiplicativeexpr is called when entering the multiplicativeexpr production.
	EnterMultiplicativeexpr(c *MultiplicativeexprContext)

	// EnterUnionexpr is called when entering the unionexpr production.
	EnterUnionexpr(c *UnionexprContext)

	// EnterIntersectexceptexpr is called when entering the intersectexceptexpr production.
	EnterIntersectexceptexpr(c *IntersectexceptexprContext)

	// EnterInstanceofexpr is called when entering the instanceofexpr production.
	EnterInstanceofexpr(c *InstanceofexprContext)

	// EnterTreatexpr is called when entering the treatexpr production.
	EnterTreatexpr(c *TreatexprContext)

	// EnterCastableexpr is called when entering the castableexpr production.
	EnterCastableexpr(c *CastableexprContext)

	// EnterCastexpr is called when entering the castexpr production.
	EnterCastexpr(c *CastexprContext)

	// EnterArrowexpr is called when entering the arrowexpr production.
	EnterArrowexpr(c *ArrowexprContext)

	// EnterUnaryexpr is called when entering the unaryexpr production.
	EnterUnaryexpr(c *UnaryexprContext)

	// EnterValueexpr is called when entering the valueexpr production.
	EnterValueexpr(c *ValueexprContext)

	// EnterGeneralcomp is called when entering the generalcomp production.
	EnterGeneralcomp(c *GeneralcompContext)

	// EnterValuecomp is called when entering the valuecomp production.
	EnterValuecomp(c *ValuecompContext)

	// EnterNodecomp is called when entering the nodecomp production.
	EnterNodecomp(c *NodecompContext)

	// EnterSimplemapexpr is called when entering the simplemapexpr production.
	EnterSimplemapexpr(c *SimplemapexprContext)

	// EnterPathexpr is called when entering the pathexpr production.
	EnterPathexpr(c *PathexprContext)

	// EnterRelativepathexpr is called when entering the relativepathexpr production.
	EnterRelativepathexpr(c *RelativepathexprContext)

	// EnterStepexpr is called when entering the stepexpr production.
	EnterStepexpr(c *StepexprContext)

	// EnterAxisstep is called when entering the axisstep production.
	EnterAxisstep(c *AxisstepContext)

	// EnterForwardstep is called when entering the forwardstep production.
	EnterForwardstep(c *ForwardstepContext)

	// EnterForwardaxis is called when entering the forwardaxis production.
	EnterForwardaxis(c *ForwardaxisContext)

	// EnterAbbrevforwardstep is called when entering the abbrevforwardstep production.
	EnterAbbrevforwardstep(c *AbbrevforwardstepContext)

	// EnterReversestep is called when entering the reversestep production.
	EnterReversestep(c *ReversestepContext)

	// EnterReverseaxis is called when entering the reverseaxis production.
	EnterReverseaxis(c *ReverseaxisContext)

	// EnterAbbrevreversestep is called when entering the abbrevreversestep production.
	EnterAbbrevreversestep(c *AbbrevreversestepContext)

	// EnterNodetest is called when entering the nodetest production.
	EnterNodetest(c *NodetestContext)

	// EnterNametest is called when entering the nametest production.
	EnterNametest(c *NametestContext)

	// EnterWildcard is called when entering the wildcard production.
	EnterWildcard(c *WildcardContext)

	// EnterPostfixexpr is called when entering the postfixexpr production.
	EnterPostfixexpr(c *PostfixexprContext)

	// EnterArgumentlist is called when entering the argumentlist production.
	EnterArgumentlist(c *ArgumentlistContext)

	// EnterPredicatelist is called when entering the predicatelist production.
	EnterPredicatelist(c *PredicatelistContext)

	// EnterPredicate is called when entering the predicate production.
	EnterPredicate(c *PredicateContext)

	// EnterLookup is called when entering the lookup production.
	EnterLookup(c *LookupContext)

	// EnterKeyspecifier is called when entering the keyspecifier production.
	EnterKeyspecifier(c *KeyspecifierContext)

	// EnterArrowfunctionspecifier is called when entering the arrowfunctionspecifier production.
	EnterArrowfunctionspecifier(c *ArrowfunctionspecifierContext)

	// EnterPrimaryexpr is called when entering the primaryexpr production.
	EnterPrimaryexpr(c *PrimaryexprContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterNumericliteral is called when entering the numericliteral production.
	EnterNumericliteral(c *NumericliteralContext)

	// EnterVarref is called when entering the varref production.
	EnterVarref(c *VarrefContext)

	// EnterVarname is called when entering the varname production.
	EnterVarname(c *VarnameContext)

	// EnterParenthesizedexpr is called when entering the parenthesizedexpr production.
	EnterParenthesizedexpr(c *ParenthesizedexprContext)

	// EnterContextitemexpr is called when entering the contextitemexpr production.
	EnterContextitemexpr(c *ContextitemexprContext)

	// EnterFunctioncall is called when entering the functioncall production.
	EnterFunctioncall(c *FunctioncallContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterArgumentplaceholder is called when entering the argumentplaceholder production.
	EnterArgumentplaceholder(c *ArgumentplaceholderContext)

	// EnterFunctionitemexpr is called when entering the functionitemexpr production.
	EnterFunctionitemexpr(c *FunctionitemexprContext)

	// EnterNamedfunctionref is called when entering the namedfunctionref production.
	EnterNamedfunctionref(c *NamedfunctionrefContext)

	// EnterInlinefunctionexpr is called when entering the inlinefunctionexpr production.
	EnterInlinefunctionexpr(c *InlinefunctionexprContext)

	// EnterMapconstructor is called when entering the mapconstructor production.
	EnterMapconstructor(c *MapconstructorContext)

	// EnterMapconstructorentry is called when entering the mapconstructorentry production.
	EnterMapconstructorentry(c *MapconstructorentryContext)

	// EnterMapkeyexpr is called when entering the mapkeyexpr production.
	EnterMapkeyexpr(c *MapkeyexprContext)

	// EnterMapvalueexpr is called when entering the mapvalueexpr production.
	EnterMapvalueexpr(c *MapvalueexprContext)

	// EnterArrayconstructor is called when entering the arrayconstructor production.
	EnterArrayconstructor(c *ArrayconstructorContext)

	// EnterSquarearrayconstructor is called when entering the squarearrayconstructor production.
	EnterSquarearrayconstructor(c *SquarearrayconstructorContext)

	// EnterCurlyarrayconstructor is called when entering the curlyarrayconstructor production.
	EnterCurlyarrayconstructor(c *CurlyarrayconstructorContext)

	// EnterUnarylookup is called when entering the unarylookup production.
	EnterUnarylookup(c *UnarylookupContext)

	// EnterSingletype is called when entering the singletype production.
	EnterSingletype(c *SingletypeContext)

	// EnterTypedeclaration is called when entering the typedeclaration production.
	EnterTypedeclaration(c *TypedeclarationContext)

	// EnterSequencetype is called when entering the sequencetype production.
	EnterSequencetype(c *SequencetypeContext)

	// EnterOccurrenceindicator is called when entering the occurrenceindicator production.
	EnterOccurrenceindicator(c *OccurrenceindicatorContext)

	// EnterItemtype is called when entering the itemtype production.
	EnterItemtype(c *ItemtypeContext)

	// EnterAtomicoruniontype is called when entering the atomicoruniontype production.
	EnterAtomicoruniontype(c *AtomicoruniontypeContext)

	// EnterKindtest is called when entering the kindtest production.
	EnterKindtest(c *KindtestContext)

	// EnterAnykindtest is called when entering the anykindtest production.
	EnterAnykindtest(c *AnykindtestContext)

	// EnterDocumenttest is called when entering the documenttest production.
	EnterDocumenttest(c *DocumenttestContext)

	// EnterTexttest is called when entering the texttest production.
	EnterTexttest(c *TexttestContext)

	// EnterCommenttest is called when entering the commenttest production.
	EnterCommenttest(c *CommenttestContext)

	// EnterNamespacenodetest is called when entering the namespacenodetest production.
	EnterNamespacenodetest(c *NamespacenodetestContext)

	// EnterPitest is called when entering the pitest production.
	EnterPitest(c *PitestContext)

	// EnterAttributetest is called when entering the attributetest production.
	EnterAttributetest(c *AttributetestContext)

	// EnterAttribnameorwildcard is called when entering the attribnameorwildcard production.
	EnterAttribnameorwildcard(c *AttribnameorwildcardContext)

	// EnterSchemaattributetest is called when entering the schemaattributetest production.
	EnterSchemaattributetest(c *SchemaattributetestContext)

	// EnterAttributedeclaration is called when entering the attributedeclaration production.
	EnterAttributedeclaration(c *AttributedeclarationContext)

	// EnterElementtest is called when entering the elementtest production.
	EnterElementtest(c *ElementtestContext)

	// EnterElementnameorwildcard is called when entering the elementnameorwildcard production.
	EnterElementnameorwildcard(c *ElementnameorwildcardContext)

	// EnterSchemaelementtest is called when entering the schemaelementtest production.
	EnterSchemaelementtest(c *SchemaelementtestContext)

	// EnterElementdeclaration is called when entering the elementdeclaration production.
	EnterElementdeclaration(c *ElementdeclarationContext)

	// EnterAttributename is called when entering the attributename production.
	EnterAttributename(c *AttributenameContext)

	// EnterElementname is called when entering the elementname production.
	EnterElementname(c *ElementnameContext)

	// EnterSimpletypename is called when entering the simpletypename production.
	EnterSimpletypename(c *SimpletypenameContext)

	// EnterTypename is called when entering the typename production.
	EnterTypename(c *TypenameContext)

	// EnterFunctiontest is called when entering the functiontest production.
	EnterFunctiontest(c *FunctiontestContext)

	// EnterAnyfunctiontest is called when entering the anyfunctiontest production.
	EnterAnyfunctiontest(c *AnyfunctiontestContext)

	// EnterTypedfunctiontest is called when entering the typedfunctiontest production.
	EnterTypedfunctiontest(c *TypedfunctiontestContext)

	// EnterMaptest is called when entering the maptest production.
	EnterMaptest(c *MaptestContext)

	// EnterAnymaptest is called when entering the anymaptest production.
	EnterAnymaptest(c *AnymaptestContext)

	// EnterTypedmaptest is called when entering the typedmaptest production.
	EnterTypedmaptest(c *TypedmaptestContext)

	// EnterArraytest is called when entering the arraytest production.
	EnterArraytest(c *ArraytestContext)

	// EnterAnyarraytest is called when entering the anyarraytest production.
	EnterAnyarraytest(c *AnyarraytestContext)

	// EnterTypedarraytest is called when entering the typedarraytest production.
	EnterTypedarraytest(c *TypedarraytestContext)

	// EnterParenthesizeditemtype is called when entering the parenthesizeditemtype production.
	EnterParenthesizeditemtype(c *ParenthesizeditemtypeContext)

	// EnterEqname is called when entering the eqname production.
	EnterEqname(c *EqnameContext)

	// EnterAuxilary is called when entering the auxilary production.
	EnterAuxilary(c *AuxilaryContext)

	// ExitXpath is called when exiting the xpath production.
	ExitXpath(c *XpathContext)

	// ExitParamlist is called when exiting the paramlist production.
	ExitParamlist(c *ParamlistContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitFunctionbody is called when exiting the functionbody production.
	ExitFunctionbody(c *FunctionbodyContext)

	// ExitEnclosedexpr is called when exiting the enclosedexpr production.
	ExitEnclosedexpr(c *EnclosedexprContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitExprsingle is called when exiting the exprsingle production.
	ExitExprsingle(c *ExprsingleContext)

	// ExitForexpr is called when exiting the forexpr production.
	ExitForexpr(c *ForexprContext)

	// ExitSimpleforclause is called when exiting the simpleforclause production.
	ExitSimpleforclause(c *SimpleforclauseContext)

	// ExitSimpleforbinding is called when exiting the simpleforbinding production.
	ExitSimpleforbinding(c *SimpleforbindingContext)

	// ExitLetexpr is called when exiting the letexpr production.
	ExitLetexpr(c *LetexprContext)

	// ExitSimpleletclause is called when exiting the simpleletclause production.
	ExitSimpleletclause(c *SimpleletclauseContext)

	// ExitSimpleletbinding is called when exiting the simpleletbinding production.
	ExitSimpleletbinding(c *SimpleletbindingContext)

	// ExitQuantifiedexpr is called when exiting the quantifiedexpr production.
	ExitQuantifiedexpr(c *QuantifiedexprContext)

	// ExitIfexpr is called when exiting the ifexpr production.
	ExitIfexpr(c *IfexprContext)

	// ExitOrexpr is called when exiting the orexpr production.
	ExitOrexpr(c *OrexprContext)

	// ExitAndexpr is called when exiting the andexpr production.
	ExitAndexpr(c *AndexprContext)

	// ExitComparisonexpr is called when exiting the comparisonexpr production.
	ExitComparisonexpr(c *ComparisonexprContext)

	// ExitStringconcatexpr is called when exiting the stringconcatexpr production.
	ExitStringconcatexpr(c *StringconcatexprContext)

	// ExitRangeexpr is called when exiting the rangeexpr production.
	ExitRangeexpr(c *RangeexprContext)

	// ExitAdditiveexpr is called when exiting the additiveexpr production.
	ExitAdditiveexpr(c *AdditiveexprContext)

	// ExitMultiplicativeexpr is called when exiting the multiplicativeexpr production.
	ExitMultiplicativeexpr(c *MultiplicativeexprContext)

	// ExitUnionexpr is called when exiting the unionexpr production.
	ExitUnionexpr(c *UnionexprContext)

	// ExitIntersectexceptexpr is called when exiting the intersectexceptexpr production.
	ExitIntersectexceptexpr(c *IntersectexceptexprContext)

	// ExitInstanceofexpr is called when exiting the instanceofexpr production.
	ExitInstanceofexpr(c *InstanceofexprContext)

	// ExitTreatexpr is called when exiting the treatexpr production.
	ExitTreatexpr(c *TreatexprContext)

	// ExitCastableexpr is called when exiting the castableexpr production.
	ExitCastableexpr(c *CastableexprContext)

	// ExitCastexpr is called when exiting the castexpr production.
	ExitCastexpr(c *CastexprContext)

	// ExitArrowexpr is called when exiting the arrowexpr production.
	ExitArrowexpr(c *ArrowexprContext)

	// ExitUnaryexpr is called when exiting the unaryexpr production.
	ExitUnaryexpr(c *UnaryexprContext)

	// ExitValueexpr is called when exiting the valueexpr production.
	ExitValueexpr(c *ValueexprContext)

	// ExitGeneralcomp is called when exiting the generalcomp production.
	ExitGeneralcomp(c *GeneralcompContext)

	// ExitValuecomp is called when exiting the valuecomp production.
	ExitValuecomp(c *ValuecompContext)

	// ExitNodecomp is called when exiting the nodecomp production.
	ExitNodecomp(c *NodecompContext)

	// ExitSimplemapexpr is called when exiting the simplemapexpr production.
	ExitSimplemapexpr(c *SimplemapexprContext)

	// ExitPathexpr is called when exiting the pathexpr production.
	ExitPathexpr(c *PathexprContext)

	// ExitRelativepathexpr is called when exiting the relativepathexpr production.
	ExitRelativepathexpr(c *RelativepathexprContext)

	// ExitStepexpr is called when exiting the stepexpr production.
	ExitStepexpr(c *StepexprContext)

	// ExitAxisstep is called when exiting the axisstep production.
	ExitAxisstep(c *AxisstepContext)

	// ExitForwardstep is called when exiting the forwardstep production.
	ExitForwardstep(c *ForwardstepContext)

	// ExitForwardaxis is called when exiting the forwardaxis production.
	ExitForwardaxis(c *ForwardaxisContext)

	// ExitAbbrevforwardstep is called when exiting the abbrevforwardstep production.
	ExitAbbrevforwardstep(c *AbbrevforwardstepContext)

	// ExitReversestep is called when exiting the reversestep production.
	ExitReversestep(c *ReversestepContext)

	// ExitReverseaxis is called when exiting the reverseaxis production.
	ExitReverseaxis(c *ReverseaxisContext)

	// ExitAbbrevreversestep is called when exiting the abbrevreversestep production.
	ExitAbbrevreversestep(c *AbbrevreversestepContext)

	// ExitNodetest is called when exiting the nodetest production.
	ExitNodetest(c *NodetestContext)

	// ExitNametest is called when exiting the nametest production.
	ExitNametest(c *NametestContext)

	// ExitWildcard is called when exiting the wildcard production.
	ExitWildcard(c *WildcardContext)

	// ExitPostfixexpr is called when exiting the postfixexpr production.
	ExitPostfixexpr(c *PostfixexprContext)

	// ExitArgumentlist is called when exiting the argumentlist production.
	ExitArgumentlist(c *ArgumentlistContext)

	// ExitPredicatelist is called when exiting the predicatelist production.
	ExitPredicatelist(c *PredicatelistContext)

	// ExitPredicate is called when exiting the predicate production.
	ExitPredicate(c *PredicateContext)

	// ExitLookup is called when exiting the lookup production.
	ExitLookup(c *LookupContext)

	// ExitKeyspecifier is called when exiting the keyspecifier production.
	ExitKeyspecifier(c *KeyspecifierContext)

	// ExitArrowfunctionspecifier is called when exiting the arrowfunctionspecifier production.
	ExitArrowfunctionspecifier(c *ArrowfunctionspecifierContext)

	// ExitPrimaryexpr is called when exiting the primaryexpr production.
	ExitPrimaryexpr(c *PrimaryexprContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitNumericliteral is called when exiting the numericliteral production.
	ExitNumericliteral(c *NumericliteralContext)

	// ExitVarref is called when exiting the varref production.
	ExitVarref(c *VarrefContext)

	// ExitVarname is called when exiting the varname production.
	ExitVarname(c *VarnameContext)

	// ExitParenthesizedexpr is called when exiting the parenthesizedexpr production.
	ExitParenthesizedexpr(c *ParenthesizedexprContext)

	// ExitContextitemexpr is called when exiting the contextitemexpr production.
	ExitContextitemexpr(c *ContextitemexprContext)

	// ExitFunctioncall is called when exiting the functioncall production.
	ExitFunctioncall(c *FunctioncallContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitArgumentplaceholder is called when exiting the argumentplaceholder production.
	ExitArgumentplaceholder(c *ArgumentplaceholderContext)

	// ExitFunctionitemexpr is called when exiting the functionitemexpr production.
	ExitFunctionitemexpr(c *FunctionitemexprContext)

	// ExitNamedfunctionref is called when exiting the namedfunctionref production.
	ExitNamedfunctionref(c *NamedfunctionrefContext)

	// ExitInlinefunctionexpr is called when exiting the inlinefunctionexpr production.
	ExitInlinefunctionexpr(c *InlinefunctionexprContext)

	// ExitMapconstructor is called when exiting the mapconstructor production.
	ExitMapconstructor(c *MapconstructorContext)

	// ExitMapconstructorentry is called when exiting the mapconstructorentry production.
	ExitMapconstructorentry(c *MapconstructorentryContext)

	// ExitMapkeyexpr is called when exiting the mapkeyexpr production.
	ExitMapkeyexpr(c *MapkeyexprContext)

	// ExitMapvalueexpr is called when exiting the mapvalueexpr production.
	ExitMapvalueexpr(c *MapvalueexprContext)

	// ExitArrayconstructor is called when exiting the arrayconstructor production.
	ExitArrayconstructor(c *ArrayconstructorContext)

	// ExitSquarearrayconstructor is called when exiting the squarearrayconstructor production.
	ExitSquarearrayconstructor(c *SquarearrayconstructorContext)

	// ExitCurlyarrayconstructor is called when exiting the curlyarrayconstructor production.
	ExitCurlyarrayconstructor(c *CurlyarrayconstructorContext)

	// ExitUnarylookup is called when exiting the unarylookup production.
	ExitUnarylookup(c *UnarylookupContext)

	// ExitSingletype is called when exiting the singletype production.
	ExitSingletype(c *SingletypeContext)

	// ExitTypedeclaration is called when exiting the typedeclaration production.
	ExitTypedeclaration(c *TypedeclarationContext)

	// ExitSequencetype is called when exiting the sequencetype production.
	ExitSequencetype(c *SequencetypeContext)

	// ExitOccurrenceindicator is called when exiting the occurrenceindicator production.
	ExitOccurrenceindicator(c *OccurrenceindicatorContext)

	// ExitItemtype is called when exiting the itemtype production.
	ExitItemtype(c *ItemtypeContext)

	// ExitAtomicoruniontype is called when exiting the atomicoruniontype production.
	ExitAtomicoruniontype(c *AtomicoruniontypeContext)

	// ExitKindtest is called when exiting the kindtest production.
	ExitKindtest(c *KindtestContext)

	// ExitAnykindtest is called when exiting the anykindtest production.
	ExitAnykindtest(c *AnykindtestContext)

	// ExitDocumenttest is called when exiting the documenttest production.
	ExitDocumenttest(c *DocumenttestContext)

	// ExitTexttest is called when exiting the texttest production.
	ExitTexttest(c *TexttestContext)

	// ExitCommenttest is called when exiting the commenttest production.
	ExitCommenttest(c *CommenttestContext)

	// ExitNamespacenodetest is called when exiting the namespacenodetest production.
	ExitNamespacenodetest(c *NamespacenodetestContext)

	// ExitPitest is called when exiting the pitest production.
	ExitPitest(c *PitestContext)

	// ExitAttributetest is called when exiting the attributetest production.
	ExitAttributetest(c *AttributetestContext)

	// ExitAttribnameorwildcard is called when exiting the attribnameorwildcard production.
	ExitAttribnameorwildcard(c *AttribnameorwildcardContext)

	// ExitSchemaattributetest is called when exiting the schemaattributetest production.
	ExitSchemaattributetest(c *SchemaattributetestContext)

	// ExitAttributedeclaration is called when exiting the attributedeclaration production.
	ExitAttributedeclaration(c *AttributedeclarationContext)

	// ExitElementtest is called when exiting the elementtest production.
	ExitElementtest(c *ElementtestContext)

	// ExitElementnameorwildcard is called when exiting the elementnameorwildcard production.
	ExitElementnameorwildcard(c *ElementnameorwildcardContext)

	// ExitSchemaelementtest is called when exiting the schemaelementtest production.
	ExitSchemaelementtest(c *SchemaelementtestContext)

	// ExitElementdeclaration is called when exiting the elementdeclaration production.
	ExitElementdeclaration(c *ElementdeclarationContext)

	// ExitAttributename is called when exiting the attributename production.
	ExitAttributename(c *AttributenameContext)

	// ExitElementname is called when exiting the elementname production.
	ExitElementname(c *ElementnameContext)

	// ExitSimpletypename is called when exiting the simpletypename production.
	ExitSimpletypename(c *SimpletypenameContext)

	// ExitTypename is called when exiting the typename production.
	ExitTypename(c *TypenameContext)

	// ExitFunctiontest is called when exiting the functiontest production.
	ExitFunctiontest(c *FunctiontestContext)

	// ExitAnyfunctiontest is called when exiting the anyfunctiontest production.
	ExitAnyfunctiontest(c *AnyfunctiontestContext)

	// ExitTypedfunctiontest is called when exiting the typedfunctiontest production.
	ExitTypedfunctiontest(c *TypedfunctiontestContext)

	// ExitMaptest is called when exiting the maptest production.
	ExitMaptest(c *MaptestContext)

	// ExitAnymaptest is called when exiting the anymaptest production.
	ExitAnymaptest(c *AnymaptestContext)

	// ExitTypedmaptest is called when exiting the typedmaptest production.
	ExitTypedmaptest(c *TypedmaptestContext)

	// ExitArraytest is called when exiting the arraytest production.
	ExitArraytest(c *ArraytestContext)

	// ExitAnyarraytest is called when exiting the anyarraytest production.
	ExitAnyarraytest(c *AnyarraytestContext)

	// ExitTypedarraytest is called when exiting the typedarraytest production.
	ExitTypedarraytest(c *TypedarraytestContext)

	// ExitParenthesizeditemtype is called when exiting the parenthesizeditemtype production.
	ExitParenthesizeditemtype(c *ParenthesizeditemtypeContext)

	// ExitEqname is called when exiting the eqname production.
	ExitEqname(c *EqnameContext)

	// ExitAuxilary is called when exiting the auxilary production.
	ExitAuxilary(c *AuxilaryContext)
}
