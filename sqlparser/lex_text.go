package sql

import (
	"errors"
)

/**
 * For Anltr3 Defination:

SINGLE_QUOTED_TEXT
@init { int escape_count = 0; }:
    SINGLE_QUOTE
    (
        SINGLE_QUOTE SINGLE_QUOTE { escape_count++; }
        | {!SQL_MODE_ACTIVE(SQL_MODE_NO_BACKSLASH_ESCAPES)}? => ESCAPE_OPERATOR .  { escape_count++; }
        | {SQL_MODE_ACTIVE(SQL_MODE_NO_BACKSLASH_ESCAPES)}? => ~(SINGLE_QUOTE)
        | {!SQL_MODE_ACTIVE(SQL_MODE_NO_BACKSLASH_ESCAPES)}? => ~(SINGLE_QUOTE | ESCAPE_OPERATOR)
    )*
    SINGLE_QUOTE
    { EMIT(); LTOKEN->user1 = escape_count; }
;

DOUBLE_QUOTED_TEXT
@init { int escape_count = 0; }:
    DOUBLE_QUOTE
    (
        DOUBLE_QUOTE DOUBLE_QUOTE { escape_count++; }
        | {!SQL_MODE_ACTIVE(SQL_MODE_NO_BACKSLASH_ESCAPES)}? => ESCAPE_OPERATOR .  { escape_count++; }
        | {SQL_MODE_ACTIVE(SQL_MODE_NO_BACKSLASH_ESCAPES)}? => ~(DOUBLE_QUOTE)
        | {!SQL_MODE_ACTIVE(SQL_MODE_NO_BACKSLASH_ESCAPES)}? => ~(DOUBLE_QUOTE | ESCAPE_OPERATOR)
    )*
    DOUBLE_QUOTE
    { EMIT(); LTOKEN->user1 = escape_count; }
;
*/

var StringFormatError error = errors.New("text string format error")

func (lexer *SQLLexer) getQuotedText() ([]byte, error) {
	var dq bool
	var sep byte

	if sep = lexer.yyLookHead(); sep == '"' {
		dq = true
	}

	for lexer.ptr < uint(len(lexer.buf)) {
		c := lexer.yyNext()

		if c == '\\' && !lexer.sqlMode.MODE_NO_BACKSLASH_ESCAPES {
			if lexer.yyPeek() == EOF {
				return nil, StringFormatError
			}

			lexer.yySkip() // skip next char
		} else if matchQuote(c, dq) {
			if matchQuote(lexer.yyPeek(), dq) {
				// found a escape quote. Eg. '' ""
				lexer.yySkip() // skip for the second quote
				continue
			}
			// we have found the last quote
			return lexer.buf[lexer.tok_start:lexer.ptr], nil
		}
	}

	return nil, StringFormatError
}

func matchQuote(c byte, double_quote bool) bool {
	if double_quote {
		return c == '"'
	}

	return c == '\''
}
