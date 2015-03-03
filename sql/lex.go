package sql

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/wangjild/go-mysql-proxy/sql/charset"
	. "github.com/wangjild/go-mysql-proxy/sql/state"
	"strconv"
	"strings"
)

const EOFCHAR = 0x100
const NAMES_SEP_CHAR byte = '\377' /* Char to sep. names */

const MYSQL_VERSION_ID = 50109

// SQLLexer is the struct used to generate SQL
// tokens for the
type SQLLexer struct {
	reader *bufio.Reader

	buf []byte
	ptr uint

	yytoklen uint

	tok_start uint
	tok_end   uint

	tok_start_prev uint
	tok_end_prev   uint

	next_state uint // next should be state

	in_comment uint

	cs *charset.CharsetInfo

	stmt_prepare_mode bool
	ignore_space      bool
	sqlMode           SQLMode

	errorToken *string

	ParseTree IStatement
	LastError string
}

type SQLMode struct {
	MODE_ANSI_QUOTES          bool
	MODE_HIGH_NOT_PRECEDENCE  bool
	MODE_PIPES_AS_CONCAT      bool
	MODE_NO_BACKSLASH_ESCAPES bool
	MODE_IGNORE_SPACE         bool
}

// NewStringSQLLexer creates a new SQLLexer for the
// sql string.
func NewSQLLexer(sql string) *SQLLexer {
	return &SQLLexer{
		reader:     bufio.NewReader(strings.NewReader(sql)),
		buf:        []byte(sql),
		next_state: MY_LEX_START,
		cs:         charset.CSUtf8GeneralCli,
	}
}

