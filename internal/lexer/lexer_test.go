package lexer

import "testing"

func TestTokenize(t *testing.T) {
	//expect comments to be ignored
	//check that the lexed of "1/2, 3/1, 4/5" and "1/2,3/1this is a comment,4/5" are the same
	input1 := "1/2, 3/1, 4/5"
	input2 := "1/2,3/1this is a comment,4/5"

	lexer := New()

	tokens1, err1 := lexer.Tokenize(input1)
	if err1 != nil {
		t.Fatalf("unexpected error: %v", err1)
	}

	tokens2, err2 := lexer.Tokenize(input2)
	if err2 != nil {
		t.Fatalf("unexpected error: %v", err2)
	}

	if len(tokens1) != len(tokens2) {
		t.Fatalf("expected same number of tokens, got %d and %d", len(tokens1), len(tokens2))
	}

	for i := range tokens1 {
		if tokens1[i] != tokens2[i] {
			t.Errorf("expected token %d to be the same, got %v and %v", i, tokens1[i], tokens2[i])
		}
	}
}
