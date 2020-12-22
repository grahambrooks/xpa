package xpa

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/grahambrooks/xpa/pkg/xpa/XPath31"
	"os"
	"runtime"
	"strings"
)

type Analyser struct{}

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

type CustomErrorHandler struct {
	SyntaxErrors []string
}

func (c *CustomErrorHandler) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	c.SyntaxErrors = append(c.SyntaxErrors, fmt.Sprintf("line %d column %d %s", line, column, msg))
}

func (c CustomErrorHandler) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
}

func (c CustomErrorHandler) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
}

func (c CustomErrorHandler) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
}

func (c *CustomErrorHandler) Error() error {
	if len(c.SyntaxErrors) == 0 {
		return nil
	} else {
		return fmt.Errorf("parser failure %s", strings.Join(c.SyntaxErrors, ","))
	}
}

func (*Analyser) Analyse(text string) (*Analysis, error) {
	trace = func(format string, args ...interface{}) {}
	stream := antlr.NewInputStream(text)
	lexer := XPath31.NewXPath31Lexer(stream)
	tokenstream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := XPath31.NewXPath31Parser(tokenstream)
	errorHandler := CustomErrorHandler{}
	parser.AddErrorListener(&errorHandler)

	listener := xpathListener{}
	antlr.ParseTreeWalkerDefault.Walk(&listener, parser.Xpath())

	return &Analysis{
		TerminalCount:         listener.TerminalCount,
		ComparisonChange:      listener.GeneralCompChange,
		BooleanOperatorChange: listener.BooleanOperatorChange,
		FunctionCallCount:     listener.FunctionCallCount,
		Pattern:               text,
	}, errorHandler.Error()
}

type Analysis struct {
	TerminalCount         int    `json:"terminal_count"`
	ComparisonChange      int    `json:"comparison_change"`
	BooleanOperatorChange int    `json:"boolean_operator_change"`
	FunctionCallCount     int    `json:"function_call_count"`
	Pattern               string `json:"pattern"`
}

func (observation *Analysis) CognitiveComplexity() int {
	return observation.ComparisonChange + observation.FunctionCallCount*2 + observation.BooleanOperatorChange*2
}
