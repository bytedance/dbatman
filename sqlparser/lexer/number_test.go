package lexer

import (
	"testing"
)

func TestInt(t *testing.T) {
	//testMatchReturn(t, `123456`, NUM, false)
	//testMatchReturn(t, `2147483646`, NUM, false)  // NUM
	//testMatchReturn(t, `2147483647`, NUM, false)  // 2^31 - 1
	testMatchReturn(t, `+2147483648`, NUM, true) // -2^31
}

/*
func TestFloatNum(t *testing.T) {
	testMatchReturn(t, " 10e-10", FLOAT_NUM, false)
	testMatchReturn(t, " 10E+10", FLOAT_NUM, false)
	testMatchReturn(t, "   10E10", FLOAT_NUM, false)
}*/
