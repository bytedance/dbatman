package sql

import (
	"fmt"
	"github.com/bytedance/dbatman/sql/charset"
	. "github.com/bytedance/dbatman/sql/state"
)

func (lex *SQLLexer) getPureIdentifier() (int, []byte) {
	ident_map := lex.cs.IdentMap
	c := lex.yyPeek()
	rs := int(c)

	for ident_map[lex.yyPeek()] != 0 {
		rs |= int(c)
		c = lex.yyNext()
	}

	if rs&0x80 != 0 {
		rs = IDENT_QUOTED
	} else {
		rs = IDENT
	}

	if lex.yyPeek() == '.' && ident_map[int(lex.yyPeek2())] != 0 {
		lex.next_state = MY_LEX_IDENT_SEP
	}

	return rs, lex.buf[lex.tok_start:lex.ptr]
}

func (lex *SQLLexer) getIdentifier() (int, []byte) {

	ident_map := lex.cs.IdentMap

	c := lex.yyPeek()
	rs := int(c)

	for ident_map[lex.yyPeek()] != 0 {
		rs |= int(c)
		c = lex.yyNext()
	}

	if rs&0x80 != 0 {
		rs = IDENT_QUOTED
	} else {
		rs = IDENT
	}

	idc := lex.buf[lex.tok_start:lex.ptr]
	DEBUG(fmt.Sprintf("idc:[" + string(idc) + "]\n"))

	start := lex.ptr

	/*
		for ; lex.ignore_space && state_map[c] == MY_LEX_SKIP; c = lex.yyNext() {
		}*/

	c = lex.yyPeek()
	if start == lex.ptr && lex.yyPeek() == '.' && ident_map[int(lex.yyPeek())] != 0 {
		lex.next_state = MY_LEX_IDENT_SEP
	} else if ret, ok := findKeywords(idc, c == '('); ok {
		lex.next_state = MY_LEX_START
		return ret, idc
	}

	if idc[0] == '_' && charset.IsValidCharsets(idc[1:]) {
		return UNDERSCORE_CHARSET, idc
	}

	return rs, idc
}
