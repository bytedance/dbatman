package lexer

import (
	. "github.com/wangjild/go-mysql-proxy/sqlparser/lexer/charset"
	. "github.com/wangjild/go-mysql-proxy/sqlparser/lexer/state"
)

// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

const EOFCHAR = 0x100
const NAMES_SEP_CHAR byte = '\377' /* Char to sep. names */

const MYSQL_VERSION_ID = 50109

// MySQLLexer is the struct used to generate SQL
// tokens for the parser.
type MySQLLexer struct {
	reader *bufio.Reader

	buf []byte
	ptr uint

	yylineno uint
	yytoklen uint

	tok_start uint
	tok_end   uint

	tok_start_prev uint
	tok_end_prev   uint

	next_state uint // next should be state

	in_comment uint

	charset *CharsetInfo

	stmt_prepare_mode bool
	ignore_space      bool
	sqlMode           SQLMode

	errorToken *string

	ParseTree string
	LastError string
}

type SQLMode struct {
	MODE_ANSI_QUOTES          bool
	MODE_HIGH_NOT_PRECEDENCE  bool
	MODE_PIPES_AS_CONCAT      bool
	MODE_NO_BACKSLASH_ESCAPES bool
	MODE_IGNORE_SPACE         bool
}

// NewStringMySQLLexer creates a new MySQLLexer for the
// sql string.
func NewMySQLLexer(sql string) *MySQLLexer {
	return &MySQLLexer{
		reader:     bufio.NewReader(strings.NewReader(sql)),
		buf:        []byte(sql),
		next_state: MY_LEX_START,
		charset:    CSUtf8GeneralCli,
	}
}

