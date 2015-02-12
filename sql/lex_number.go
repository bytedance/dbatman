package sql

import (
	"fmt"
)

const (
	LONG_LEN              = 10
	LONGLONG_LEN          = 19
	SIGNED_LONGLONG_LEN   = 19
	UNSIGNED_LONGLONG_LEN = 20
)

var (
	LONG              []byte = []byte{'2', '1', '4', '7', '4', '8', '3', '6', '4', '7'}
	SIGNED_LONG       []byte = []byte{'-', '2', '1', '4', '7', '4', '8', '3', '6', '4', '8'}
	LONGLONG          []byte = []byte{'9', '2', '2', '3', '3', '7', '2', '0', '3', '6', '8', '5', '4', '7', '7', '5', '8', '0', '7'}
	SIGNED_LONGLONG   []byte = []byte{'-', '9', '2', '2', '3', '3', '7', '2', '0', '3', '6', '8', '5', '4', '7', '7', '5', '8', '0', '8'}
	UNSIGNED_LONGLONG []byte = []byte{'1', '8', '4', '4', '6', '7', '4', '4', '0', '7', '3', '7', '0', '9', '5', '5', '1', '6', '1', '5'}
)

func (lex *SQLLexer) scanInt(lval *MySQLSymType) int {
	length := lex.ptr - lex.tok_start
	lval.bytes = lex.buf[lex.tok_start:lex.ptr]

	if length < LONG_LEN {
		return NUM
	}

	neg := false
	start := lex.tok_start
	if lex.buf[start] == '+' {
		start += 1
		length -= 1
	} else if lex.buf[start] == '-' {
		start += 1
		length -= 1
		neg = true
	}

	// ignore any '0' character
	for start < lex.ptr && lex.buf[start] == '0' {
		start += 1
		length -= 1
	}

	if length < LONG_LEN {
		return NUM
	}

	var cmp []byte
	var smaller int
	var bigger int
	if neg {
		if length == LONG_LEN {
			cmp = SIGNED_LONG[1:len(SIGNED_LONG)]
			smaller = NUM
			bigger = LONG_NUM
		} else if length < SIGNED_LONGLONG_LEN {
			return LONG_NUM
		} else if length > SIGNED_LONGLONG_LEN {
			return DECIMAL_NUM
		} else {
			cmp = SIGNED_LONGLONG[1:len(SIGNED_LONGLONG)]
			smaller = LONG_NUM
			bigger = DECIMAL_NUM
		}
	} else {
		if length == LONG_LEN {
			cmp = LONG
			smaller = NUM
			bigger = LONG_NUM
		} else if length < LONGLONG_LEN {
			return LONG_NUM
		} else if length > LONGLONG_LEN {
			if length > UNSIGNED_LONGLONG_LEN {
				return DECIMAL_NUM
			}
			cmp = UNSIGNED_LONGLONG
			smaller = ULONGLONG_NUM
			bigger = DECIMAL_NUM
		} else {
			cmp = LONGLONG
			smaller = LONG_NUM
			bigger = ULONGLONG_NUM
		}
	}

	idx := 0
	for idx < len(cmp) && cmp[idx] == lex.buf[start] {
		DEBUG(fmt.Sprintf("cmp:[%c] buf[%c]\n", cmp[idx], lex.buf[start]))
		idx += 1
		start += 1
	}

	if idx == len(cmp) {
		return smaller
	}

	if lex.buf[start] <= cmp[idx] {
		return smaller
	}
	return bigger
}

func (lex *SQLLexer) scanFloat(lval *MySQLSymType, c *byte) (int, bool) {
	cs := lex.cs

	// try match (+|-)? digit+
	if lex.yyPeek() == '+' || lex.yyPeek() == '-' {
		lex.yySkip() // ignore this char
	}

	// at least we have 1 digit-char
	if cs.IsDigit(lex.yyPeek()) {
		for ; cs.IsDigit(lex.yyPeek()); lex.yySkip() {
		}

		lval.bytes = lex.buf[lex.tok_start:lex.ptr]
		return FLOAT_NUM, true
	}

	return 0, false
}
