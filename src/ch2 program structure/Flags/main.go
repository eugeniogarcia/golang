package main

import (
	"flag"
	"fmt"
	"strings"
)

//Define un parámetro de tipo booleano
var n = flag.Bool("n", false, "omit trailing newline")

//Define un parámetro de tipo string
var sep = flag.String("s", " ", "separator")

func main() {
	//Procesa todos los parametros
	flag.Parse()
	//Muestra todos los parametros que no sean -n y -s. Une estos parametros con el separador especificado con -n
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
