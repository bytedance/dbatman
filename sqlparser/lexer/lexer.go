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
	ptr uint

	yylineno uint
	yytoklen uint

	tok_start uint
	tok_end   uint

	tok_start_pre uint
	tok_end_pre   uint

	state      uint // current state
	next_state uint // next should be state

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
		status: MY_LEX_START,
	}
}

// Lex returns the next token form the MySQLLexer.
// This function is used by go yacc.
func (lex *MySQLLexer) Lex(lval *yySymType) int {

	var result_state int
	cs := lex.charset
	state_map := cs.StateMap
	ident_map := cs.IdentMap

	token_start_lineno := lex.yylineno

	lex.tok_start_prev = lex.tok_start
	lex.tok_end_prev = lex.tok_end

	lex.tok_start = lex.ptr
	lex.tok_end = lex.ptr

	lex.state = lex.next_state
	lex.next_state = MY_LEX_OPERATOR_OR_IDENT

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
			break TAG_CHAR_OR_SKIP
		case MY_LEX_CHAR, MY_LEX_SKIP:
		TAG_CHAR_OR_SKIP:
			if c == '-' && lex.yyPeek() == '-' && (cs.isspace(lex.yyPeek2()) ||
				cs.iscntrl(lex.yyPeek2())) {
				state = MY_LEX_COMMENT
				break
			}

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

		case MY_LEX_IDENT_OR_NCHAR:
			if lex.yyPeek() != '\'' {
				state = MY_LEX_IDENT
				break
			}

			var ret int
			ret, c = lex.scanNChar(lval)

			return ret
		case MY_LEX_IDENT_OR_HEX:
			if lex.yyPeek() == '\'' {
				state = MY_LEX_HEX_NUMBER
			}

		case MY_LEX_IDENT_OR_BIN:
			if lex.yyPeek() == '\'' {
				state = MY_LEX_BIN_NUMBER
			}

		case MY_LEX_IDENT:
			var start byte

			c = lex.yyGet()
			for result_state = c; ident_map[c]; result_state |= c {
				c = lex.yyGet()
			}

			if result_state & 0x80 {
				result_state = IDENT_QUOTED
			} else {
				result_state = IDENT
			}

			start := lex.ptr
			idc := lex.buf[lex.tok_start:lex.ptr]
			if lex.ignore_space {
				for ; state_map[c] == MY_LEX_SKIP; c = lex.yyGet() {
				}
			}

			if start == lex.ptr && c == '.' && ident_map[lex.yyPeek()] {
				lex.next_state = MY_LEX_IDENT_SEP
			} else {
				yyUnget()
				if tokval := token.FindKeyword(idc, c == '('); tokval {
                    lex.next_state = MY_LEX_START
                    return tokval
				}
                yySkip()
			}

            // match _charsername
            if idc[0] == '_' {
                if _, ok : = charsets[string(idc)]; ok {
                    return UNDERSCORE_CHARSET
                }
            }

            return result_state

        case MY_LEX_IDENT_SEP: // Found ident before
                               // And Now '.'
            c = lex.yyGet()
            lex.next_state= MY_LEX_IDENT_START
            if !ident_map[lex.yyPeek()] {
                lex.next_state = MY_LEX_START;
            }

            return int(c)

        case MY_LEX_NUMER_IDENT:
            for c = lex.yyGet(); cs.isdigit(c); c = lex.yyGet() {
            }

            if !ident_map[c] {
                state = MY_LEX_INT_OR_REAL
                break
            }

            if (c == 'e' || c == 'E') {
                
            }

		}

	}
}

// return current char
func (lex *MySQLLexer) yyGet() (b byte) {

	b = lex.yyPeek()
	n := lex.yyPeek2()
	if b == '\n' || (b == '\r' && n != '\n') {
		lex.yylineno += 1
	}

	lex.ptr += 1
	return
}

func (lex *MySQLLexer) yyUnget() {
	lex.ptr -= 1
	if lex.Peek() == '\n' || (lex.Peek() == '\r' && lex.Peek2() != '\n') {
		lex.yylineno -= 1
	}
}

func (lex *MySQLLexer) yyPeek() (b byte) {
	if lex.ptr < len(lex.buf) {
		b = lex.buf[lex.ptr]
	} else {
		b.token.EOF
	}
}

func (lex *MySQLLexer) yyPeek2() (b byte) {
	if lex.ptr+1 < len(lex.buf) {
		b = lex.buf[lex.ptr+1]
	} else {
		b = token.EOF
	}
}

func (lex *MySQLLexer) yySkip() (b byte) {
	b = lex.Peek()
	n := lex.Peek2()
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
