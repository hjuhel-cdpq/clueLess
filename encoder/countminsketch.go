package encoder

import (
	"errors"
	"math"
)

// Export the interface of all CountMinSketch
type CountMinSketchTable interface {
	Push(signature []uint64) (interface{}, error) // Represent the increase int the counter, the first return parameter's type is the same as the one the structure is written for
}

// Constants exported by the package for creating tables
type tableSize string

const (
	UINT8 tableSize = "uint8"
)

// Create a new CountMinSketchTable with types set based on the provided size
func NewCountMinSketchTable(size uint64, counterType tableSize) (table CountMinSketchTable, err error) {

	switch counterType {
	case "uint8":

		t := make([][]uint8, size)
		for i := range t {
			t[i] = make([]uint8, 255) // All stored key will be modulate by 255
		}

		table = &countMinSketchTable8{
			table: t,
			size:  size,
		}

	default:
		table = nil
		err = errors.New("No matching table structure for provided counter type")
	}

	return
}

// Represent a count min sketch table with counter of size 8
type countMinSketchTable8 struct {
	table [][]uint8
	size  uint64
}

// Implement the Add operation for the countMinSketch 8 table
func (t countMinSketchTable8) Push(signature []uint64) (minCount interface{}, err error) {

	if uint64(len(signature)) != t.size {
		err = errors.New("Inconsitent signature size between the set to add and the table size.")
		return
	}

	var minCount_ uint8 = math.MaxUint8 // Shadow variable : to specify the emtpy interface
	for k, v := range signature {
		v = v % 255     // Ensure that the value is compatible with the size of the array
		t.table[k][v]++ // Increment the counter for this specific value of the key
		if minCount_ > t.table[k][v] {
			minCount_ = t.table[k][v] // Set the min counter to the new minimal value
		}
	}

	minCount = minCount_

	return
}
