package lexer

import (
	"github.com/wangjild/go-mysql-proxy/sqlparser/token"
)

// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"bytes"
	"fmt"
	"strings"

	"micode.be.xiaomi.com/wangjing1/go-mysql-proxy/sqltypes"
)

const EOFCHAR = 0x100

// MySQLLexer is the struct used to generate SQL
// tokens for the parser.
type MySQLLexer struct {
	reader bufio.Reader

	buf []byte
	cur uint

	yylineno uint
	yytoklen uint

	tok_start uint
	tok_end   uint

	tok_start_pre uint
	tok_end_pre   uint

	next_state uint

	AllowComments bool
	ForceEOF      bool
	lastChar      uint16
	Position      int
	errorToken    []byte
	LastError     string
	posVarIndex   int
	ParseTree     Statement
}

// NewStringMySQLLexer creates a new MySQLLexer for the
// sql string.
func NewMySQLLexer(sql string) *MySQLLexer {
	return &MySQLLexer{
		reader: bufio.NewReader(strings.NewReader(sql)),
		buf:    []byte(sql),
	}
}

// Lex returns the next token form the MySQLLexer.
// This function is used by go yacc.
func (lex *MySQLLexer) Lex(lval *yySymType) int {

	cs := lval.charset
	sm := cs.StateMap
	im := cs.IdentMap

	token_start_lineno := lex.yylineno

	lex.tok_end_prev = lex.tok_end
	lex.tok_start_prev = lex.tok_start

	lex.tok_start = lex.ptr
	lex.tok_end = lex.ptr

	state := lex.next_state
	lval.next_state = MY_LEX_OPERATOR_OR_IDENT

	var c byte
	for {
		switch status {
		case MY_LEX_OPERATOR_OR_IDENT, MY_LEX_START:
			for c = yyGet(); sm[c] == MY_LEX_SKIP; c = yyGet() {
			}

			lex.tok_start = lex.ptr - 1
			state = sm[c]
			token_start_lineno = lex.yylineno

		case MY_LEX_ESCAPE:
			if lex.yyGet() == 'N' {
				// Allow \N as shortcut for NULL
				return NULL_SYM
			}
		case MY_LEXCHAR, MY_LEX_SKIP:
			if c == '-' && lex.yyPeek() == '-' && (isspace(cs, lex.yyPeek2()) ||
				iscntrl(cs, lex.yyPeek2())) {
				state = MY_LEX_COMMENT
			} else {

				lex.ptr = lex.tok_start
				c = lex.yyGet()
				if c != ')' {
					lex.next_state = MY_LEX_START
				}

				if c == ',' {
					lex.tok_start = lex.ptr
				} else if c == '?' && lex.stmt_prepare_mode && !ident_map[yyPeek()] {
					return token.PARAM_MARKER
				}

				return int(c)
			}

		case MY_LEX_IDENT_OR_NCHAR:
			if lex.yyPeek() != '\'' {
				state = MY_LEX_IDENT
				break
			}

            lex.yyGet() // Skip '

            // Skip any char except '
            for c = lex.yyGet(); c != 0 && c != '\'' {}
            
            if c != '\'' {
                return ABORT_SYM
            }
            
            lval.bytes = lex.buf[lex.tok_start : lex.ptr] 

            lex.yytoklen -= 3
            return NCHAR_STRING
        
        case MY_LEX_IDENT_OR_HEX:
            if lex.yyPeek() == '\'' {
                state = MY_LEX_HEX_NUMBER
            }

        case MY_LEX_IDENT_OR_BIN:
            if lex.yyPeek() == '\'' {
                state = MY_LEX_BIN_NUMBER;
            }

        case MY_LEX_IDENT:
           var start byte

            for result_state = c; ident_map[c = lex.yyGet()]; result_state |= c)
            {}

            if result_state & 0x80 {
                result_state = IDENT_QUOTED
            } else {
                result_state = IDENT
            }




		}





	}
}

