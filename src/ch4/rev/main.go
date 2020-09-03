// Rev reverses a slice.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Declaramos un array con una dimension igual a los datos con lo que se inicializa, 6 en este caso
	a := [...]int{0, 1, 2, 3, 4, 5}
	//Pasamos un slice del array
	reverse(a[:])
	fmt.Println(a) // "[5 4 3 2 1 0]"

	//Declaramos un slice. El slice tiene un array por debajo
	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s) // "[2 3 4 5 0 1]"

	// Interactive test of reverse.
	input := bufio.NewScanner(os.Stdin)
outer:
	for input.Scan() {
		//Slice de enteros
		var ints []int
		//Con Fields se sapara el String usando el espacio como separador. Creamos un range
		for _, s := range strings.Fields(input.Text()) {
			//Convierte cada valor en int
			x, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				//Continua desde la etiqueta outer
				continue outer
			}
			//Añade al slice el int
			ints = append(ints, int(x))
		}
		reverse(ints)
		fmt.Printf("%v\n", ints)
	}
	// NOTE: ignoring potential errors from input.Err()
}

// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		//Revierte el valor. Procesa el slice desde el principio y el fin a la vez
		s[i], s[j] = s[j], s[i]
	}
}
