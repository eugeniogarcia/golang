// Package memotest provides common functions for
// testing various designs of the memo package.
package memotest

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

//Hacemos un GET http
func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	//Devuelbe un array de bytes y un error. El array de bytes implementa interface
	return ioutil.ReadAll(resp.Body)
}

//Esportamos la función
var HTTPGetBody = httpGetBody

//Esta funcion devuelve un channel que emite string
func incomingURLs() <-chan string {
	//Creamos un canal
	ch := make(chan string)
	go func() {
		//iteramos sobre un range que hemos contruido con un slice de direcciones predeterminadas
		for _, url := range []string{
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
		} {
			ch <- url
		}
		//Cierra el canal, no se enviaran más datos
		close(ch)
	}()
	return ch
}

//M Exporta un tipo llamado M que define un interface
type M interface {
	Get(key string) (interface{}, error)
}

//Sequential exporta una funcion que tiene como argumento un interface M
//La funcion procesa el canal en modo stream de forma secuencial
func Sequential(t *testing.T, m M) {
	for url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}

}

//Concurrent exporta una función que nos permite a una interface M
//La funcion procesa el canal en modo stream de forma paralela
func Concurrent(t *testing.T, m M) {
	//Crea un Waitgroup
	var n sync.WaitGroup
	for url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			//Al final de la ejecución libera el grupo
			defer n.Done()
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
		}(url)
	}

	//Esperamos a que todas las peticiones terminen
	n.Wait()
}
