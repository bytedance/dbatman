package parser

import (
	"fmt"
)

func TokenName(tok int) string {

	if tok == 0 {
		return "EOF"
	}

	if tok > 31 && tok < 126 {
		return fmt.Sprintf("%c", tok)
	}

	if (tok-ABORT_SYM) < 0 || (tok-ABORT_SYM) > len(MySQLToknames) {
		return fmt.Sprintf("Unknown Token:%d", tok)
	}

	return MySQLToknames[tok-ABORT_SYM]
}
