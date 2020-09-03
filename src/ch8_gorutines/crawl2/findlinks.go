// Crawl3 crawls web links starting with the command-line arguments.
//
// This version uses bounded parallelism.
// For simplicity, it does not address the termination problem.
//
package main

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	//Creamos un canal bidireccional con slices de strings
	worklist := make(chan []string) // lists of URLs, may have duplicates
	//Creamos un canal para intercambiar strings
	unseenLinks := make(chan string) // de-duplicated URLs

	//Lanzamos una go rutina con una función anónima que alimenta el canal con un slice de strings
	go func() { worklist <- os.Args[1:] }()

	for i := 0; i < 20; i++ {
		//Lanza 20 go rutinas
		go func() {
			//Usamos range con el canal. Lo que estamos haciendo aqui es usar el canal como un stream de datos por el que vamos recibiendo strings
			for link := range unseenLinks {
				foundLinks := crawl(link)
				//Lanzamos una go rutuna para alimentar el canal con el slice obtenido por crwal
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	seen := make(map[string]bool)
	//Usamos el canal worklist como una fuente de streams de slices de strings
	for list := range worklist {
		//iteramos sobre todos los componentes del slice
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				//Alimenta el canal
				unseenLinks <- link
			}
		}
	}
}
