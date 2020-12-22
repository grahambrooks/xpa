package main

import (
	"github.com/grahambrooks/xpa/pkg/xpa"
	"testing"
)

func TestTheParser(t *testing.T) {
	tests := []struct {
		xpath    string
		parsable bool
	}{
		{"/bookstore/book/price[text()]", true},
		{"//*[@id=\"id\"]", true},
		{"/bookstore/book[price>35]/price", true},
	}
	for _, test := range tests {
		t.Run(test.xpath, func(t *testing.T) {
			analyser := xpa.Analyser{}
			_, err := analyser.Analyse(test.xpath)
			if err !=nil {
				t.Fail()
			}
		})
	}
}
