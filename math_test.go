package vector

import "testing"

func TestVector_DotIntVector(t *testing.T) {

	var (
		v1 = NewFromInt([]int64{4, 5, 2, 5, -1})
		v2 = NewFromInt([]int64{3, 4, -4, 1, 9})
		v3 = NewFromInt([]int64{3, 1, 4, 8, 7, 1, 0, -2})
	)

	res, err := v1.DotIntVector(v2)
	if err != nil {
		t.Errorf(err.Error())
	}
	if res != 20 {
		t.Errorf("%d != %d", res, 20)
	}
	_, err = v1.DotIntVector(v3)
	if err == nil {
		t.Errorf("Vector.DotIntVector doesn't raise error with different length vectors.")
	}

}

func TestVector_DotFloatVector(t *testing.T) {

	var (
		v1 = NewFromFloat([]float64{63.1, -5, 31.4, 3174.31, 531})
		v2 = NewFromFloat([]float64{31.43, 4.31, -4.34, 1.13, 64})
		v3 = NewFromFloat([]float64{331.4, -1.334, 4.1541, 34.14, 84.1, 79.31, 21.3, 43.1})
	)

	res, err := v1.DotFloatVector(v2)
	if err != nil {
		t.Errorf(err.Error())
	}
	if res != 39396.3773 {
		t.Errorf("%f != %f", res, 39396.3773)
	}
	_, err = v1.DotFloatVector(v3)
	if err == nil {
		t.Errorf("Vector.DotFloatVector doesn't raise error with different length vectors.")
	}

}

func TestVector_DotIntScalar(t *testing.T) {
	v1 := NewFromInt([]int64{4, 3, 5, 1, -5})
	res := v1.DotIntScalar(3)
	for index, item := range res.Range() {
		if toCompare := v1.At(uint64(index)).(int64) * 3; item != toCompare {
			t.Errorf("%d != %d", item, toCompare)
		}
	}
}

func TestVector_DotFloatScalar(t *testing.T) {
	v1 := NewFromFloat([]float64{0.4, 3.31, 5.34, 1.136, -5.031})
	res := v1.DotFloatScalar(3.2)
	for index, item := range res.Range() {
		if toCompare := v1.At(uint64(index)).(float64) * 3.2; item != toCompare {
			t.Errorf("%f != %f", item, toCompare)
		}
	}
}
