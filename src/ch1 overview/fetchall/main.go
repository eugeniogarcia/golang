package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	//Obtiene el ts actual
	start := time.Now()
	//Crea un canal que devuelve string
	ch := make(chan string)
	//Para cada una de las urls especificadas en los argumentos...
	for _, url := range os.Args[1:] {
		//llama a la rutina que hemos creado, llamada recupera. Le pasamos una url y el canal
		go recupera(url, ch) // start a goroutine
	}

	//Para cada una de las url, recuperamos la información que nos llegue por el canal, y la imprimimos
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	//Medimos cuanto hemos tardado
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func recupera(url string, ch chan<- string) {
	//Empezamos a medir
	start := time.Now()
	//Hacemos la petición http get
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	//Terminamos de medir
	secs := time.Since(start).Seconds()
	//Informamos en el canal lo que se ha tardado en obtener la respuesta, el tamaño de la respuesta, y la url
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

//!-
