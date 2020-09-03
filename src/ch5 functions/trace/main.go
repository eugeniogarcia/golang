// The trace program uses defer to add entry/exit diagnostics to a function.
package main

import (
	"log"
	"time"
)

func bigSlowOperation() {
	log.Printf("inicio")
	//trace retorna una funcion. Lo que estamos haciendo es ejecutar la funcion trace inmediatamente, y la funcion que devuelve, la diferimos
	defer trace("bigSlowOperation")() // don't forget the extra parentheses

	time.Sleep(10 * time.Second) // simulate slow operation by sleeping
	log.Printf("fin")
}

func trace(msg string) func() {
	//Calculamos el momento de inicio
	start := time.Now()
	log.Printf("enter %s", msg)

	//Como la función se difiere, esta funcion anónima se ejecutara cuando la función principal termine
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

//!-main

func main() {
	bigSlowOperation()
}

/*
!+output
$ go build gopl.io/ch5/trace
$ ./trace
2015/11/18 09:53:26 enter bigSlowOperation
2015/11/18 09:53:36 exit bigSlowOperation (10.000589217s)
!-output
*/