// Lex returns the next token form the SQLLexer.
// This function is used by go yacc.
func (lex *SQLLexer) Lex(lval *MySQLSymType) (retstate int) {

	var result_state int
	var length uint
	var state uint
	var c byte

	cs := lex.cs
	state_map := cs.StateMap
	ident_map := cs.IdentMap

	lex.tok_start_prev = lex.tok_start
	lex.tok_end_prev = lex.tok_end

	lex.tok_start = lex.ptr
	lex.tok_end = lex.ptr

	state = lex.next_state
	lex.next_state = MY_LEX_OPERATOR_OR_IDENT

	// DEBUG("dbg buf:[" + string(lex.buf) + "]\ndbg enter:\n")
	for {
		DEBUG("\t" + GetLexStatus(state) + " current_buf[" + string(lex.buf[lex.ptr:]) + "]\n")
		switch state {
		case MY_LEX_OPERATOR_OR_IDENT, MY_LEX_START:
			for c = lex.yyNext(); state_map[c] == MY_LEX_SKIP; c = lex.yyNext() {
			}

			lex.tok_start = lex.ptr - 1
			state = state_map[c]
		case MY_LEX_ESCAPE:
			if lex.yyNext() == 'N' {
				// Allow \N as shortcut for NULL
				retstate = NULL_SYM
				goto TG_RET
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
			} else if c == '?' && ident_map[lex.yyPeek()] == 0 {
				retstate = PARAM_MARKER
				goto TG_RET
			}

			retstate = int(c)
			goto TG_RET

		case MY_LEX_IDENT_OR_NCHAR:
			if lex.yyPeek() != '\'' {
				state = MY_LEX_IDENT
				break
			}

			retstate, c = lex.scanNChar(lval)
			goto TG_RET
		case MY_LEX_IDENT_OR_HEX:
			if lex.yyPeek() == '\'' {
				state = MY_LEX_HEX_NUMBER
				break
			}
			DEBUG("MY_LEX_IDENT_OR_BIN")
			fallthrough
		case MY_LEX_IDENT_OR_BIN:
			if lex.yyPeek() == '\'' {
				state = MY_LEX_BIN_NUMBER
				break
			}
			DEBUG("MY_LEX_IDENT")
			fallthrough
		case MY_LEX_IDENT:

			retstate, lval.bytes = lex.getIdentifier()
			goto TG_RET

		case MY_LEX_IDENT_SEP: // Found ident before
			// And Now '.'
			c = lex.yyNext()
			lex.next_state = MY_LEX_IDENT_START
			if ident_map[lex.yyPeek()] == 0 {
				lex.next_state = MY_LEX_START
			}

			retstate = int(c)
			goto TG_RET

		case MY_LEX_NUMBER_IDENT: // number or ident which num-start
			for cs.IsDigit(lex.yyPeek()) {
				c = lex.yyNext()
			}
			c = lex.yyPeek()

			if ident_map[c] == 0 {
				// Can't be identifier
				state = MY_LEX_INT_OR_REAL
				break
			}

			if c == 'e' || c == 'E' {
				lex.yySkip() // skip for e/E
				var ok bool
				if retstate, ok = lex.scanFloat(lval, &c); ok {
					goto TG_RET
				}
			} else if c == 'x' && (lex.ptr-lex.tok_start) == 1 && lex.buf[lex.tok_start] == '0' {
				lex.yySkip() // skip for 'x'
				// 0xdddd number
				for c = lex.yyNext(); cs.IsXdigit(c); c = lex.yyNext() {
				}

				if lex.ptr-lex.tok_start >= 4 && ident_map[c] == 0 {
					retstate = HEX_NUM
					goto TG_RET
				}

				lex.yyBack()
			} else if c == 'b' && lex.ptr-lex.tok_start == 1 && lex.buf[lex.tok_start] == '0' {
				lex.yySkip() // skip for 'b'
				// binary number 0bxxxx
				for c = lex.yyNext(); cs.IsXdigit(c); c = lex.yyNext() {
				}

				if lex.ptr-lex.tok_start >= 4 && ident_map[c] == 0 {
					retstate = BIN_NUM
					goto TG_RET
				}

				lex.yyBack()
			}

			fallthrough
		case MY_LEX_IDENT_START:
			retstate, lval.bytes = lex.getPureIdentifier()
			goto TG_RET

		case MY_LEX_USER_VARIABLE_DELIMITER:
			quote_char := c

			for c = lex.yyNext(); c != EOF; c = lex.yyNext() {
				if c == NAMES_SEP_CHAR {
					break
				}

				if c == quote_char && lex.yyPeek() != quote_char {
					break
				}
			}

			if c == EOF {
				retstate = ABORT_SYM
				goto TG_RET
			}

			lex.next_state = MY_LEX_START
			lval.bytes = lex.buf[lex.tok_start:lex.ptr]
			retstate = IDENT_QUOTED
			goto TG_RET

		case MY_LEX_INT_OR_REAL:
			if c != '.' {
				retstate = lex.scanInt(lval)
				goto TG_RET
			}
			lex.yySkip()
			DEBUG("\tMY_LEX_REAL\n")
			fallthrough
		case MY_LEX_REAL:
			for cs.IsDigit(lex.yyPeek()) {
				c = lex.yyNext()
			}
			c = lex.yyPeek()

			if c == 'e' || c == 'E' {
				lex.yySkip() // skip for 'e'/'E'
				var ok bool
				if retstate, ok = lex.scanFloat(lval, &c); ok {
					goto TG_RET
				}

				state = MY_LEX_CHAR
				break
			}

			lval.bytes = lex.buf[lex.tok_start:lex.ptr]
			retstate = DECIMAL_NUM
			goto TG_RET
		case MY_LEX_HEX_NUMBER:
			lex.yySkip() // skip '
			for c = lex.yyNext(); cs.IsXdigit(c); c = lex.yyNext() {
			}

			length = lex.ptr - lex.tok_start

			if (length&1) == 0 || c != '\'' {
				retstate = ABORT_SYM
				goto TG_RET
			}

			lval.bytes = lex.buf[lex.tok_start:lex.ptr]
			retstate = HEX_NUM
			goto TG_RET

		case MY_LEX_BIN_NUMBER:
			lex.yyNext()
			for c = lex.yyNext(); c == '0' || c == '1'; c = lex.yyNext() {
			}

			length = lex.ptr - lex.tok_start
			if c != '\'' {
				retstate = ABORT_SYM
				goto TG_RET
			}

			lex.yyNext()

			retstate = BIN_NUM
			goto TG_RET
		case MY_LEX_CMP_OP:
			if state_map[lex.yyPeek()] == MY_LEX_CMP_OP || state_map[lex.yyPeek()] == MY_LEX_LONG_CMP_OP {
				lex.yySkip()
			}

			var ok bool
			if retstate, ok = findKeywords(lex.buf[lex.tok_start:lex.ptr], false); ok {
				lex.next_state = MY_LEX_START
				goto TG_RET
			}
			state = MY_LEX_CHAR

		case MY_LEX_LONG_CMP_OP:
			if state_map[lex.yyPeek()] == MY_LEX_CMP_OP || state_map[lex.yyPeek()] == MY_LEX_LONG_CMP_OP {
				lex.yySkip()
				if state_map[lex.yyPeek()] == MY_LEX_CMP_OP {
					lex.yySkip()
				}
			}

			var ok bool
			if retstate, ok = findKeywords(lex.buf[lex.tok_start:lex.ptr], false); ok {
				lex.next_state = MY_LEX_START
				goto TG_RET
			}
			state = MY_LEX_CHAR

		case MY_LEX_BOOL:
			if c != lex.yyPeek() {
				state = MY_LEX_CHAR
			} else {
				lex.yySkip()
				retstate, _ = findKeywords(lex.buf[lex.tok_start:lex.tok_start+2], false)
				lex.next_state = MY_LEX_START
				goto TG_RET
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
				retstate = ABORT_SYM
			} else {
				lval.bytes = b
				retstate = TEXT_STRING
			}
			goto TG_RET
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
				lex.yySkip() // skip '!'
				state = MY_LEX_START
				if cs.IsDigit(lex.yyPeek()) {
					start := lex.ptr
					lex.yySkip() // skip first digit
					for c = lex.yyPeek(); cs.IsDigit(c); c = lex.yyPeek() {
						lex.yySkip()
					}

					if i, err := strconv.Atoi(string(lex.buf[start:lex.ptr])); err != nil {
						lex.Error(err.Error())
						return ABORT_SYM
					} else {
						version = uint32(i)
					}
				}

				if version <= MYSQL_VERSION_ID {
					lex.in_comment = 1
					break
				}
			}

			// scan util match `*/`
			for c = lex.yyNext(); c != EOF && !(c == '*' && lex.yyPeek() == '/'); c = lex.yyNext() {
			}

			if c == '*' {
				lex.yySkip() // skip for '*/'
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
				retstate = SET_VAR
				goto TG_RET
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
				retstate = END_OF_INPUT
				goto TG_RET
			}

			state = MY_LEX_CHAR
		case MY_LEX_END:
			lex.next_state = MY_LEX_END
			retstate = 0
			goto TG_RET
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

			retstate = int('@')
			goto TG_RET

		case MY_LEX_HOSTNAME:
			for c = lex.yyNext(); cs.IsAlnum(c) || c == '.' || c == '_' || c == '$'; c = lex.yyNext() {
			}

			retstate = LEX_HOSTNAME
			goto TG_RET
		case MY_LEX_SYSTEM_VAR:
			lex.yySkip()
			lex.next_state = func() uint {
				if state_map[lex.yyPeek()] == MY_LEX_USER_VARIABLE_DELIMITER {
					return MY_LEX_OPERATOR_OR_IDENT
				} else {
					return MY_LEX_IDENT_OR_KEYWORD
				}
			}()

			retstate = int('@')
			goto TG_RET
		case MY_LEX_IDENT_OR_KEYWORD:
			result_state = 0
			c = lex.yyNext()
			for ident_map[c] != 0 {
				result_state |= int(c)
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
				retstate = ABORT_SYM
				goto TG_RET
			}

			val := lex.buf[lex.tok_start : lex.ptr-1]
			var ok bool
			if retstate, ok = findKeywords(val, false); ok {
				lex.yyBack()
				goto TG_RET
			}

			lval.bytes = val
			retstate = result_state
			goto TG_RET
		}
	}

	retstate = 0

