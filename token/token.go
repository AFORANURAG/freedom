package token

type TokenType string

// Literal is the value it holds and type is the type of token
// For example var x=10;
// x here is the identifier, 10 is a literal of type integer and var is a keyWord
type Token struct {
	Literal string
	Type TokenType
}

// TokenTypes
const (
ILLEGAL = "ILLEGAL"
EOF= "EOF"
IDENT = "IDENT" 
INT= "INT"
ASSIGN= "="
PLUS= "+"
COMMA= ","
SEMICOLON = ";"
LPAREN = "("
RPAREN = ")"
LBRACE = "{"
RBRACE = "}"
FUNCTION = "FUNCTION"
LET= "LET"
)


