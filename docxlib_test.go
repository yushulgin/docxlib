package docxlib

import "testing"

func TestNumToHans(t *testing.T) {
	s := NumToHans(101)
	t.Error(s)
}
