package vector

import (
	"testing"
)

func TestVector_Pop(t *testing.T) {
	s := []int64{3, 1, 4, 5, 3}
	v := NewFromInt(s)

	for i := len(s) - 1; i >= 0; i-- {
		if last, err := v.Pop(); last != s[i] && err != nil {
			t.Errorf("%d != %d", last, s[i])
		}
	}

	item, err := v.Pop()
	if err == nil {
		t.Errorf("empty Vector returned value: %v", item)
	}

}

func TestVector_Clear(t *testing.T) {
	v := NewFromInt([]int64{3, 1, 5, 6, 8})
	v.Clear()
	if v.Len() != 0 || v.Capacity != 0 || len(v.Data) != 0 {
		t.Errorf("v was not cleared")
	}
}

func TestVector_Remove(t *testing.T) {
	s := []int64{8, 4, 3, 1, 9, 3}
	v := NewFromInt(s)
	cv := NewFromInt(s)
	v.Remove(int64(4))
	if v.EqualsVector(cv) {
		t.Errorf("%v == %v", v.Range(), cv.Range())
	}
}
