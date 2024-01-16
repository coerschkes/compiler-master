package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ASSIGN    = "="
	PLUS      = "+"
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	GREATER   = ">"
	LESS      = "<"
	MINUS     = "-"
	SLASH     = "/"
	ASTERISK  = "*"
	BANG      = "!"
	EQ        = "=="
	NOT_EQ    = "!="
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	IDENT     = "IDENT"
	INT       = "INT"
	FUNCTION  = "FUNCTION"
	VAR       = "VAR"
	TRUE      = "TRUE"
	FALSE     = "FALSE"
	IF        = "IF"
	ELSE      = "ELSE"
	RETURN    = "RETURN"
	EXIT      = "EXIT"
)

var SimpleTokens = map[byte]TokenType{
	'=': ASSIGN,
	'+': PLUS,
	',': COMMA,
	';': SEMICOLON,
	'(': LPAREN,
	')': RPAREN,
	'{': LBRACE,
	'}': RBRACE,
	'>': GREATER,
	'<': LESS,
	'-': MINUS,
	'/': SLASH,
	'*': ASTERISK,
	'!': BANG,
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"var":    VAR,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"exit":   EXIT,
}

func NewToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
