// Package vector provides a strong dynamic array interface for storing and updating data efficiently.
package vector

func (v *Vector) full() bool {
	return v.Length == v.Capacity
}

// New creates a new vector with NULL values.
func New(capacity uint64) Vector {
	largerCap := largerCapacity(capacity)
	v := Vector{Capacity: largerCap, Length: 0}
	v.Data = make([]interface{}, largerCap)
	return v
}

// NewFromSlice creates a vector and populates it from a provided slice.
func NewFromSlice(slice []interface{}) Vector {
	v := New(uint64(len(slice)))
	for _, item := range slice {
		v.Append(item)
	}
	return v
}

// Len returns the total number of elements present in a vector.
func (v *Vector) Len() uint64 {
	return v.Length
}

// NewFromInt constructs a Vector from a given slice of int64s.
func NewFromInt(slice []int64) Vector {
	v := New(uint64(len(slice)))
	for _, item := range slice {
		v.Append(item)
	}
	return v
}

// NewFromFloat similar to NewFromInt returns a Vector
// produced by a slice of float64s.
func NewFromFloat(slice []float64) Vector {
	v := New(uint64(len(slice)))
	for _, item := range slice {
		v.Append(item)
	}
	return v
}
