package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT" // add , x,y
	INT   = "INT"

	ASSIGN = "="
	PLUS   = "+"

	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keyWords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookUpIndent(keyword string) TokenType {
	if tok, ok := keyWords[keyword]; ok {
		return tok
	}
	return IDENT
}
