package lexer

func (lexer *Lexer) scanNChar(lval *yySymType) (int, byte) {

	// found N'string'
	lex.yyGet() // Skip '

	// Skip any char except '
	for c := lex.yyGet(); c && c != '\''; c = lex.yyGet() {
	}

	if c != '\'' {
		return ABORT_SYM, c
	}

	lval.bytes = lex.buf[lex.tok_start:lex.ptr]

	return NCHAR_STRING, c
}
