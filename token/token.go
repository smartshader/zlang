package token

// TODO: Use int or byte for speed
type TokenType string

// TODO: Add filename and line numbers for better error messages
type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"

	// Identifiers + literals
	IDENT     = "IDENT"
	INT       = "INT"

	// Operators
	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	ASTERISK  = "*"
	SLASH     = "/"
	BANG      = "!"
	LT        = "<"
	GT        = ">"

	EQ        = "=="
	NOT_EQ    = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	FUNCTION  = "FUNCTION"
	LET       = "LET"
	TRUE      = "TRUE"
	FALSE     = "FALSE"
	IF        = "IF"
	ELSE      = "ELSE"
	RETURN    = "RETURN"
)

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}