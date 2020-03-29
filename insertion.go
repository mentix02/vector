package vector

// Append add a new member to the end of a vector. It reallocates memory if the vector is filled
// and updates (and manages) the length and capacity automatically.
func (v *Vector) Append(item interface{}) {

	if !v.full() {
		v.Data[v.Len()] = item
	} else {
		largerCap := largerCapacity(v.Capacity)
		data := make([]interface{}, largerCap)
		copy(data, v.Data)
		v.Data = data
		v.Capacity = largerCap
		v.Append(item)
		return
	}

	if v.Len() == 1 {
		v.First = &v.Data[0]
		v.Last = &v.Data[0]
	} else {
		v.Last = &v.Data[v.Len()]
	}

	v.Length++

}

// Insert inserts an item into a provided index while taking care of resizing and capacity management.
func (v *Vector) Insert(index uint64, item interface{}) {

	if (index == v.Len() && index <= v.Capacity) || index > v.Len() {
		v.Append(item)
		return
	}

	if v.Len()+1 < v.Capacity {
		v.Length++
		for i := v.Len(); i > index; i-- {
			v.Data[i] = v.At(i - 1)
		}
		v.Data[index] = item
	} else {
		largerCap := largerCapacity(v.Capacity)
		largerData := make([]interface{}, largerCap)

		copy(largerData, v.Data[0:index])
		largerData[index] = item
		for i := index + 1; i <= v.Len(); i++ {
			largerData[i] = v.At(i - 1)
		}

		v.Length++
		v.Data = largerData
		v.Capacity = largerCap
	}

	if index == 0 {
		v.First = &v.Data[0]
	}

}
