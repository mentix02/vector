package vector

// Equals checks for equality between two
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
