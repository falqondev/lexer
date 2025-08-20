package lexer

import (
	"io"
	"strings"
)

type Lexer struct {
	s        *Scanner
	tokens   map[Token]string
	keywords map[string]Token
}

func NewLexer(r io.Reader) *Lexer {
	l := &Lexer{
		tokens:   make(map[Token]string),
		keywords: make(map[string]Token),
	}

	l.s = newLexerScanner(r, l)
	for k, v := range baseTokens {
		l.tokens[k] = v
	}

	l.keywords["and"] = AND
	l.keywords["or"] = OR
	l.keywords["true"] = TRUE
	l.keywords["false"] = FALSE

	return l
}

func (l *Lexer) SetTokenMap(keywordTokens map[Token]string) {
	for k, v := range keywordTokens {
		l.tokens[k] = v
		l.keywords[strings.ToLower(v)] = k
	}
}

func (l *Lexer) Lookup(ident string) Token {
	if tok, ok := l.keywords[strings.ToLower(ident)]; ok {
		return tok
	}
	return IDENT
}

func (l *Lexer) Scan() (tok Token, pos Pos, lit string) {
	return l.s.Scan()
}
func (l *Lexer) Peek() rune {
	return l.s.Peek()
}
func (l *Lexer) ScanRegex() (tok Token, pos Pos, lit string) {
	return l.s.ScanRegex()
}
func (l *Lexer) Scanner() *Scanner {
	return l.s
}
