package lexer

func (lexer *MySQLLexer) scanNChar(lval *yySymType) (int, byte) {

	// found N'string'
	lexer.yyGet() // Skip '

	// Skip any char except '
	var c byte
	for c = lexer.yyGet(); c != 0 && c != '\''; c = lexer.yyGet() {
	}

	if c != '\'' {
		return ABORT_SYM, c
	}

	lval.bytes = lexer.buf[lexer.tok_start:lexer.ptr]

	return NCHAR_STRING, c
}
