package xpath

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Parser struct{}

func (*Parser) Parse(text string) {
	trace = func(format string, args ...interface{}) {}
	stream := antlr.NewInputStream(text)
	lexer := NewXPath31Lexer(stream)
	tokenstream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := NewXPath31Parser(tokenstream)

	//tree := parser.Xpath()

	listener := xpathListener{}
	antlr.ParseTreeWalkerDefault.Walk(&listener, parser.Xpath())

	if CognitiveComplexity(&Observation{
		TerminalCount:         listener.TerminalCount,
		GeneralCompChange:     listener.GeneralCompChange,
		BooleanOperatorChange: listener.BooleanOperatorChange,
		FunctionCallCount:     listener.FunctionCallCount,
	}) > 10 {
		fmt.Printf("boolean complexity %d comparison complexity %d terminals %d function calls %d %s\n", listener.BooleanOperatorChange, listener.GeneralCompChange, listener.TerminalCount, listener.FunctionCallCount, text)
	}

}

type Observation struct {
	TerminalCount         int
	GeneralCompChange     int
	BooleanOperatorChange int
	FunctionCallCount     int
}

func CognitiveComplexity(observations *Observation) int {
	return observations.GeneralCompChange + observations.FunctionCallCount*2 + observations.BooleanOperatorChange*2
}
