package encoder

import (
	"errors"
	"sort"
	"sync"
)

// Export the interface of all CountMinSketch
type CountMedianSketchTable interface {
	Push(signature []uint64) (interface{}, error) // Represent the increase int the counter, the first return parameter's type is the same as the one the structure is written for
}

// Represent a count min sketch table with counter of size 8
// Synchronisation is done through mutex, because there is not utin8 atomic operations
type CountMedianSketchTable8 struct {
	table [][]uint8
	size  uint64
	sync.RWMutex
}

// Implement the Add operation for the countMinSketch 8 table
func (t *CountMedianSketchTable8) Push(signature []uint64) (median interface{}, err error) {

	if uint64(len(signature)) != t.size {
		err = errors.New("Inconsitent signature size between the set to add and the table size.")
		return
	}

	counter := make([]uint8, t.size) // Holds the state of the counter

	t.Lock() // Lock the counter during the increment
	for k, v := range signature {
		v = v % 255     // Ensure that the value is compatible with the size of the array
		t.table[k][v]++ // Increment the counter for this specific value of the key
		counter[k] = t.table[k][v]
	}
	t.Unlock()

	// Compute the median of the counter
	sort.Slice(counter, func(i, j int) bool { return counter[i] < counter[j] })
	if t.size%2 != 0 {
		median = counter[(t.size-1)/2+1]
	} else {
		median = (counter[(t.size)/2-1] + counter[(t.size)/2+1]) / 2
	}

	return
}

// Implement the new table of size 8
func NewCountMedianSketchTable8(size uint64) *CountMedianSketchTable8 {

	t := make([][]uint8, size)
	for i := range t {
		t[i] = make([]uint8, 255)
	}

	return &CountMedianSketchTable8{
		table: t,
		size:  size,
	}
}
