package ast

import (
	"compiler/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&VarStatement{
				Token: token.Token{Type: token.VAR, Literal: "var"},
				Name:  &Identifier{Token: token.Token{Type: token.IDENT, Literal: "x"}, Value: "x"},
				Value: &Identifier{Token: token.Token{Type: token.IDENT, Literal: "y"}, Value: "y"},
			},
		},
	}

	if program.String() != "var x = y;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
