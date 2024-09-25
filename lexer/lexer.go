package lexer

import "freedom/token"
type Lexer struct {
	input string
	position int
	readPosition int 
	ch byte //(input[readPosition])
}

func New(input string)*Lexer{
l:=&Lexer{input: input}
// set the lexer to initial position
l.readChar()
return l
}

func isLetter(ch byte) bool{
	return 'a'<=ch&&ch<='z'||'A'<=ch&&ch<='Z';
}

func isDigit(ch byte) bool{
	return '0'<=ch&&ch<='9'
}
// why pointer reciever method,because they operate on the state of data of the lexer
// if not function pointer, then we have to write addition code,reassignment,copies of lexer and soon
// which is not feasible and would make things really hard, when our lexer would expands.
func (l *Lexer) readNumber() string{
	tempPosition:=l.position
	for isDigit(l.ch){
		l.readChar()
	}
	return l.input[tempPosition:l.position]
}
func (l *Lexer) readIdentifier() string{
	tempPosition:=l.position
	for isLetter(l.ch){
		l.readChar()
	}
	return l.input[tempPosition:l.position]
}

func (l *Lexer) readChar(){
	if l.readPosition>=len(l.input){
		// we are finished reading
		l.ch=0
	}else {
		l.ch=l.input[l.readPosition]
	}
	l.position=l.readPosition
	l.readPosition+=1;
}
func (l *Lexer) skipWhitespace() {
	// skip newlines,whitespaces,tabs etc.
for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
l.readChar()
}
}
func (l  *Lexer) NextToken()token.Token{
	// while reading tokens, we need to skip the whitespaces.
	// we won't be analysing whitespaces,whitespaces are not significant
	// but reading them is not neccassarily tough also
	var tok token.Token 
	l.skipWhitespace()
	switch l.ch {
		case '=':
		if l.PeekChar()=='='{
		ch:=l.ch
		l.readChar()
		tok=token.Token{Type: token.EQ,Literal: string(ch)+ string(l.ch)}
		}else{
			tok=newToken(token.ASSIGN,l.ch)
		}
		case '!':
		if l.PeekChar() == '=' {
		ch := l.ch
		l.readChar()
		tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
		tok = newToken(token.BANG, l.ch)
		}
		case ';':
		tok = newToken(token.SEMICOLON, l.ch)
		case '(':
		tok = newToken(token.LPAREN, l.ch)
		case ')':
		tok = newToken(token.RPAREN, l.ch)
		case ',':
		tok = newToken(token.COMMA, l.ch)
		case '+':
		tok = newToken(token.PLUS, l.ch)
		case '{':
		tok = newToken(token.LBRACE, l.ch)
		case '}':
		tok = newToken(token.RBRACE, l.ch)
		case'<':
		tok=newToken(token.LT,l.ch)
		case '>':
		tok=newToken(token.GT,l.ch);
		case '/':
		tok = newToken(token.SLASH, l.ch)
		case '*':
		tok = newToken(token.ASTERISK, l.ch)
		case '-':
		tok=newToken(token.MINUS,l.ch)
		case 0:
		tok.Literal = ""
		tok.Type = token.EOF
		

		default:
			// we need the literal or value and token
			// type
			if isLetter(l.ch){
				tok.Literal=l.readIdentifier();
				tok.Type=token.LookupIdent(tok.Literal)
				// early return because readPosition is already in place.
			return tok;
			}else if isDigit(l.ch){
				tok.Literal=l.readNumber();
				tok.Type=token.INT;
				return tok
			}else{
				tok=newToken(token.ILLEGAL,l.ch)
			}
			
	}
	l.readChar()
return tok
}

func (l *Lexer) PeekChar()byte{
	// readPosition is also used for peeking character
	if l.readPosition>=len(l.input){
		return 0
	}else {
		return l.input[l.readPosition]
	}
}
func newToken(tokenType token.TokenType,ch byte)token.Token{
	return token.Token{Type: tokenType,Literal: string(ch)}
}