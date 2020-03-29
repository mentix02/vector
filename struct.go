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
