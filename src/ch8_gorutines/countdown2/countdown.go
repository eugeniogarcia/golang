// Countdown implements the countdown for a rocket launch.
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	//Creamos un canal. Cuando publiquemos un mensaje en el canal, sabremos que tenemos que abortar
	abort := make(chan struct{})

	//Lanza la go rutina
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		//Publica un token en el abort channel
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown.  Press return to abort.")

	//Un multiplexador por el canal, escucha por varios canales a la vez
	select {
	//Con After se retorna un canal en el que tras 10 segundos se publicara un Timer - uno solo, una sola vez
	case <-time.After(10 * time.Second):
		// Do nothing.
	//Si nos llega algo por este canal
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}
	launch()
}

func launch() {
	fmt.Println("Lift off!")
}