// return current char
func (lex *MySQLLexer) yyGet() (b byte) {

	if lex.buf[lex.ptr] == '\n' || (lex.buf[lex.ptr] == '\r' && lex.buf[lex.ptr+1] != '\n') {
		lex.yylineno += 1
	}
	b = lex.buf[lex.ptr]
	lex.ptr += 1

	return
}

func (lex *MySQLLexer) yyPeek() (b byte) {
	b = lex.buf[lex.ptr]
}

func (lex *MySQLLexer) yyPeek2() (b byte) {
	b = lex.buf[lex.ptr+1]
}

func (lex *MySQLLexer) yySkip() (b byte) {
	b := lex.buf[lex.ptr]
	n := lex.buf[lex.ptr+1]
	if b == '\n' || (b == '\r' && n != '\n') {
		lex.yylineno += 1
	}

	lex.ptr += 1
}

// Error is called by go yacc if there's a parsing error.
func (tkn *MySQLLexer) Error(err string) {
	buf := bytes.NewBuffer(make([]byte, 0, 32))
	if tkn.errorToken != nil {
		fmt.Fprintf(buf, "%s at position %v near %s", err, tkn.Position, tkn.errorToken)
	} else {
		fmt.Fprintf(buf, "%s at position %v", err, tkn.Position)
	}
	tkn.LastError = buf.String()
}

// Scan scans the tokenizer for the next token and returns
// the token type and an optional value.
func (tkn *MySQLLexer) Scan() (int, []byte) {
	if tkn.ForceEOF {
		return 0, nil
	}

	if tkn.lastChar == 0 {
		tkn.next()
	}
	tkn.skipBlank()
	switch ch := tkn.lastChar; {
	case isLetter(ch):
		return tkn.scanIdentifier()
	case isDigit(ch):
		return tkn.scanNumber(false)
	case ch == ':':
		return tkn.scanBindVar()
	default:
		tkn.next()
		switch ch {
		case EOFCHAR:
			return 0, nil
		case '=', ',', ';', '(', ')', '+', '*', '%', '&', '|', '^', '~':
			return int(ch), nil
		case '?':
			tkn.posVarIndex++
			buf := new(bytes.Buffer)
			fmt.Fprintf(buf, ":v%d", tkn.posVarIndex)
			return VALUE_ARG, buf.Bytes()
		case '.':
			if isDigit(tkn.lastChar) {
				return tkn.scanNumber(true)
			} else {
				return int(ch), nil
			}
		case '/':
			switch tkn.lastChar {
			case '/':
				tkn.next()
				return tkn.scanCommentType1("//")
			case '*':
				tkn.next()
				return tkn.scanCommentType2()
			default:
				return int(ch), nil
			}
		case '-':
			if tkn.lastChar == '-' {
				tkn.next()
				return tkn.scanCommentType1("--")
			} else {
				return int(ch), nil
			}
		case '<':
			switch tkn.lastChar {
			case '>':
				tkn.next()
				return NE, nil
			case '=':
				tkn.next()
				switch tkn.lastChar {
				case '>':
					tkn.next()
					return NULL_SAFE_EQUAL, nil
				default:
					return LE, nil
				}
			default:
				return int(ch), nil
			}
		case '>':
			if tkn.lastChar == '=' {
				tkn.next()
				return GE, nil
			} else {
				return int(ch), nil
			}
		case '!':
			if tkn.lastChar == '=' {
				tkn.next()
				return NE, nil
			} else {
				return LEX_ERROR, []byte("!")
			}
		case '\'', '"':
			return tkn.scanString(ch, STRING)
		case '`':
			return tkn.scanLiteralIdentifier()
		default:
			return LEX_ERROR, []byte{byte(ch)}
		}
	}
}

func (tkn *MySQLLexer) skipBlank() {
	ch := tkn.lastChar
	for ch == ' ' || ch == '\n' || ch == '\r' || ch == '\t' {
		tkn.next()
		ch = tkn.lastChar
	}
}

