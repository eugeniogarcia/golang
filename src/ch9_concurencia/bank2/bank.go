// Package bank provides a concurrency-safe bank with one account.
package bank

var (
	sema    = make(chan struct{}, 1) // a binary semaphore guarding balance
	balance int
)
var valores = make([]int, 0, 100)

func Valores() []int { return valores }

func Deposit(amount int) {
	sema <- struct{}{} // acquire token
	balance = balance + amount
	valores = append(valores, balance)
	<-sema // release token
}

func Balance() int {
	sema <- struct{}{} // acquire token
	b := balance
	<-sema // release token
	return b
}
