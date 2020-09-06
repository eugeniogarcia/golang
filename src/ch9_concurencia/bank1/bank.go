// Package bank provides a concurrency-safe bank with one account.
package bank

//Creamos dos canales para solicitar escrituras y lecturas
var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var valores = make([]int, 0, 100)

//Deposit Funcion que envia por el canal
func Deposit(amount int) { deposits <- amount }

//Balance Funcion que recive por el canal
func Balance() int { return <-balances }

//Valores devuelve todos los saldos que ha tenido la cuenta
func Valores() []int { return valores }

//Esta funcion es la única que tiene acceso a las variables, de modoque se evitan inconsistencias. Quien necesite actualizar o leer tendrá que usar los canales para contactar con esta funcion. Viene a ser un singleton
func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		//lee del canal deposits
		case amount := <-deposits:
			balance += amount
			valores = append(valores, balance)
		//Escribe en el canal balance
		case balances <- balance:
		}
	}
}

//Se lanza la go rutina al iniciar el paquete
func init() {
	go teller() // start the monitor goroutine
}
