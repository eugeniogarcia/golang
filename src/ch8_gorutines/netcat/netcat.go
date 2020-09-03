// Netcat is a simple read/write client for TCP servers.
package main

import (
	"io"
	"log"
	"net"
	"os"
)

//!+
func main() {
	//Se conecta con un servidor tcp
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	//Crea un channel
	done := make(chan struct{})
	//Lanza esta go rutine
	go func() {
		//Compia en la salida lo que hayamos recibido desde el servidor
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		//Cuando no haya más que recibir, escribe done
		log.Println("done")
		//Escribe en el canal
		done <- struct{}{} // signal the main goroutine
	}()

	//Envia datos al servidor
	mustCopy(conn, os.Stdin)
	//Cierra la conexión
	conn.Close()
	//Se bloquea hasta no recibir datos en el canal
	<-done // wait for background goroutine to finish
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
