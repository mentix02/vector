package vector

import "errors"

// Pop removes the last element from a vector.
func (v *Vector) Pop() (interface{}, error) {
	if v.Len() < 1 {
		return nil, errors.New("pop from empty vector")
	} else {
		item := v.Data[v.Len()-1]
		v.Data[v.Len()-1] = nil
		v.Length--
		if v.Len() == 0 {
			v.Last = nil
		} else {
			v.Last = &v.Data[v.Len()-1]
		}
		return item, nil
	}
}

// Clear deleted all items from a Vector and
// allocates Vector.Data to new 0 byte long memory chunk.
func (v *Vector) Clear() {
	v.Last, v.First = nil, nil
	v.Length, v.Capacity = 0, 0
	v.Data = make([]interface{}, 0)
}

// Remove removes the first occurrence of item provided.
func (v *Vector) Remove(item interface{}) {
	index, err := v.Index(item)
	if err == nil {
		copy(v.Data[index:], v.Data[index+1:])
		v.Data[v.Len()-1] = nil
		v.Length--
		v.Data = v.Data[:v.Capacity - 1]
	}
}
