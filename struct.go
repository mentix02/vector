package vector

// Vector is a generic data type inspired by C++'s STL vector container -
//https://en.cppreference.com/w/cpp/container/vector.
type Vector struct {
	Last  *interface{} `json:"last"`
	First *interface{} `json:"first"`

	Data []interface{} `json:"data"`

	Length   uint64 `json:"length"`
	Capacity uint64 `json:"capacity"`
}

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

// NewFromSlice creates a vector and populates it from a provided slcie.
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
