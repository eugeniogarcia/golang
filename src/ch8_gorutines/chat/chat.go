// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

//!+broadcaster
type client chan<- string // an outgoing message channel

var (
	//Canal de con un canal de recepción
	entering = make(chan client)
	//Canal de con un canal de recepción
	leaving = make(chan client)
	//Canal de strings
	messages = make(chan string) // all incoming client messages
)

//Ejecutado en una go rutina que se lanza al abrir el servidor de sockets
func broadcaster() {
	//Mapa con clientes conectados. El tipo del key es un tipo custom, client, que tiene como underlying un canal de envio para strings
	clients := make(map[client]bool) // all connected clients

	//De forma indefinida
	for {
		//Escucha
		select {
		//Un mensaje recibido en cualquiera de los clientes
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli <- msg
			}

		//Que recibamos algo por aqui significa que un cliente se ha conectado. Lo que recibimos es el canal por que podemos enviar mensajes
		case cli := <-entering:
			//Actualizamos el mapa
			clients[cli] = true

		//Que recibamos algo por aqui significa que un cliente se ha desconectado. Lo que recibimos es el canal por que podemos enviar mensajes
		case cli := <-leaving:
			//Borramos la entrada del mapa
			delete(clients, cli)
			//Cerramos el canal. Ya no podremos enviar mensajes
			close(cli)
		}
	}
}

//Gestiona cada conexión
func handleConn(conn net.Conn) {
	//Crea un canal de strings
	ch := make(chan string) // outgoing client messages
	//Lanza una gorutina que se encargara de enviar al cliente todo lo que se publique en el canal ch
	go clientWriter(conn, ch)

	//Obtiene la direccion del cliente
	who := conn.RemoteAddr().String()
	ch <- "You are " + who

	//Todo lo que pongamos en este canal, se envia al cliente
	messages <- who + " has arrived"

	//Enviamos un canal por el canal
	entering <- ch

	//Procesa la información enviada desde el cliente
	input := bufio.NewScanner(conn)
	for input.Scan() {
		//Publicamos en el canal lo que hemos recibido de este cliente
		messages <- who + ": " + input.Text()
	}

	//Cuando no haya más que leer en el canal de entrada
	//Enviamos un canal por el canal
	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

//Envia al cliente todo lo que llegue por el canal ch
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func main() {
	//Escucha con un socket server en 8000
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	//Lanza una go rutina
	go broadcaster()
	for {
		//Acepta una conexión de un cliente
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		//Gestiona la conexión por separado en una gorutina
		go handleConn(conn)
	}
}
