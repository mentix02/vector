package vector

func largerCapacity(capacity uint64) uint64 {
	return capacity + (capacity * 2) + 1
}
