// Coloredpoint demonstrates struct embedding.
package main

import (
	"fmt"
	"image/color"
	"math"
)

//Definimos un tipo llamado Point
type Point struct{ X, Y float64 }

//Definimos un segundo tipo, que incluye un componente de tipo Point. Como no le damos nombre el campo se llamara Point, y es de tipo Point
type ColoredPoint struct {
	Point
	Color color.RGBA
}

//Definimos un método para el tipo Point
func (p Point) Distance(q Point) float64 {
	dX := q.X - p.X
	dY := q.Y - p.Y
	return math.Sqrt(dX*dX + dY*dY)
}

//Definimos un método para el tipo Point, pero el receiver es un puntero. Esto nos permitira que con el método cambiemos el valor de nuestra variable
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	//Creamos dos variables de tipo ColoredPoint y las inicializamos
	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}

	//ColoredPoint no tiene el metodo Distance, pero el compilador ve que Point si lo tiene, asi que hace la operación con el campo Point de nuestro tipo ColoredPoint
	fmt.Println(p.Distance(q.Point)) // "5"

	//Como en el caso anterior el compilador utilizara el método de Point con el campo Point de ColoredPoint. El compilador sabe que el receiver es un puntero, así que le pasara la dirección del Point
	p.ScaleBy(2)
	q.ScaleBy(2)

	//Equivalente a lo anterior. En este caso aplicamos el metodo sobre el campo Point de ColoredPoint
	var p1 = ColoredPoint{Point{1, 1}, red}
	p1.Point.ScaleBy(2)

	//Equivalente a lo anterior. . En este caso aplicamos el metodo sobre un Point
	var p2 = Point{1, 1}
	p2.ScaleBy(2)

	//Equivalente a lo anterior. En este caso aplicamos el metodo sobre un puntero de Point
	var p4 = Point{1, 1}
	var p3 = &p4
	p3.ScaleBy(2)

	fmt.Println(p.Distance(q.Point))  // "10"
	fmt.Println(p1.Distance(q.Point)) // "10"
	fmt.Println(p2.Distance(q.Point)) // "10"
	fmt.Println(p3.Distance(q.Point)) // "10"
	fmt.Println(p4.Distance(q.Point)) // "10"
}

//Demuestra como utilizar funciones de inicializacion. Estas funciones se ejecutan antes que el resto, y sirven para inicializar
func init() {
	p := Point{1, 2}
	q := Point{4, 6}

	//Podemos acceder al método, y ponerle un alias
	distance := Point.Distance   // method expression
	fmt.Println(distance(p, q))  // "5"
	fmt.Printf("%T\n", distance) // "func(Point, Point) float64"

	scale := (*Point).ScaleBy
	scale(&p, 2)
	fmt.Println(p)            // "{2 4}"
	fmt.Printf("%T\n", scale) // "func(*Point, float64)"
	//!-methodexpr
}

//Otra funcion de inicializacion
func init() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	//Define el tipo dentro del scope de la función. Tapa al tipo definido en el Paquete
	type ColoredPoint struct {
		*Point
		Color color.RGBA
	}

	p := ColoredPoint{&Point{1, 1}, red}
	q := ColoredPoint{&Point{5, 4}, blue}

	//El tipo Point tiene el metodo Distance. El receiver el compilador es "listo" y sabe que si pasamos un ColoredPoint en realidad nos referimos al campo Point, y sabe que tenemos que pasar su dirección
	fmt.Println(p.Distance(*q.Point)) // "5"
	//Son referencias Point
	q.Point = p.Point // p and q now share the same Point
	p.ScaleBy(2)
	fmt.Println(*p.Point, *q.Point) // "{2 2} {2 2}"
}
