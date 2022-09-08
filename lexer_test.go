package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTokens(l *Lexer) []Token {
	tokens := make([]Token, 0)
	for {
		_, tok, _ := l.Lex()
		if tok == EOF {
			break
		}

		tokens = append(tokens, tok)
	}

	return tokens
}

func TestBasicToken(t *testing.T) {
	input := "+-*/="
	expected := []Token{ADD, SUB, MUL, DIV, ASSIGN}

	reader := strings.NewReader(input)
	l := NewLexer(reader)
	tokens := getTokens(l)

	assert.Len(t, expected, len(tokens))
	assert.Equal(t, expected, tokens)
}

func TestIntToken(t *testing.T) {
	cases := []struct {
		input    string
		expected []Token
	}{
		{"123+23", []Token{INT, ADD, INT}},
		{"123-", []Token{INT, SUB}},
		{"123 456", []Token{INT, INT}},
		{"123One", []Token{INT, IDENT}},
	}

	for _, c := range cases {
		reader := strings.NewReader(c.input)
		l := NewLexer(reader)
		tokens := getTokens(l)

		assert.Len(t, c.expected, len(tokens))
		assert.Equal(t, c.expected, tokens)
	}
}

func TestIdentToken(t *testing.T) {
	cases := []struct {
		input    string
		expected []Token
	}{
		{"testIdent", []Token{IDENT}},
		{"ill)egal", []Token{IDENT, ILLEGAL, IDENT}},
		{"ill.egal", []Token{IDENT, ILLEGAL, IDENT}},
		{"one two", []Token{IDENT, IDENT}},
	}

	for _, c := range cases {
		reader := strings.NewReader(c.input)
		l := NewLexer(reader)
		tokens := getTokens(l)

		assert.Len(t, c.expected, len(tokens))
		assert.Equal(t, c.expected, tokens)
	}
}
