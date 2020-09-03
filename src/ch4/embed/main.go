// Embed demonstrates basic struct embedding.
package main

import "fmt"

//Define un tipo
type Point struct{ X, Y int }

//Define un tipo compuesto.
type Circle struct {
	Point
	Radius int
}

//Define un tipo compuesto.
type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w Wheel
	//!+
	//Al especificar los valores, NO estamos indicando el key, así que tenemos que especificar los valores en el mismo orden en que se han definido
	w = Wheel{Circle{Point{8, 8}, 5}, 20}

	//Al especificar los valores, estamos indicando el key y el valor
	w = Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20, // NOTE: trailing comma necessary here (and at Radius)
	}

	fmt.Printf("%#v\n", w)
	// Output:
	// Wheel{Circle:Circle{Point:Point{X:8, Y:8}, Radius:5}, Spokes:20}

	w.X = 42

	fmt.Printf("%#v\n", w)
	// Output:
	// Wheel{Circle:Circle{Point:Point{X:42, Y:8}, Radius:5}, Spokes:20}
	//!-
}
