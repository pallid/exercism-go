package accumulate

type fn func(string) string

// Accumulate implements the accumulate operation, which,
// given a collection and an operation to perform
// on each element of the collection, returns a new collection containing
// the result of applying that operation to each element
// of the input collection
func Accumulate(input []string, converter fn) []string {

	var collection []string

	for _, word := range input {
		collection = append(collection, converter(word))
	}

	return collection

}
