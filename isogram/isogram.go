package isogram

import (
	"unicode"
)

// IsIsogram determine if a word or phrase is an isogram
func IsIsogram(s string) bool {

	var arr []rune

	for _, v := range s {

		if v == 45 || v == 32 {
			continue
		}

		uv := unicode.ToUpper(v)
		if !search(arr, uv) {
			arr = append(arr, uv)
		} else {
			return false
		}
	}
	return true
}

func search(arr []rune, s rune) bool {

	for i := 0; i < len(arr); i++ {
		if arr[i] == s {
			return true
		}
	}

	return false
}
