package vector

// EqualsSlice checks for equality between two slices.
func (v *Vector) EqualsSlice(arr []interface{}) bool {

	if v.Len() != uint64(len(arr)) {
		return false
	}

	for index, item := range v.Range() {
		if arr[index] != item {
			return false
		}
	}

	return true
}

// EqualsVector checks for equality between two vectors.
func (v *Vector) EqualsVector(ov Vector) bool {

	if v.Len() != ov.Len() {
		return false
	}

	for index, item := range v.Range() {
		if ov.At(uint64(index)) != item {
			return false
		}
	}

	return true
}