func (tkn *MySQLLexer) scanIdentifier() (int, []byte) {
	buffer := bytes.NewBuffer(make([]byte, 0, 8))
	buffer.WriteByte(byte(tkn.lastChar))
	for tkn.next(); isLetter(tkn.lastChar) || isDigit(tkn.lastChar); tkn.next() {
		buffer.WriteByte(byte(tkn.lastChar))
	}
	lowered := bytes.ToLower(buffer.Bytes())
	if keywordId, found := keywords[string(lowered)]; found {
		return keywordId, lowered
	}
	return ID, buffer.Bytes()
}

func (tkn *MySQLLexer) scanLiteralIdentifier() (int, []byte) {
	buffer := bytes.NewBuffer(make([]byte, 0, 8))
	buffer.WriteByte(byte(tkn.lastChar))
	if !isLetter(tkn.lastChar) {
		return LEX_ERROR, buffer.Bytes()
	}
	for tkn.next(); isLetter(tkn.lastChar) || isDigit(tkn.lastChar); tkn.next() {
		buffer.WriteByte(byte(tkn.lastChar))
	}
	if tkn.lastChar != '`' {
		return LEX_ERROR, buffer.Bytes()
	}
	tkn.next()
	return ID, buffer.Bytes()
}

func (tkn *MySQLLexer) scanBindVar() (int, []byte) {
	buffer := bytes.NewBuffer(make([]byte, 0, 8))
	buffer.WriteByte(byte(tkn.lastChar))
	token := VALUE_ARG
	tkn.next()
	if tkn.lastChar == ':' {
		token = LIST_ARG
		buffer.WriteByte(byte(tkn.lastChar))
		tkn.next()
	}
	if !isLetter(tkn.lastChar) {
		return LEX_ERROR, buffer.Bytes()
	}
	for isLetter(tkn.lastChar) || isDigit(tkn.lastChar) || tkn.lastChar == '.' {
		buffer.WriteByte(byte(tkn.lastChar))
		tkn.next()
	}
	return token, buffer.Bytes()
}

func (tkn *MySQLLexer) scanMantissa(base int, buffer *bytes.Buffer) {
	for digitVal(tkn.lastChar) < base {
		tkn.ConsumeNext(buffer)
	}
}

func (tkn *MySQLLexer) scanNumber(seenDecimalPoint bool) (int, []byte) {
	buffer := bytes.NewBuffer(make([]byte, 0, 8))
	if seenDecimalPoint {
		buffer.WriteByte('.')
		tkn.scanMantissa(10, buffer)
		goto exponent
	}

	if tkn.lastChar == '0' {
		// int or float
		tkn.ConsumeNext(buffer)
		if tkn.lastChar == 'x' || tkn.lastChar == 'X' {
			// hexadecimal int
			tkn.ConsumeNext(buffer)
			tkn.scanMantissa(16, buffer)
		} else {
			// octal int or float
			seenDecimalDigit := false
			tkn.scanMantissa(8, buffer)
			if tkn.lastChar == '8' || tkn.lastChar == '9' {
				// illegal octal int or float
				seenDecimalDigit = true
				tkn.scanMantissa(10, buffer)
			}
			if tkn.lastChar == '.' || tkn.lastChar == 'e' || tkn.lastChar == 'E' {
				goto fraction
			}
			// octal int
			if seenDecimalDigit {
				return LEX_ERROR, buffer.Bytes()
			}
		}
		goto exit
	}

	// decimal int or float
	tkn.scanMantissa(10, buffer)

fraction:
	if tkn.lastChar == '.' {
		tkn.ConsumeNext(buffer)
		tkn.scanMantissa(10, buffer)
	}

exponent:
	if tkn.lastChar == 'e' || tkn.lastChar == 'E' {
		tkn.ConsumeNext(buffer)
		if tkn.lastChar == '+' || tkn.lastChar == '-' {
			tkn.ConsumeNext(buffer)
		}
		tkn.scanMantissa(10, buffer)
	}

exit:
	return NUMBER, buffer.Bytes()
}

