package main

import (
	"github.com/grahambrooks/xpath-parser/parser/xpath"
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
			parser := xpath.Parser{}
			parser.Parse(test.xpath)
		})
	}
}
