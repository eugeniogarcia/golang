package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//Crea el mapa
	counts := make(map[string]int)
	//Crea un rango con los argumentos del programa...
	for _, filename := range os.Args[1:] {
		//... y abre un archivo para cada uno
		data, err := ioutil.ReadFile(filename)
		//Si hay un error
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		//Accede a un rango
		//
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

//!-
