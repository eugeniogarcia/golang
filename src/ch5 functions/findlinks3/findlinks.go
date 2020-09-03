// Findlinks3 crawls the web, starting with the URLs on the command line.
package main

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, os.Args[1:])
}

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	//Crea un mapa para guardar que hemos y que no hemos procesado ya
	seen := make(map[string]bool)
	//Mientras el slice tenga contenido
	for len(worklist) > 0 {
		//Limpia el slice
		items := worklist
		worklist = nil
		//Procesa un rango. Con cada item...
		for _, item := range items {
			//Verifica si le hemos visto. Notese que si la key no existia, no pasa nada
			if !seen[item] {
				seen[item] = true
				//Añade el item, crawled, a la worklist
				//Notese los ... Lo que hacen es extraer los items del slice que devuelve la funcion
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