// Lex returns the next token form the MySQLLexer.
// This function is used by go yacc.
func (lex *MySQLLexer) Lex(lval *yySymType) int {

	var result_state int
	var length uint
	var state uint
	var c byte

	cs := lex.charset
	state_map := cs.StateMap
	ident_map := cs.IdentMap

	lex.tok_start_prev = lex.tok_start
	lex.tok_end_prev = lex.tok_end

	lex.tok_start = lex.ptr
	lex.tok_end = lex.ptr

	state = lex.next_state
	lex.next_state = MY_LEX_OPERATOR_OR_IDENT

	DEBUG("buf:[" + string(lex.buf) + "]")
	DEBUG("\ndbg enter:\n")
	defer DEBUG("dbg leave\n")
	for {
		DEBUG("\t" + GetLexStatus(state) + "\n")
		switch state {
		case MY_LEX_OPERATOR_OR_IDENT, MY_LEX_START:
			for c = lex.yyNext(); state_map[c] == MY_LEX_SKIP; c = lex.yyNext() {
			}

			lex.tok_start = lex.ptr - 1
			state = state_map[c]
		case MY_LEX_ESCAPE:
			if lex.yyNext() == 'N' {
				// Allow \N as shortcut for NULL
				return NULL_SYM
			}
			fallthrough
		case MY_LEX_CHAR, MY_LEX_SKIP:
			if c == '-' && lex.yyPeek() == '-' && (cs.IsSpace(lex.yyPeek2()) ||
				cs.IsCntrl(lex.yyPeek2())) {
				state = MY_LEX_COMMENT
				break
			}

			lex.ptr = lex.tok_start
			c = lex.yyNext()
			if c != ')' {
				lex.next_state = MY_LEX_START
			}

			if c == ',' {
				lex.tok_start = lex.ptr
			} else if c == '?' && lex.stmt_prepare_mode && ident_map[lex.yyPeek()] != 0 {
				return PARAM_MARKER
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
			var start uint

			c = lex.yyNext()
			for result_state = int(c); ident_map[int(c)] != 0; result_state |= int(c) {
				c = lex.yyNext()
			}

			if result_state&0x80 != 0 {
				result_state = IDENT_QUOTED
			} else {
				result_state = IDENT
			}

			start = lex.ptr
			idc := lex.buf[lex.tok_start:lex.ptr]
			if lex.ignore_space {
				for ; state_map[c] == MY_LEX_SKIP; c = lex.yyNext() {
				}
			}

			if start == lex.ptr && c == '.' && ident_map[int(lex.yyPeek())] != 0 {
				lex.next_state = MY_LEX_IDENT_SEP
			} else {
				lex.yyBack()
				if tokval, ok := findKeywords(idc, c == '('); ok {
					lex.next_state = MY_LEX_START
					return tokval
				}
				lex.yySkip()
			}

			// match _charsername
			if idc[0] == '_' {
				if _, ok := ValidCharsets[string(idc)]; ok {
					return UNDERSCORE_CHARSET
				}
			}

			return result_state

		case MY_LEX_IDENT_SEP: // Found ident before
			// And Now '.'
			c = lex.yyNext()
			lex.next_state = MY_LEX_IDENT_START
			if ident_map[lex.yyPeek()] == 0 {
				lex.next_state = MY_LEX_START
			}

			return int(c)

		case MY_LEX_NUMBER_IDENT: // number or ident which num-start
			for c = lex.yyNext(); cs.IsDigit(c); c = lex.yyNext() {
			}

			if ident_map[c] == 0 {
				// Can't be identifier
				state = MY_LEX_INT_OR_REAL
				break
			}

			if c == 'e' || c == 'E' {
				if cs.IsDigit(lex.yyPeek()) { // Allow 1E10
					if cs.IsDigit(lex.yyPeek()) { // Number must have digit after sign
						lex.yySkip()
						for tmpc := lex.yyNext(); cs.IsDigit(tmpc); tmpc = lex.yyNext() {
						} // until non-numberic char

						return FLOAT_NUM
					}
				} else if c = lex.yyNext(); c == '+' || c == '-' { // Allow 1E+10
					if cs.IsDigit(lex.yyPeek()) { // Number must have digit after sign
						lex.yySkip()
						for tmpc := lex.yyNext(); cs.IsDigit(tmpc); tmpc = lex.yyNext() {
						} // until non-numberic char

						return FLOAT_NUM
					}
				} else {
					lex.yyBack()
				}
			} else if c == 'x' && (lex.ptr-lex.tok_start) == 2 && lex.buf[lex.tok_start] == '0' {
				// 0xdddd number
				for c = lex.yyNext(); cs.IsXdigit(c); c = lex.yyNext() {
				}

				if lex.ptr-lex.tok_start >= 4 && ident_map[c] == 0 {
					return HEX_NUM
				}

				lex.yyBack()
			} else if c == 'b' && lex.ptr-lex.tok_start == 2 && lex.buf[lex.tok_start] == '0' {
				// binary number 0bxxxx
				for c = lex.yyNext(); cs.IsXdigit(c); c = lex.yyNext() {
				}

				if lex.ptr-lex.tok_start >= 4 && ident_map[c] == 0 {
					return BIN_NUM
				}

				lex.yyBack()
			}

			fallthrough
		case MY_LEX_IDENT_START:
			result_state = 0
			for c = lex.yyNext(); ident_map[int(c)] != 0; result_state |= int(c) {
			}

			result_state = result_state & 0x80
			if result_state != 0 {
				result_state = IDENT_QUOTED
			} else {
				result_state = IDENT
			}

			if c == '.' && ident_map[int(lex.yyPeek())] != 0 {
				lex.next_state = MY_LEX_IDENT_SEP
			}

			return result_state

		case MY_LEX_USER_VARIABLE_DELIMITER:
			double_quotes := 0
			quote_char := c

			lex.tok_start = lex.ptr
			for c = lex.yyNext(); c != 0; c = lex.yyNext() {
				if c == NAMES_SEP_CHAR {
					break
				}

				if c == quote_char {
					if lex.yyPeek() != quote_char {
						break
					}
					c = lex.yyNext()
					double_quotes += 1
				}
			}

			if double_quotes != 0 {

			} else {

			}

			if c == quote_char {
				lex.yySkip()
			}

			lex.next_state = MY_LEX_START
			return IDENT_QUOTED

		case MY_LEX_INT_OR_REAL:
			if c != '.' {
				// TODO
				// return getNumberType()
			}
			fallthrough
		case MY_LEX_REAL:
			for c = lex.yyNext(); cs.IsDigit(c); c = lex.yyNext() {
			}

			if c == 'e' || c == 'E' {
				c = lex.yyNext()
				if c == '-' || c == '+' {
					c = lex.yyNext() // skip sign
				}

				if !cs.IsDigit(c) {
					state = MY_LEX_CHAR
					break
				}

				for tmpc := lex.yyNext(); cs.IsDigit(tmpc); tmpc = lex.yyNext() {
				}

				return FLOAT_NUM
			}

			return DECIMAL_NUM
		case MY_LEX_HEX_NUMBER:
			lex.yyNext() // skip '
			for c = lex.yyNext(); cs.IsXdigit(c); c = lex.yyNext() {
			}

			length = lex.ptr - lex.tok_start

			if (length&1) == 0 || c != '\'' {
				return ABORT_SYM
			}

			lex.yyNext() //
			return HEX_NUM

		case MY_LEX_BIN_NUMBER:
			lex.yyNext()
			for c = lex.yyNext(); c == '0' || c == '1'; c = lex.yyNext() {
			}

			length = lex.ptr - lex.tok_start
			if c != '\'' {
				return ABORT_SYM
			}

			lex.yyNext()

			return BIN_NUM

		case MY_LEX_CMP_OP:
			if state_map[lex.yyPeek()] == MY_LEX_CMP_OP || state_map[lex.yyPeek()] == MY_LEX_LONG_CMP_OP {
				lex.yySkip()
			}

			if tokval, ok := findKeywords(lex.buf[lex.tok_start:lex.ptr], false); ok {
				lex.next_state = MY_LEX_START
				return tokval
			}
			state = MY_LEX_CHAR

		case MY_LEX_LONG_CMP_OP:
			if state_map[lex.yyPeek()] == MY_LEX_CMP_OP || state_map[lex.yyPeek()] == MY_LEX_LONG_CMP_OP {
				lex.yySkip()
				if state_map[lex.yyPeek()] == MY_LEX_CMP_OP {
					lex.yySkip()
				}
			}

			if tokval, ok := findKeywords(lex.buf[lex.tok_start:lex.ptr], false); ok {
				lex.next_state = MY_LEX_START
				return tokval
			}
			state = MY_LEX_CHAR

		case MY_LEX_BOOL:
			if c != lex.yyPeek() {
				state = MY_LEX_CHAR
			} else {
				lex.yySkip()
				tokval, _ := findKeywords(lex.buf[lex.tok_start:lex.tok_start+2], false)
				lex.next_state = MY_LEX_START
				return tokval
			}

		case MY_LEX_STRING_OR_DELIMITER:
			if lex.sqlMode.MODE_ANSI_QUOTES {
				state = MY_LEX_USER_VARIABLE_DELIMITER
				break
			}
			fallthrough
		case MY_LEX_STRING:
			b, err := lex.getQuotedText()
			if err != nil {
				lex.Error(err.Error())
				return ABORT_SYM
			} else {
				lval.bytes = b
				return TEXT_STRING
			}
		case MY_LEX_COMMENT:
			c = lex.yyNext()
			n := lex.yyPeek()
			for c != '\n' && !(c == '\r' && n != '\n') {
				c = lex.yyNext()
				n = lex.yyPeek()
			}

			lex.yyBack() // Safety against eof
			state = MY_LEX_START
		case MY_LEX_LONG_COMMENT:
			if lex.yyPeek() != '*' {
				state = MY_LEX_CHAR
				break
			}

			lex.yySkip() // skip '*'
			if lex.yyPeek() == '!' {
				var version uint32 = MYSQL_VERSION_ID
				lex.yySkip()
				state = MY_LEX_START
				if cs.IsDigit(lex.yyPeek()) {
					// TODO version = atoi
				}

				if version <= MYSQL_VERSION_ID {
					lex.in_comment = 1
					break
				}
			}

			for lex.ptr != uint(len(lex.buf)) {
				if c = lex.yyNext(); c != '*' || lex.yyPeek() != '/' {
					continue
				}
			} //

			if lex.ptr != uint(len(lex.buf)) {
				lex.yySkip()
			}

			state = MY_LEX_START

		case MY_LEX_END_LONG_COMMENT:
			if lex.in_comment != 0 && lex.yyPeek() == '/' {
				lex.yySkip()
				lex.in_comment = 0
				state = MY_LEX_START
			} else {
				state = MY_LEX_CHAR
			}
		case MY_LEX_SET_VAR:
			if lex.yyPeek() != '=' {
				state = MY_LEX_CHAR
			} else {
				lex.yySkip()
				return SET_VAR
			}

		case MY_LEX_SEMICOLON:
			if lex.yyPeek() != 0 {
				state = MY_LEX_CHAR
				break
			}
			fallthrough
		case MY_LEX_EOL:
			if lex.ptr >= uint(len(lex.buf)) {
				lex.next_state = MY_LEX_END
				return END_OF_INPUT
			}

			state = MY_LEX_CHAR
		case MY_LEX_END:
			lex.next_state = MY_LEX_END
			return 0
		case MY_LEX_REAL_OR_POINT:
			if cs.IsDigit(lex.yyPeek()) {
				state = MY_LEX_REAL
			} else {
				state = MY_LEX_IDENT_SEP
				lex.yyBack()
			}
		case MY_LEX_USER_END: // end '@' of user@hostname
			switch state_map[lex.yyPeek()] {
			case MY_LEX_STRING, MY_LEX_USER_VARIABLE_DELIMITER, MY_LEX_STRING_OR_DELIMITER:
			case MY_LEX_USER_END:
				lex.next_state = MY_LEX_SYSTEM_VAR
			default:
				lex.next_state = MY_LEX_HOSTNAME
			}

			return int('@')

		case MY_LEX_HOSTNAME:
			for c = lex.yyNext(); cs.IsAlnum(c) || c == '.' || c == '_' || c == '$'; c = lex.yyNext() {
			}

			return LEX_HOSTNAME
		case MY_LEX_SYSTEM_VAR:
			lex.yySkip()
			lex.next_state = func() uint {
				if state_map[lex.yyPeek()] == MY_LEX_USER_VARIABLE_DELIMITER {
					return MY_LEX_OPERATOR_OR_IDENT
				} else {
					return MY_LEX_IDENT_OR_KEYWORD
				}
			}()

			return int('@')
		case MY_LEX_IDENT_OR_KEYWORD:
			result_state = 0
			c = lex.yyNext()
			for ; ident_map[c] != 0; result_state |= int(c) {
				c = lex.yyNext()
			}

			if result_state&0x80 != 0 {
				result_state = IDENT_QUOTED
			} else {
				result_state = IDENT
			}

			if c == '.' {
				lex.next_state = MY_LEX_IDENT_SEP
			}

			length = lex.ptr - lex.tok_start - 1
			if length == 0 {
				return ABORT_SYM
			}

			if tokval, ok := findKeywords(lex.buf[lex.tok_start:lex.ptr-1], false); ok {
				lex.yyBack()
				return tokval
			}

			return result_state
		}
	}

	return 0
}

// return current char
func (lex *MySQLLexer) yyNext() (b byte) {

	b = lex.yyPeek()
	n := lex.yyPeek2()
	if b == '\n' || (b == '\r' && n != '\n') {
		lex.yylineno += 1
	}

	lex.ptr += 1
	return
}

func (lex *MySQLLexer) yyBack() {
	lex.ptr -= 1
	if lex.yyPeek() == '\n' || (lex.yyPeek() == '\r' && lex.yyPeek2() != '\n') {
		lex.yylineno -= 1
	}
}

func (lex *MySQLLexer) yyPeek() (b byte) {
	if lex.ptr < uint(len(lex.buf)) {
		b = lex.buf[lex.ptr]
	} else {
		b = EOF
	}
	return
}

func (lex *MySQLLexer) yyPeek2() (b byte) {
	if lex.ptr+1 < uint(len(lex.buf)) {
		b = lex.buf[lex.ptr+1]
	} else {
		b = EOF
	}

	return
}

func (lex *MySQLLexer) yyLookHead() (b byte) {
	b = lex.buf[lex.ptr-1]
	return
}

func (lex *MySQLLexer) yySkip() (b byte) {
	b = lex.yyPeek()
	n := lex.yyPeek2()
	if b == '\n' || (b == '\r' && n != '\n') {
		lex.yylineno += 1
	}

	lex.ptr += 1
	return
}

// Error is called by go yacc if there's a parsing error.
func (lexer *MySQLLexer) Error(err string) {
	buf := bytes.NewBuffer(make([]byte, 0, 32))
	if lexer.errorToken != nil {
		fmt.Fprintf(buf, "%s at position %v near %s", err, lexer.ptr, lexer.errorToken)
	} else {
		fmt.Fprintf(buf, "%s at position %v", err, lexer.ptr)
	}

	lexer.LastError = buf.String()
}
