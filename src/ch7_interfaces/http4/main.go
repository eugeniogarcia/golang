// Http4 is an e-commerce server that registers the /list and /price
// endpoint by calling http.HandleFunc.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	//Especifica para cada uri el handler que debe usarse
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	//Empieza a servir peticiones
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//Tipo que incluye un método para convertir a String. Esti hace que se implemente el interface Stringer
type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

//Define un tipo. Se trata de un mapa que devuelve el tipo dollars para entradas de tipo string
type database map[string]dollars

//Metodo list. Retorna todo el contenido del mapa
func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

//Metodo price
func (db database) price(w http.ResponseWriter, req *http.Request) {
	//Toma el parametro item del query parameter
	item := req.URL.Query().Get("item")
	//Si el item esta en el mapa, lo retorna
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		//Responde con la cabecera http status code 404
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}
