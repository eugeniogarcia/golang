// Package cake provides a simulation of
// a concurrent cake shop with numerous parameters.
//
// Use this command to run the benchmarks:
// 	$ go test -bench=. gopl.io/ch8/cake
package cake

import (
	"fmt"
	"math/rand"
	"time"
)

//Define y exporta un tipo
type Shop struct {
	Verbose        bool
	Cakes          int           // number of cakes to bake
	BakeTime       time.Duration // time to bake one cake
	BakeStdDev     time.Duration // standard deviation of baking time
	BakeBuf        int           // buffer slots between baking and icing
	NumIcers       int           // number of cooks doing icing
	IceTime        time.Duration // time to ice one cake
	IceStdDev      time.Duration // standard deviation of icing time
	IceBuf         int           // buffer slots between icing and inscribing
	InscribeTime   time.Duration // time to inscribe one cake
	InscribeStdDev time.Duration // standard deviation of inscribing time
}

//Tipo interno. No se exporta
type cake int

//Metodo que no se exporta
//Toma un canal de emision y lo aplica sobre un Shop
func (s *Shop) baker(baked chan<- cake) {
	//Para cada Cake del shop
	for i := 0; i < s.Cakes; i++ {
		//Convierte a tipo cake
		c := cake(i)
		if s.Verbose {
			fmt.Println("baking", c)
		}
		//elabora el cake (bake)
		work(s.BakeTime, s.BakeStdDev)
		//Informa el canal con el cake
		baked <- c
	}
	//Cierra el canal. Este baker no va a emitir nada más
	close(baked)
}

//Metodo que tampoco se exporta. Tenemos dos canales como argumentos, el primero es de emisión, el segundo es de recepción
func (s *Shop) icer(iced chan<- cake, baked <-chan cake) {
	//Un stream con el canal baked
	for c := range baked {
		if s.Verbose {
			fmt.Println("icing", c)
		}
		//elabora el cake (icing)
		work(s.IceTime, s.IceStdDev)
		//Informa el canal con el cake
		iced <- c
	}
}

//Metodo que tampoco se exporta. Tenemos un canal de recepcion como argumento
func (s *Shop) inscriber(iced <-chan cake) {
	//Para cada cake de la tienda
	for i := 0; i < s.Cakes; i++ {
		//esperamos a que se haya hecho el icing. Recivimos un cake que tiene hecho el icing
		c := <-iced
		if s.Verbose {
			fmt.Println("inscribing", c)
		}
		//hace el inscribing
		work(s.InscribeTime, s.InscribeStdDev)
		if s.Verbose {
			fmt.Println("finished", c)
		}
	}
}

// Work runs the simulation 'runs' times.
//Este método se exporta. Sobre una determinada tienda hace un número de simulaciones. Cada simulacion lanza una gorutina que se encarga
func (s *Shop) Work(runs int) {
	for run := 0; run < runs; run++ {
		//Crea un buffered channel con un tamaño s.BakeBuf
		baked := make(chan cake, s.BakeBuf)
		iced := make(chan cake, s.IceBuf)
		go s.baker(baked)
		for i := 0; i < s.NumIcers; i++ {
			go s.icer(iced, baked)
		}
		s.inscriber(iced)
	}
}

//Espera un tiempo aleatorio. El tiempo de espera sigue una distribución normal de media d y desviación estandar stddev
func work(d, stddev time.Duration) {
	delay := d + time.Duration(rand.NormFloat64()*float64(stddev))
	time.Sleep(delay)
}
