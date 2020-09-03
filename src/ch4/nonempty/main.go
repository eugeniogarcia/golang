// Nonempty is an example of an in-place slice algorithm.
package main

import "fmt"

// nonempty returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func main() {
	//Define un slice
	data := []string{"one", "", "three"}
	//Muestra el string con quotes
	fmt.Printf("%q\n", nonempty(data)) // `["one" "three"]`
	fmt.Printf("%q\n", data)           // `["one" "three" "three"]`
}

//Toma un slice y retorna otro
func nonempty2(strings []string) []string {

	out := strings[:0] // zero-length slice of original
	for _, s := range strings {
		if s != "" {
			//Añade al slice el string
			out = append(out, s)
		}
	}
	return out
}
