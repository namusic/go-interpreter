package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// id + literal
	IDENT = "IDENT" //add, foobar,x,y,・・・
	INT   = "INT"   //1234556

	// operator
	ASSIGN = "="
	PLUS   = "+"

	//delimiter
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//keyward
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
