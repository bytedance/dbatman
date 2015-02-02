package charset

import (
	"testing"
)

func TestUtf8(t *testing.T) {

	func() { // TEST for utf8 digit
		for i := 0; i < 10; i++ {
			b := byte('0') + byte(i)
			if CSUtf8GeneralCli.IsDigit(b) == false {
				t.Fatalf("%v is not digit type", b)
			}
		}
	}()

	func() { // TEST for utf8 digit
		for i := 0; i < 26; i++ {
			b := byte('A') + byte(i)
			if CSUtf8GeneralCli.IsAlpha(b) == false {
				t.Fatalf("%v is not digit type", b)
			}

			b = byte('a') + byte(i)
			if CSUtf8GeneralCli.IsAlpha(b) == false {
				t.Fatalf("%v is not digit type", b)
			}
		}
	}()
}
