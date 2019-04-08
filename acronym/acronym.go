package acronym

import (
	"strings"
)

var replacer = strings.NewReplacer(",", "", "'", "", "-", " ")

// Convert a phrase to its acronym
func Abbreviate(s string) string {

	res := ""
	str := replacer.Replace(s)

	words := strings.Split(str, " ")
	for _, word := range words {
		if len(word) != 0 {

			res = res + string([]rune(word)[0])
		}
	}

	res = strings.ToUpper(res)

	return res
}
