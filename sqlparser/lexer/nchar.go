package lexer

func (lexer *MySQLLexer) scanNChar(lval *yySymType) (int, byte) {

	// found N'string'
	lexer.yyNext() // Skip '

	// Skip any char except '
	var c byte
	for c = lexer.yyNext(); c != 0 && c != '\''; c = lexer.yyNext() {
	}

	if c != '\'' {
		return ABORT_SYM, c
	}

	lval.bytes = lexer.buf[lexer.tok_start:lexer.ptr]

	return NCHAR_STRING, c
}
