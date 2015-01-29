package lexer

type (
	CharsetInfo struct {
		Number        int
		PrimaryNumber int
		BinaryNumber  int

		CSName string
		Name   string

		CType []byte

		StateMap []byte
		IdentMap []byte
	}
)

func initStateMaps(cs *CharsetInfo) {

	var state_map [256]byte

	for i := 0; i < 256; i++ {
		if isalpha(cs, byte(i)) == true {
			state_map[i] = byte(MY_LEX_IDENT)
		} else if isdigit(cs, byte(i)) {
			state_map[i] = byte(MY_LEX_NUMBER_IDENT)
		} else if isspace(cs, byte(i)) {
			state_map[i] = byte(MY_LEX_IDENT)
		} else {
			state_map[i] = byte(MY_LEX_CHAR)
		}
	}
	state_map[0] = byte(MY_LEX_EOL)
	state_map['_'] = byte(MY_LEX_IDENT)
	state_map['$'] = byte(MY_LEX_IDENT)
	state_map['\''] = byte(MY_LEX_STRING)
	state_map['.'] = byte(MY_LEX_REAL_OR_POINT)
	state_map['>'] = byte(MY_LEX_CMP_OP)
	state_map['='] = byte(MY_LEX_CMP_OP)
	state_map['!'] = byte(MY_LEX_CMP_OP)
	state_map['<'] = byte(MY_LEX_LONG_CMP_OP)
	state_map['&'] = byte(MY_LEX_BOOL)
	state_map['|'] = byte(MY_LEX_BOOL)
	state_map['#'] = byte(MY_LEX_COMMENT)
	state_map[';'] = byte(MY_LEX_SEMICOLON)
	state_map[':'] = byte(MY_LEX_SET_VAR)
	state_map['\\'] = byte(MY_LEX_ESCAPE)
	state_map['/'] = byte(MY_LEX_LONG_COMMENT)
	state_map['*'] = byte(MY_LEX_END_LONG_COMMENT)
	state_map['@'] = byte(MY_LEX_USER_END)
	state_map['`'] = byte(MY_LEX_USER_VARIABLE_DELIMITER)
	state_map['"'] = byte(MY_LEX_STRING_OR_DELIMITER)

	var ident_map [256]byte
	for i := 0; i < 256; i++ {
		ident_map[i] = byte((state_map[i] == MY_LEX_IDENT || state_map[i] == MY_LEX_NUMBER_IDENT))
	}

	state_map['x'] = byte(MY_LEX_IDENT_OR_HEX)
	state_map['X'] = byte(MY_LEX_IDENT_OR_HEX)
	state_map['b'] = byte(MY_LEX_IDENT_OR_BIN)
	state_map['B'] = byte(MY_LEX_IDENT_OR_BIN)
	state_map['n'] = byte(MY_LEX_IDENT_OR_NCHAR)
	state_map['N'] = byte(MY_LEX_IDENT_OR_NCHAR)

	cs.IdentMap = ident_map
	cs.StateMap = state_map
}

func isalpha(cs *CharsetInfo, c byte) bool {
	return cs.CType[c+1] & (_MY_U | _MY_L)
}

func isdigit(cs *CharsetInfo, c byte) bool {
	return cs.CType[c+1] & _MY_U
}

func isspace(cs *CharsetInfo, c byte) bool {
	return cs.CType[c+1] & _MY_SPC
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
