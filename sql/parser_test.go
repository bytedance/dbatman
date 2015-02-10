package sql

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	if tree, err := Parse("Select version()"); err != nil {
		t.Fatalf("%v", err)
	} else {
		fmt.Println(tree)
	}
}
