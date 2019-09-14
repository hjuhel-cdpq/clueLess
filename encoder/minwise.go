package encoder

import (
	"errors"
	"math"
)

// Represent a 64b hash function
type Hash64Function func(s []byte) uint64

// Represent a collections of minimum hashes for a set
type MinWise struct {
	minimums []uint64
	h1       Hash64Function
	h2       Hash64Function
}

// Getter for set signature
func (m *MinWise) GetSignature() []uint64 {
	return m.minimums
}

// Generate a new collections of minimum hashes for a set
func NewMinWise(h1, h2 Hash64Function, size int) *MinWise {

	minimums := make([]uint64, size)
	for i := range minimums {
		minimums[i] = math.MaxUint64
	}

	return &MinWise{
		h1:       h1,
		h2:       h2,
		minimums: minimums,

		
	}

}

// Add a new item to the set and update the signature
func (m *MinWise) Push(b []byte) {

	v1 := m.h1(b)
	v2 := m.h2(b)

	for i, v := range m.minimums {
		hv := v1 + uint64(i)*v2
		if hv < v {
			m.minimums[i] = hv
		}
	}

}

// Compute similarity with another set
func (s *MinWise) Similarity(t *MinWise) (sim float32, err error) {

	if len(s.minimums) != len(t.minimums) {
		err = errors.New("Inconsitent signature size between sets.")
		return
	}

	var intersection float32 = 0
	for k, v := range s.minimums {
		if v == t.minimums[k] {
			intersection++
		}
	}

	sim = intersection / float32(len(s.minimums))
	return
}
