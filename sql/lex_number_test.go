package sql

import (
	"testing"
)

func TestInt(t *testing.T) {
	testMatchReturn(t, `123456`, NUM, false)
	testMatchReturn(t, `0000000000000000000000000123456`, NUM, false)
	testMatchReturn(t, `2147483646`, NUM, false)                           // NUM
	testMatchReturn(t, `2147483647`, NUM, false)                           // 2^31 - 1
	testMatchReturn(t, `2147483648`, LONG_NUM, false)                      // 2^31
	testMatchReturn(t, `0000000000000000000002147483648`, LONG_NUM, false) // 2^31
	testMatchReturn(t, `2147483648`, LONG_NUM, false)                      // 2^31
	testMatchReturn(t, `2147483648`, LONG_NUM, false)                      // 2^31
	testMatchReturn(t, `2147483648`, LONG_NUM, false)                      // 2^31

	testMatchReturn(t, `9223372036854775807`, LONG_NUM, false)
	testMatchReturn(t, `9223372036854775808`, ULONGLONG_NUM, false)
	testMatchReturn(t, `18446744073709551615`, ULONGLONG_NUM, false)
	testMatchReturn(t, `18446744073709551616`, DECIMAL_NUM, false)
}

func TestNum(t *testing.T) {
	testMatchReturn(t, `0x1234`, HEX_NUM, false)
	testMatchReturn(t, `0xa4234`, HEX_NUM, false)
	testMatchReturn(t, `0b0110`, BIN_NUM, false)
}

func TestFloatNum(t *testing.T) {
	testMatchReturn(t, " 10e-10", FLOAT_NUM, false)
	testMatchReturn(t, " 10E+10", FLOAT_NUM, false)
	testMatchReturn(t, "   10E10", FLOAT_NUM, false)
	testMatchReturn(t, "1.20E10", FLOAT_NUM, false)
	testMatchReturn(t, "1.20E-10", FLOAT_NUM, false)
}

func TestDecimalNum(t *testing.T) {
	testMatchReturn(t, `.21`, DECIMAL_NUM, false)
	testMatchReturn(t, `72.21`, DECIMAL_NUM, false)
}

func TestHex(t *testing.T) {
	testMatchReturn(t, `X'4D7953514C'`, HEX_NUM, false)

	testMatchReturn(t, `x'D34F2X`, ABORT_SYM, false)
	testMatchReturn(t, `x'`, ABORT_SYM, false)

}

func TestBin(t *testing.T) {
	testMatchReturn(t, `b'0101010111000'`, BIN_NUM, false)
	testMatchReturn(t, `b'0S01010111000'`, ABORT_SYM, false)
	testMatchReturn(t, `b'12312351123`, ABORT_SYM, false)
}

func TestMultiNum(t *testing.T) {
	str := `123     'string1' 18446744073709551616    1.20E-10 .312  x'4D7953514C' `
	lex, lval := getLexer(str)

	lexExpect(t, lex, lval, NUM)
	lvalExpect(t, lval, `123`)

	lexExpect(t, lex, lval, TEXT_STRING)
	lvalExpect(t, lval, `'string1'`)

	lexExpect(t, lex, lval, DECIMAL_NUM)
	lvalExpect(t, lval, `18446744073709551616`)

	lexExpect(t, lex, lval, FLOAT_NUM)
	lvalExpect(t, lval, `1.20E-10`)

	lexExpect(t, lex, lval, DECIMAL_NUM)
	lvalExpect(t, lval, `.312`)

	lexExpect(t, lex, lval, HEX_NUM)
	lvalExpect(t, lval, `x'4D7953514C'`)

	lexExpect(t, lex, lval, END_OF_INPUT)
}

func TestNumberInPlacehold(t *testing.T) {
	str := ` (5)`
	lex, lval := getLexer(str)
	lexExpect(t, lex, lval, '(')
	lexExpect(t, lex, lval, NUM)
	lexExpect(t, lex, lval, ')')
}
