package token

import (
	"testing"
)


func TestTokenTypeLookup(t *testing.T){
	if LookupIdentifier("let") != LET {
		t.Fatalf("expected LET")
	}

	if LookupIdentifier("fn") != FUNCTION {
		t.Fatalf("expected FUNCTION")
	}

	if LookupIdentifier("x_y") != IDENT {
		t.Fatalf("expected IDENT")
	}
}
