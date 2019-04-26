package etl

import (
	"strings"
)

// Transform - transforms map[int][]string to map[string]int
func Transform(inputMap map[int][]string) map[string]int {

	x := make(map[string]int)

	for point, arrString := range inputMap {

		for i := 0; i < len(arrString); i++ {

			letter := strings.ToLower(arrString[i])
			x[letter] = point
		}

	}

	return x
}
