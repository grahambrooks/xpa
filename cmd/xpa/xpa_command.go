package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/grahambrooks/xpa/pkg/xpa"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

type XpaCommand struct {
	cmd        *cobra.Command
	JsonOutput bool
}

func NewXpaCommand() *XpaCommand {
	var cmd = &cobra.Command{
		Use:   "xpa",
		Short: "XPath analyser",
		Long:  `A static analyser for XPath expressions`,
	}
	x := &XpaCommand{cmd: cmd}
	cmd.Flags().BoolVarP(&x.JsonOutput, "json-output", "j", false, "Output the analysis in json format")
	x.cmd.RunE = x.Run
	return x
}

func (c *XpaCommand) Formatter() *OutputFormatter {
	switch {
	case c.JsonOutput:
		encoder := json.NewEncoder(os.Stdout)
		return &OutputFormatter{
			ProgressFn: nil,
			ErrorFn:    nil,
			ObservationFn: func(observation *xpa.Analysis) {
				if err := encoder.Encode(observation); err != nil {
					_, _ = fmt.Fprintf(os.Stderr, "error encoding json")
				}
			},
		}
	case isTTY():
		red := color.New(color.FgRed)
		errorColour := red.Add(color.Bold)
		titleColour := color.New(color.FgGreen)
		progressColor := color.New(color.FgGreen)
		return &OutputFormatter{
			ProgressFn: func(format string, args ...interface{}) {
				_, _ = progressColor.Fprintf(os.Stderr, format, args...)
				fmt.Println()
			},
			ErrorFn: func(format string, args ...interface{}) {
				_, _ = errorColour.Fprintf(os.Stderr, format, args...)
				fmt.Println()
			},
			ObservationFn: func(observation *xpa.Analysis) {
				_, _ = fmt.Printf("complexity:         ")
				_, _ = titleColour.Printf("%d\n", observation.CognitiveComplexity())
				_, _ = fmt.Printf("function calls:     ")
				_, _ = titleColour.Printf("%d\n",  observation.FunctionCallCount)
				_, _ = fmt.Printf("boolean changes:    ")
				_, _ = titleColour.Printf("%d\n",  observation.BooleanOperatorChange)
				_, _ = fmt.Printf("comparison changes: ")
				_, _ = titleColour.Printf("%d\n", observation.ComparisonChange,)
				_, _ = fmt.Printf("tokens:             ")
				_, _ = titleColour.Printf("%d\n", observation.TerminalCount)
				_, _ = fmt.Printf("pattern:            ")
				_, _ = titleColour.Printf("%s\n", observation.Pattern)
			},
		}
	default:
		return &OutputFormatter{
			ProgressFn: func(format string, args ...interface{}) {
				_, _ = fmt.Fprintf(os.Stderr, format, args...)
			},
			ErrorFn: func(format string, args ...interface{}) {
				_, _ = fmt.Fprintf(os.Stderr, format, args...)
			},
			ObservationFn: func(analysis *xpa.Analysis) {
				_, _ = fmt.Printf("complexity: %d, function calls %d, boolean changes: %d, comparison changes %d, tokens %d pattern %s\n", analysis.CognitiveComplexity(), analysis.FunctionCallCount, analysis.BooleanOperatorChange, analysis.ComparisonChange, analysis.TerminalCount, analysis.Pattern)
			},
		}
	}
}

func (c *XpaCommand) Run(_ *cobra.Command, args []string) error {
	output := c.Formatter()
	output.Progress("XPath Analyser (xpa)\n")
	analyser := xpa.Analyser{}

	if len(args) == 0 {
		content, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			output.Error("Failed to read from stdin")
		} else {
			c.analyseString(analyser, string(content), output)
		}
	} else {
		for _, arg := range args {
			if _, err := os.Stat(arg); err == nil {
				c.analyseFile(analyser, arg, output)
			} else {
				c.analyseString(analyser, arg, output)
			}
		}
	}
	return nil
}

func (c *XpaCommand) analyseFile(analyser xpa.Analyser, arg string, output *OutputFormatter) {
	if f, err := os.Open(arg); err != nil {
		output.Error("could not read from file %s:%v", arg, err)
	} else {
		if content, err := ioutil.ReadAll(f); err != nil {
			output.Error("error reading from file %s:%v", arg, err)
		} else {
			if analysis, err := analyser.Analyse(string(content)); err != nil {
				output.Error("could not analyse %s: %v", string(content), err)
			} else {
				output.Observation(analysis)
			}
		}
	}
}

func (c *XpaCommand) analyseString(analyser xpa.Analyser, arg string, output *OutputFormatter) {
	if analysis, err := analyser.Analyse(arg); err != nil {
		output.Error("failed to analyse %s", arg)
	} else {
		output.Observation(analysis)
	}
}

func (c XpaCommand) Execute() error {
	return c.cmd.Execute()
}
