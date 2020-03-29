package vector

import "testing"

func TestVector_Equals(t *testing.T) {
	v := New(5)
	n := []interface{}{4, 3, 2}

	if v.EqualsSlice(n) {
		t.Errorf("%v == %v", v.Data, n)
	}

	v.Append(2)
	v.Append(3)
	v.Append(5)

	if v.EqualsSlice(n) {
		t.Errorf("%v == %v", v.Data, n)
	}

}
