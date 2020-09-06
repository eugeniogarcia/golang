// Package bank provides a concurrency-safe single-account bank.
package bank

import "sync"

var (
	mu      sync.Mutex // guards balance
	balance int
)

var valores = make([]int, 0, 100)

func Valores() []int { return valores }

func Deposit(amount int) {
	mu.Lock()
	balance = balance + amount
	valores = append(valores, balance)
	mu.Unlock()
}

func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}
