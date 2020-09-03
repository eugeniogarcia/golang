package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	//Para cada url especificada en los argumentos
	for _, url := range os.Args[1:] {
		//... hace un http GET
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//...y lee toda la respuesta
		b, err := ioutil.ReadAll(resp.Body)
		//y cierra el stream
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
