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
)