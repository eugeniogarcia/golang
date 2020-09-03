// Countdown implements the countdown for a rocket launch.
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Commencing countdown.")
	//Esta funcion devuelbe un channel, especificamente un sender, que envia Time. Cuando leamos de este channel leeremos un Time. La particularidad es que se publicara un Timer cada segundo
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		//Leemos un valor del canal. Como se publica el valor una vez por segundo, en esencia ponemos una espera por segundo
		<-tick
	}
	launch()
}

func launch() {
	fmt.Println("Lift off!")
}
