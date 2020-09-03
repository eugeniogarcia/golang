// Bytecounter demonstrates an implementation of io.Writer that counts bytes.
package main

import (
	"fmt"
)

//Definimos un tipo custom que tiene como underlaying uno básico. Esto nos permite comparar valores de este tipo - custom
type ByteCounter int

//Añade el método Write de modo que *ByteCounter implementa el interface
func (c *ByteCounter) Write(p []byte) (int, error) {
	//Toma un slice de bytes, mira su tamaño, y actualiza el valor del puntero
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	//Podemos usar Println con nuestro tipo, porque hemos implementado el interface
	fmt.Println(c) // "5", = len("hello")

	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")
}