func (tkn *MySQLLexer) scanString(delim uint16, typ int) (int, []byte) {
	buffer := bytes.NewBuffer(make([]byte, 0, 8))
	for {
		ch := tkn.lastChar
		tkn.next()
		if ch == delim {
			if tkn.lastChar == delim {
				tkn.next()
			} else {
				break
			}
		} else if ch == '\\' {
			if tkn.lastChar == EOFCHAR {
				return LEX_ERROR, buffer.Bytes()
			}
			if decodedChar := sqltypes.SqlDecodeMap[byte(tkn.lastChar)]; decodedChar == sqltypes.DONTESCAPE {
				ch = tkn.lastChar
			} else {
				ch = uint16(decodedChar)
			}
			tkn.next()
		}
		if ch == EOFCHAR {
			return LEX_ERROR, buffer.Bytes()
		}
		buffer.WriteByte(byte(ch))
	}
	return typ, buffer.Bytes()
}

func (tkn *MySQLLexer) scanCommentType1(prefix string) (int, []byte) {
	buffer := bytes.NewBuffer(make([]byte, 0, 8))
	buffer.WriteString(prefix)
	for tkn.lastChar != EOFCHAR {
		if tkn.lastChar == '\n' {
			tkn.ConsumeNext(buffer)
			break
		}
		tkn.ConsumeNext(buffer)
	}
	return COMMENT, buffer.Bytes()
}

func (tkn *MySQLLexer) scanCommentType2() (int, []byte) {
	buffer := bytes.NewBuffer(make([]byte, 0, 8))
	buffer.WriteString("/*")
	for {
		if tkn.lastChar == '*' {
			tkn.ConsumeNext(buffer)
			if tkn.lastChar == '/' {
				tkn.ConsumeNext(buffer)
				break
			}
			continue
		}
		if tkn.lastChar == EOFCHAR {
			return LEX_ERROR, buffer.Bytes()
		}
		tkn.ConsumeNext(buffer)
	}
	return COMMENT, buffer.Bytes()
}

func (tkn *MySQLLexer) ConsumeNext(buffer *bytes.Buffer) {
	if tkn.lastChar == EOFCHAR {
		// This should never happen.
		panic("unexpected EOF")
	}
	buffer.WriteByte(byte(tkn.lastChar))
	tkn.next()
}

func (tkn *MySQLLexer) next() {
	if ch, err := tkn.InStream.ReadByte(); err != nil {
		// Only EOF is possible.
		tkn.lastChar = EOFCHAR
	} else {
		tkn.lastChar = uint16(ch)
	}
	tkn.Position++
}

func isLetter(ch uint16) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '@'
}

func digitVal(ch uint16) int {
	switch {
	case '0' <= ch && ch <= '9':
		return int(ch) - '0'
	case 'a' <= ch && ch <= 'f':
		return int(ch) - 'a' + 10
	case 'A' <= ch && ch <= 'F':
		return int(ch) - 'A' + 10
	}
	return 16 // larger than any legal digit val
}

func isDigit(ch uint16) bool {
	return '0' <= ch && ch <= '9'
}

const (
	AST_AUTO_INCR = "auto_increment"
	AST_BIGINT    = "bigint"
	AST_CHAR      = "char"
	AST_CHARACTER = "character"
	AST_DATE      = "date"
	AST_ERRORS    = "errors"
	AST_EXPLAIN   = "explain"
	AST_INDEX     = "index"
	AST_INDEXES   = "indexes"
	AST_INT       = "int"
	AST_KEYS      = "keys"
	AST_FROM      = "from"
	AST_SMALLINT  = "smallint"
	AST_STATUS    = "status"
	AST_TEMPORARY = "temporary"
	AST_TIME      = "time"
	AST_TIMESTAMP = "timestamp"
	AST_TINYINT   = "tinyint"
	AST_VCHAR     = "varchar"
	AST_WARNINGS  = "warnings"
	AST_PRIMARY   = "primary"
)
