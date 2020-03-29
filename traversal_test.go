package vector

import (
	"testing"

	. "github.com/mentix02/gopi"
)

func TestVector_Range(t *testing.T) {
	v := New(10)

	nums := PosRange(15)

	for i, num := range nums {
		nums[i] = num * num
	}

	for _, item := range nums {
		v.Append(item)
	}

	for index := range v.Range() {
		if v.At(uint64(index)) != nums[index] {
			t.Errorf("%d != %d", v.At(uint64(index)), nums[index])
		}
	}

	if *v.First != nums[0] {
		t.Errorf("%v != %v", *v.First, nums[0])
	}

	if *v.Last != nums[len(nums)-1] {
		t.Errorf("%v != %v", *v.Last, nums[len(nums)-1])
	}

}

func TestVector_SafeAt(t *testing.T) {
	v := New(10)

	for i := range []int{1, 2, 3, 4, 5, 6} {
		v.Append(i)
	}

	if item, err := v.SafeAt(8); err == nil {
		t.Errorf("error not raised and item found at v.Data[%d] = %d", 8, item)
	}

	if item, err := v.SafeAt(3); err != nil {
		t.Errorf("error raised and item not found at v.Data[%d] = %d", 8, item)
	}

}

func TestVector_Contains(t *testing.T) {
	v := NewFromSlice([]interface{}{2, 4, 6, 8})
	if v.Contains(10) {
		t.Errorf("%v should not contain %d", v.Range(), 10)
	}
	if !v.Contains(4) {
		t.Errorf("%v should contain %d", v.Range(), 4)
	}
}

func TestVector_Index(t *testing.T) {
	v := NewFromSlice([]interface{}{2, 4, 6, 8})
	if _, err := v.Index(10); err == nil {
		t.Errorf("%v should not contain %d", v.Range(), 10)
	}
	if _, err := v.Index(4); err != nil {
		t.Errorf("%v should contain %d", v.Range(), 4)
	}
}
