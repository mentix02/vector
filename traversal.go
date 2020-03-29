package vector

import (
	"errors"
	"fmt"
)

// Range returns data sliced until the length of the list. Traversal
// or any other 'read' operations should not interact with Vector.Data directly.
func (v *Vector) Range() []interface{} {
	return v.Data[:v.Len()]
}

// At returns the value stored at the index provided.
func (v *Vector) At(index uint64) interface{} {
	return v.Data[index]
}

// SafeAt implements an elementary check for out of range exception
// and returns an error if an index is called for access.
func (v *Vector) SafeAt(index uint64) (interface{}, error) {
	if index >= v.Len() {
		return nil, errors.New("vector index out of range")
	}
	return v.Data[index], nil
}

// Contains performs a simple linear search on the data to find the item
// provided in argument and returns true if found else false.
func (v *Vector) Contains(toFind interface{}) bool {
	for _, item := range v.Range() {
		if item == toFind {
			return true
		}
	}
	return false
}

// Index is similar to Contains as in it performs as linear search
// but instead of returning a boolean, it returns the index if it's
// found else returns an error.
func (v *Vector) Index(toFind interface{}) (uint64, error) {
	for index, item := range v.Range() {
		if item == toFind {
			return uint64(index), nil
		}
	}
	return 0, fmt.Errorf("'%v' not found in vector.", toFind)
}
