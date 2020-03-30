package vector

import "testing"

func TestVector_EqualsSlice(t *testing.T) {
	v := New(5)
	n := []interface{}{4, 3, 2}

	if v.EqualsSlice(n) {
		t.Errorf("%v == %v", v.Range(), n)
	}

	v.Append(2)
	v.Append(3)
	v.Append(5)

	if v.EqualsSlice(n) {
		t.Errorf("%v == %v", v.Range(), n)
	}

}

func TestVector_EqualsVector(t *testing.T) {
	var (
		v1 = NewFromInt([]int64{4, 3, 1, 6, 3})
		v2 = NewFromInt([]int64{4, 3, 1, 6, 3})
		v3 = NewFromInt([]int64{4, -3, 2, 7, 9})
	)

	if !v1.EqualsVector(v2) {
		t.Errorf("%v != %v", v1.Range(), v2.Range())
	}

	if v1.EqualsVector(v3) {
		t.Errorf("%v != %v", v1.Range(), v3.Range())
	}

	v2.Append(uint64(31))

	if v1.EqualsVector(v2) {
		t.Errorf("%v == %v", v1.Range(), v2.Range())
	}

}
