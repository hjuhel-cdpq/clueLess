package encoder

import "io"

// Represents a Shingle reader
type shingleReader struct {
	src string
	cur int
}

// Generate a new shingle reader from a string
func newShingleReader(src string) *shingleReader {
	return &shingleReader{src: src}
}

// Implements the IO.Reader interface
func (s *shingleReader) Read(p []byte) (int, error) {

	if s.cur > len(s.src) {
		return 0, io.EOF
	}

	// Compute the maximal size of the next chunk, and the required padding
	var bound int
	x := len(s.src) - s.cur
	if x >= len(p) {
		bound = len(p)
	} else {
		bound = x
	}

	// Get the next characters until reach the bound
	buffer := make([]byte, bound)
	for n := 0; n < bound; n++ {
		buffer[n] = s.src[s.cur+n]
	}

	// Update the value of the slice through the buffer
	n := copy(p, buffer)

	// Move shingle cursor to the next character
	s.cur++

	return n, nil

}
