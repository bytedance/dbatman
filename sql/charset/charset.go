package charset

import (
	"bytes"
	. "github.com/wangjild/go-mysql-proxy/sql/state"
)

type (
	CharsetInfo struct {
		Number        int
		PrimaryNumber int
		BinaryNumber  int

		CSName string
		Name   string

		CType []byte

		StateMap []uint
		IdentMap []uint
	}
)

func init() {
	ValidCharsets = make(map[string]*CharsetInfo)
	ValidCharsets["utf8_general_cli"] = CSUtf8GeneralCli

	for _, v := range ValidCharsets {
		initStateMaps(v)
	}
}

var ValidCharsets map[string]*CharsetInfo

func IsValidCharsets(cs []byte) bool {
	if _, ok := ValidCharsets[string(bytes.ToLower(cs))]; ok {
		return true
	}

	return false
}

func initStateMaps(cs *CharsetInfo) {

	var state_map [256]uint

	for i := 0; i < 256; i++ {
		if cs.IsAlpha(byte(i)) == true {
			state_map[i] = (MY_LEX_IDENT)
		} else if cs.IsDigit(byte(i)) {
			state_map[i] = MY_LEX_NUMBER_IDENT
		} else if cs.IsSpace(byte(i)) {
			state_map[i] = MY_LEX_SKIP
		} else {
			state_map[i] = MY_LEX_CHAR
		}
	}
	state_map[0] = MY_LEX_EOL
	state_map['_'] = MY_LEX_IDENT
	state_map['$'] = MY_LEX_IDENT
	state_map['\''] = MY_LEX_STRING
	state_map['.'] = MY_LEX_REAL_OR_POINT
	state_map['>'] = MY_LEX_CMP_OP
	state_map['='] = MY_LEX_CMP_OP
	state_map['!'] = MY_LEX_CMP_OP
	state_map['<'] = MY_LEX_LONG_CMP_OP
	state_map['&'] = MY_LEX_BOOL
	state_map['|'] = MY_LEX_BOOL
	state_map['#'] = MY_LEX_COMMENT
	state_map[';'] = MY_LEX_SEMICOLON
	state_map[':'] = MY_LEX_SET_VAR
	state_map['\\'] = MY_LEX_ESCAPE
	state_map['/'] = MY_LEX_LONG_COMMENT
	state_map['*'] = MY_LEX_END_LONG_COMMENT
	state_map['@'] = MY_LEX_USER_END
	state_map['`'] = MY_LEX_USER_VARIABLE_DELIMITER
	state_map['"'] = MY_LEX_STRING_OR_DELIMITER

	var ident_map [256]uint
	for i := 0; i < 256; i++ {
		ident_map[i] = func() uint {
			if state_map[i] == MY_LEX_IDENT || state_map[i] == MY_LEX_NUMBER_IDENT {
				return 1
			}
			return 0
		}()
	}

	state_map['x'] = MY_LEX_IDENT_OR_HEX
	state_map['X'] = MY_LEX_IDENT_OR_HEX
	state_map['b'] = MY_LEX_IDENT_OR_BIN
	state_map['B'] = MY_LEX_IDENT_OR_BIN
	state_map['n'] = (MY_LEX_IDENT_OR_NCHAR)
	state_map['N'] = (MY_LEX_IDENT_OR_NCHAR)

	cs.IdentMap = ident_map[:]
	cs.StateMap = state_map[:]
}

func (cs *CharsetInfo) IsAlpha(c byte) bool {
	if cs.CType[c+1]&(_MY_U|_MY_L) == 0 {
		return false
	}
	return true
}

func (cs *CharsetInfo) IsDigit(c byte) bool {
	if cs.CType[c+1]&_MY_NMR == 0 {
		return false
	}

	return true
}

func (cs *CharsetInfo) IsSpace(c byte) bool {
	if cs.CType[c+1]&_MY_SPC == 0 {
		return false
	}

	return true
}

func (cs *CharsetInfo) IsCntrl(c byte) bool {
	if cs.CType[c+1]&_MY_CTR == 0 {
		return false
	}

	return true
}

func (cs *CharsetInfo) IsXdigit(c byte) bool {
	if cs.CType[c+1]&_MY_X == 0 {
		return false
	}
	return true
}

func (cs *CharsetInfo) IsAlnum(c byte) bool {
	if cs.CType[c+1]&(_MY_U|_MY_L|_MY_NMR) == 0 {
		return false
	}

	return true
}

const (
	_MY_U   = 01
	_MY_L   = 02
	_MY_NMR = 04   /* Numeral (digit) */
	_MY_SPC = 010  /* Spacing character */
	_MY_PNT = 020  /* Punctuation */
	_MY_CTR = 040  /* Control character */
	_MY_B   = 0100 /* Blank */
	_MY_X   = 0200 /* heXadecimal digit */
)
