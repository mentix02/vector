package vector

import (
	"testing"

	"github.com/mentix02/gopi"
)

func TestVector_Len(t *testing.T) {

	v := New(0)

	if v.Len() != 0 {
		t.Errorf("%d != 0", v.Len())
	}

	for _, i := range gopi.PosRange(5) {
		v.Append(i)
	}

	if v.Len() != 5 {
		t.Errorf("%d != 5", v.Len())
	}

}
