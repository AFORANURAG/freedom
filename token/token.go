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
BANG= "!"
ASTERISK = "*"
SLASH= "/"
LT = "<"
GT = ">"
MINUS= "-"

EQ= "=="
NOT_EQ = "!="

// keywords are

FUNCTION = "FUNCTION"
LET= "LET"
TRUE= "TRUE"
FALSE= "FALSE"
IF= "IF"
ELSE= "ELSE"
RETURN= "RETURN"

)


// let and func are keywords
var keywords = map[string]TokenType{
"fn":FUNCTION,
"let":LET,
"true":TRUE,
"false":FALSE,
"if":IF,
"else":ELSE,
"return": RETURN,
}

func LookupIdent(ident string) TokenType {
if tok, ok := keywords[ident]; ok {
return tok
}
return IDENT
}