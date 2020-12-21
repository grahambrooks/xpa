package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/grahambrooks/xpath-parser/parser/xpath"
	"os"
)

type Input struct {
	Pattern string `json:"pattern"`
}

func main() {
	fmt.Println("Hello all")
	parser := xpath.Parser{}
	for _, arg := range os.Args[1:] {
		if _, err := os.Stat(arg); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Analysing %s\n", arg)
			parser.Parse(arg)
		} else {
			if f, err := os.Open(arg); err != nil {
				_, _ = fmt.Fprint(os.Stderr, "Failed opening input file %s", arg)
			} else {
				s := bufio.NewScanner(f)
				for s.Scan() {
					var v Input
					if err := json.Unmarshal(s.Bytes(), &v); err != nil {
						_, _ = fmt.Fprintf(os.Stderr, "Failed to decode line")
					}
					parser.Parse(v.Pattern)
				}
			}
		}
	}
}
