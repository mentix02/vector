package vector

import "fmt"

// DotIntScalar implements a simple scalar multiplication function to
// multiply each item of a vector (that is required to be an int) with a
// scalar.
func (v *Vector) DotIntScalar(num int64) Vector {
	res := New(v.Len())
	for _, item := range v.Range() {
		res.Append(num * item.(int64))
	}
	return res
}

// DotFloatScalar is similar DotIntScalar but for floats.
func (v *Vector) DotFloatScalar(num float64) Vector {
	res := New(v.Len())
	for _, item := range v.Range() {
		res.Append(num * item.(float64))
	}
	return res
}

// DotIntVector computes the product of each individual item of
// the vectors provided and returns the sum of all the products.
func (v *Vector) DotIntVector(ov Vector) (int64, error) {
	if v.Len() != ov.Len() {
		return 0, fmt.Errorf("mismatch of sizes of vector '%d' & '%d'", v.Len(), ov.Len())
	} else {
		var total int64
		for index, item := range v.Range() {
			total += item.(int64) * ov.At(uint64(index)).(int64)
		}
		return total, nil
	}
}

// DotFloatVector is similar to DotIntVector but for floats.
func (v *Vector) DotFloatVector(ov Vector) (float64, error) {
	if v.Len() != ov.Len() {
		return 0, fmt.Errorf("mismatch of sizes of vector '%d' & '%d'", v.Len(), ov.Len())
	} else {
		var total float64
		for index, item := range v.Range() {
			total += item.(float64) * ov.At(uint64(index)).(float64)
		}
		return total, nil
	}
}
