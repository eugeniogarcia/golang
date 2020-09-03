// Findlinks2 does an HTTP GET on each URL, parses the
// result as HTML, and prints the links within it.
//
// Usage:
//	findlinks url ...
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			continue
		}
		//Imprimimos cada link encontrado
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

// findLinks performs an HTTP GET request for url, parses the
// response as HTML, and extracts and returns the links.
func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	//Si hubo un error
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	//Parseamos el documento html
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		//Devolvemos nil en los slices, y un mensaje de error
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	//Pasamos nil como slice. El slice esta vacio. El segundo argumento es el documento html parseado. Lo que buscamos es obtener todos los links. Devolvemos el slice con los links y el código de error
	return visit(nil, doc), nil
}

// visit appends to links each link found in n, and returns the result.
func visit(links []string, n *html.Node) []string {
	//Busca nodos de tipo a
	if n.Type == html.ElementNode && n.Data == "a" {
		//Recorre todos los atributos buscando el href
		for _, a := range n.Attr {
			if a.Key == "href" {
				//Actualiza el slice, añadiendo el href
				links = append(links, a.Val)
			}
		}
	}
	//Recorre todos los children. n es un puntero
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		//recursivo
		links = visit(links, c)
	}
	//retorna el slice
	return links
}
