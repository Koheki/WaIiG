package token

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

// データ型
type TokenType string

// リテラル
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // ILLEGAL：未知のトークン、文字
	EOF     = "EOF"     // EOF：ファイルの終端（End of File）

	// 識別子＋リテラル
	IDENT = "IDENT"
	INT   = "INT"

	// 演算子
	ASSIGN = "="
	PLUS   = "+"

	// デリミタ（区切り文字）
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// キーワード
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
