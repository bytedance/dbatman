package lexer

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

func (lexer *MySQLLexer) getQuotedText() ([]byte, error) {
	var c byte
	var found_escape bool

	sep := lexer.yyLookHead()
	for lexer.ptr < uint(len(lexer.buf)) {
		c = lexer.yyNext()
		if c == '\\' && !lexer.sqlMode.MODE_NO_BACKSLASH_ESCAPES {
			// backslash as a escape charactar
			found_escape = true
			if lexer.ptr == uint(len(lexer.buf)) {
				return nil, nil
			}
			lexer.yySkip() // skip
		} else if c == sep {
			if sep == lexer.yyNext() { // // Check if two separators in a row
				found_escape = true
				continue
			} else {
				lexer.yyBack() // Backward 1 char
			}

			/* Found end. Unescape and return string */
			if !found_escape {
				return lexer.buf[lexer.tok_start:lexer.ptr], nil
			}
			found_escape = true
			ret := make([]byte, 0, 32)
			if !found_escape && !lexer.sqlMode.MODE_NO_BACKSLASH_ESCAPES && lexer.yyPeek() == '\\' && lexer.yyPeek2() != EOF {
				switch lexer.yyNext() {
				case 'n':
					ret = append(ret, '\n')
				case 't':
					ret = append(ret, '\t')
				case 'r':
					ret = append(ret, '\r')
				case 'b':
					ret = append(ret, '\b')
				case '0':
					ret = append(ret, 0)
				case 'Z':
					ret = append(ret, '\032')
				case '_', '%':
					ret = append(ret, '\\')
				default:
					found_escape = true
					lexer.yyBack()
				}
			} else if !found_escape && lexer.yyPeek() == sep {
				found_escape = true
			} else {
				ret = append(ret, lexer.yyNext())
				found_escape = false
			}

			return ret, nil
		}
	}

	return nil, StringFormatError
}
