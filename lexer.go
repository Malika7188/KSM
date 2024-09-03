package lexer

type Lexer struct {
	input        string //the source code lexer reading and a string containing the entire code to be broken into token
	position     int // keeps track of current position in the input and points to the character
	readposition int //current reading position in the input/is the index in the input where the lexer will read the next character.
	ch           byte // current character to be tokenized
}

//initializes the lexer with the input string and sets the initial state.
func NewLexer(input string) *Lexer{
	lex.:= &Lexer{input:input}
	lex.readchar() // It calls readChar to load the first character into ch.
	return lex
}

func (lex *Lexer) readchar() { //This function advances the position and readPosition pointers in the input string and updates ch to the current character.
	if lex.readposition >= len(lex.input) {
		lex.ch = 0 // sets ch to 0 if readposition go beyond len of input indicating EOL
	} else {
		lex.ch = i.input[i.readposition]
	}
	lex.position = lex.readposition
	readposition++
}

func (lex *Lexer) skipWhitespace() {
    for lex.ch == ' ' || lex.ch == '\t' || lex.ch == '\n' || lex.ch == '\r' {
        lex.readChar()
    }
}


func (lex *Lexer) peekChar() byte {
    if lex.readPosition >= len(lex.input) {
        return 0
    } else {
        return lex.input[lex.readPosition]
    }
}

func Identifier(ident string) TokenType {
    keywords := map[string]TokenType{
        "fn":     FUNCTION,
        "let":    LET,
        "true":   TRUE,
        "false":  FALSE,
        "if":     IF,
        "else":   ELSE,
        "return": RETURN,
        "null":   NULL,
    }

    if tokn, ok := keywords[ident]; ok {
        return tokn
    }
    return IDENT
}

func (lex *Lexer) readIdentifier() string {
    position := lex.position
    for isLetter(lex.ch) {
        lex.readChar()
    }
    return lex.input[position:lex.position]
}


func newToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}


func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (lex *Lexer) NewToke() Token {
	var tokn Token

	// lex.skipWhitespace()
	// if lex.ch =
}

	

	
	

