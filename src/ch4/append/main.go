// Append illustrates the behavior of the built-in append function.
package main

import "fmt"

//Toma un slice como primer argumento, y un número variable de argumentos a continuación
func appendslice(x []int, y ...int) []int {
	var z []int
	//la variable y se materializara como un slice. Esto será el número total de datos
	//len nos dice el tamaño del slice
	zlen := len(x) + len(y)
	//cap es la capacidad del slice
	if zlen <= cap(x) {
		//Hemos expandido el slice x, porque podemos, la capacidad es suficiente. z y x apuntaran al mismo array subyacente
		z = x[:zlen]
	} else {
		// There is insufficient space.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		//Creamos un nuevo slice. Indicamos el tamaño y la capacidad
		z = make([]int, zlen, zcap)
		//copia el contenido de x en z
		copy(z, x)
	}
	//Copiamos la parte que nos faltaba, y
	copy(z[len(x):], y)
	return z
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow.  Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space.  Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	z[len(x)] = y
	return z
}

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d  cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}
