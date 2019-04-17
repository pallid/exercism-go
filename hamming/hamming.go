// Package hamming provides methods counting 'Hamming distance'.
package hamming

import "errors"

// Distance finds The Hamming distance between two DNA strand.
// Distance returns result as a int.
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, errors.New("The Hamming distance is only defined for sequences of equal length")
	}

	if a == "" {
		return 0, nil
	}

	distance := 0

	ra := []rune(a)
	for i, sb := range b {
		if ra[i] != sb {
			distance++
		}
	}

	return distance, nil
}
