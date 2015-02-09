package lexer

import (
	"github.com/wangjild/go-mysql-proxy/sqlparser/parser"
)

func (lexer *SQLLexer) scanNChar(lval *parser.MySQLSymType) (int, byte) {

	// found N'string'
	lexer.yyNext() // Skip '

	// Skip any char except '
	var c byte
	for c = lexer.yyNext(); c != 0 && c != '\''; c = lexer.yyNext() {
	}

	if c != '\'' {
		return parser.ABORT_SYM, c
	}

	lval.Bytes = lexer.buf[lexer.tok_start:lexer.ptr]

	return parser.NCHAR_STRING, c
}
