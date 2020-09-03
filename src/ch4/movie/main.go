// Movie prints Movies as JSON.
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//Definimos una estructura. Con vistas a los jsons, estamos indicando que el campo Year debe serializarse ene el json como released. El campo Color se serializara como color, y además indicamos que solo se incluirá en el json cuando el campo este informado
type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

//slice de Movie, con 3 items
var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	// ...
}

//!-

func main() {
	{
		//Creamos un json a partir del slice
		data, err := json.Marshal(movies)
		if err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		}
		fmt.Printf("%s\n", data)
	}

	{
		//Usamos MarshalIndent para que se formate bonito
		data, err := json.MarshalIndent(movies, "", "    ")
		if err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		}
		fmt.Printf("%s\n", data)

		//Unmarshal. Creamos un slice a partir de un json. Solo tomamos el campo Title del json
		var titles []struct{ Title string }
		if err := json.Unmarshal(data, &titles); err != nil {
			log.Fatalf("JSON unmarshaling failed: %s", err)
		}
		fmt.Println(titles) // "[{Casablanca} {Cool Hand Luke} {Bullitt}]"
	}
}
