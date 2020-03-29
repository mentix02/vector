package vector

import (
	"testing"

	. "github.com/mentix02/gopi"
)

func TestVector_Insert(t *testing.T) {

	v := New(0)

	for _, item := range PosRange(5) {
		v.Append(item + 1)
	}

	v.Append(uint(7))
	v.Insert(5, uint(6))

	for index, item := range PosRange(7) {
		if v.At(uint64(index)) != item+1 {
			t.Errorf("%d != %d", v.At(uint64(index)), item+1)
		}
	}

}

func TestVector_Insert2(t *testing.T) {
	v := New(0)

	for num := 5; num > 0; num-- {
		v.Insert(0, num)
	}

	equals := []interface{}{1, 2, 3, 4, 5}

	if !v.EqualsSlice(equals) {
		t.Errorf("%v != %v", v.Range(), equals)
	}

}
