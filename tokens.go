package main

func main() {

	type TokenType string

	type Token struct {
		Type    TokenType
		Literal string
	}

	const (
		ILLEGAL = "ILLEGAL"
		EOF     = "EOF"

		// Identifiers + literals
		IDENT  = "IDENT"  
		INT    = "INT"    
		STRING = "STRING" 

		// Operators
		ASSIGN   = "="
		PLUS     = "+"
		MINUS    = "-"
		BANG     = "!"
		ASTERISK = "*"
		SLASH    = "/"

		// Delimiters
		COMMA     = ","
		SEMICOLON = ";"
		LPAREN = "("
		RPAREN = ")"
		LBRACE = "{"
		RBRACE = "}"

		// Keywords
		FUNCTION = "FUNCTION"
		LET      = "LET"
		TRUE     = "TRUE"
		FALSE    = "FALSE"
		IF       = "IF"
		ELSE     = "ELSE"
		RETURN   = "RETURN"
		NULL     = "NULL"
	)

}