TG_RET:

	DEBUG(fmt.Sprintf("dbg return [%s]\n", TokenName(retstate)))
	return
}

// return current char
func (lex *SQLLexer) yyNext() (b byte) {

	if lex.ptr < uint(len(lex.buf)) {
		b = lex.buf[lex.ptr]
		lex.ptr += 1
	} else {
		b = EOF
	}
	return
}

func (lex *SQLLexer) yyBack() {
	lex.ptr -= 1
}

func (lex *SQLLexer) yyPeek() (b byte) {
	if lex.ptr < uint(len(lex.buf)) {
		b = lex.buf[lex.ptr]
	} else {
		b = EOF
	}
	return
}

func (lex *SQLLexer) yyPeek2() (b byte) {
	if lex.ptr+1 < uint(len(lex.buf)) {
		b = lex.buf[lex.ptr+1]
	} else {
		b = EOF
	}

	return
}

func (lex *SQLLexer) yyLookHead() (b byte) {
	b = lex.buf[lex.ptr-1]
	return
}

func (lex *SQLLexer) yySkip() {
	lex.yyNext()
	return
}

// Error is called by go yacc if there's a parsing error.
func (lexer *SQLLexer) Error(err string) {
	buf := bytes.NewBuffer(make([]byte, 0, 32))
	fmt.Fprintf(buf, "%s at position %v near %s", err, lexer.ptr, string(lexer.buf[lexer.tok_start:lexer.ptr]))
	lexer.LastError = buf.String()
}
