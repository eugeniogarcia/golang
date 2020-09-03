// Pipeline2 demonstrates a finite 3-stage pipeline.
package main

import "fmt"

func main() {
	//Define dos canales que intercambian int
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 100; x++ {
			//Envia el valor por el canal
			naturals <- x
		}
		//Cierra el canal
		close(naturals)
	}()

	// Squarer
	go func() {
		//Obtiene el valor desde el canal
		for x := range naturals {
			//Envia el valor por el canal
			squares <- x * x
		}
		//Cierra el canal
		close(squares)
	}()

	// Printer (in main goroutine)
	//escribe el valor obteneido desde el canal
	for x := range squares {
		fmt.Println(x)
	}
}

//!-
