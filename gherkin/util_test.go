package gherkin

import (
	"strings"
	"testing"

	"github.com/l3pp4rd/go-behat/gherkin/lexer"
)

func (a *AST) assertMatchesTypes(expected []lexer.TokenType, t *testing.T) {
	key := -1
	for item := a.head; item != nil; item = item.next {
		key += 1
		if expected[key] != item.value.Type {
			t.Fatalf("expected ast token '%s', but got '%s' at position: %d", expected[key], item.value.Type, key)
		}
	}
	if len(expected)-1 != key {
		t.Fatalf("expected ast length %d, does not match actual: %d", len(expected), key+1)
	}
}

func indent(n int, s string) string {
	return strings.Repeat(" ", n) + s
}