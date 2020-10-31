package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//Demuestra como acceder a un campo de una variable
	acceso()

}

//Demuestra como acceder a un campo de una variable
func acceso() {
	//Definimos una variable de un tipo compuesto
	var x struct {
		a bool
		b int16
		c []int
	}

	//con uintptr hacemos un cast a un entero. Este entero tiene el tamaño suficiente para almacenar una direccion
	//con unsafe.Pointer obtenemos la dirección, un puntero
	//con unsafe.Offsetof el offset de un campo en un tipo complejo, respecto a la direccion base

	/*
		obtiene la direccion base de x, hacemos el cast a uintptr para poder sumarle el offset del campo. El resultado es un puntero, al que hacemos un cast a un puntero de go a entero de 16 bits
	*/
	pb := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	//Cambiamos el valor con el puntero
	*pb = 42
	//Verificamos que efectivamente hemos accedido al campo b de la estructura
	fmt.Println(x.b) // "42"
}

/*
trocear el código no seria correcto, porque el garbage collector no identifica tmp como un puntero. Esto significa que si el garbage collector reorganizara la memoria, no sabria que tiene que actualizar el valor de tmp

tmp := uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)
	pb := (*int16)(unsafe.Pointer(tmp))
	*pb = 42
*/
