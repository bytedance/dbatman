package sql

import (
	"fmt"
)

func TokenName(tok int) string {
	if (tok-ABORT_SYM) < 0 || (tok-ABORT_SYM) > len(MySQLToknames) {
		return fmt.Sprintf("Unknown Token:%d", tok)
	}

	return MySQLToknames[tok-ABORT_SYM]
}
