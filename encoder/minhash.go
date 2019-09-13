package encoder

import (
	"sort"

	"github.com/tenfyzhong/cityhash"
)

// Represent a 64b hash function
type Hash64Function func(s []byte) uint64

// Build a hash function from a seed
func BuildHash64Function(seed uint64) Hash64Function {

	curryHash := func(s []byte) uint64 {
		hash := cityhash.CityHash64WithSeed(s, seed)
		return hash
	}

	return curryHash
}

// Compute the signature of a slice of
func Minhash(s [][]byte, h Hash64Function, size uint8) []uint64 {

	// Compute all hashes
	hashs := make([]uint64, len(s))
	for i, item := range s {
		hashs[i] = h(item)
	}

	// Select the size first value
	sort.Slice(hashs, func(i, j int) bool { return hashs[i] < hashs[j] })
	signature := hashs[0:size]

	return signature

}
