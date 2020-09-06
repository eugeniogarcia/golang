// Package bank provides a concurrency-safe bank with one account.
package bank

var balance int
var valores = make([]int, 0, 100)

//Deposit Funcion que envia por el canal
func Deposit(amount int) {
	balance = balance + amount
	valores = append(valores, balance)
}

//Balance Funcion que recive por el canal
func Balance() int {
	return balance
}

//Valores devuelve todos los saldos que ha tenido la cuenta
func Valores() []int { return valores }

//Se lanza la go rutina al iniciar el paquete
func init() {
	balance = 0
}
