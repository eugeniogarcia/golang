// Pipeline3 demonstrates a finite 3-stage pipeline
// with range, close, and unidirectional channel types.
package main

import "fmt"

//Funcion que tiene un channel que envia datos - sender
func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		//Envia datos al channel
		out <- x
	}
	close(out)
}

//Funcion que tiene un channel que envia datos - sender -, y una channel que recibe datos - receiver
func squarer(out chan<- int, in <-chan int) {
	//Recive datos...
	for v := range in {
		//...y los envia
		out <- v * v
	}
	close(out)
}

//Funcion que recive datos - receiver
func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	//Crea dos canales
	naturals := make(chan int)
	squares := make(chan int)

	//pasa los channels como argumentos
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

//!-
