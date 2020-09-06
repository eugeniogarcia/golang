package main

import (
	"fmt"
	"sync"

	"ch9/bank"
	bank1 "ch9/bank1"
	bank2 "ch9/bank2"
	bank3 "ch9/bank3"
)

func main() {
	fmt.Println("Sin definir una zona critica")
	resp := banco()
	fmt.Println("Centralizando todo el acceso en un singleton")
	resp1 := banco1()
	fmt.Println("Emulando un mutex")
	resp2 := banco2()
	fmt.Println("Usando un mutex")
	resp3 := banco3()
	for i := 0; i < max(len(resp), len(resp1), len(resp2), len(resp3)); i++ {
		fmt.Printf("%d\t%d\t%d\t%d\t%d\n", i, resp[i], resp1[i], resp2[i], resp3[i])
	}
}

func banco() []int {
	var n sync.WaitGroup
	for b, i := true, 0; i < 100; i++ {
		n.Add(1)
		go func(b bool) {
			if b {
				bank.Deposit(50)
			} else {
				bank.Deposit(-10)
			}
			n.Done()
		}(b)
		b = !b
	}

	n.Wait()
	fmt.Printf("Saldo final: %d\n", bank.Balance())
	return bank.Valores()
}

func banco1() []int {
	var n sync.WaitGroup

	for b, i := true, 0; i < 100; i++ {
		n.Add(1)
		go func(b bool) {
			if b {
				bank1.Deposit(50)
			} else {
				bank1.Deposit(-10)
			}
			n.Done()
		}(b)
		b = !b
	}

	n.Wait()
	fmt.Printf("Saldo final: %d\n", bank1.Balance())
	return bank1.Valores()
}

func banco2() []int {
	var n sync.WaitGroup

	for b, i := true, 0; i < 100; i++ {
		n.Add(1)
		go func(b bool) {
			if b {
				bank2.Deposit(50)
			} else {
				bank2.Deposit(-10)
			}
			n.Done()
		}(b)
		b = !b
	}

	n.Wait()
	fmt.Printf("Saldo final: %d\n", bank2.Balance())
	return bank2.Valores()
}

func banco3() []int {
	var n sync.WaitGroup

	for b, i := true, 0; i < 100; i++ {
		n.Add(1)
		go func(b bool) {
			if b {
				bank3.Deposit(50)
			} else {
				bank3.Deposit(-10)
			}
			n.Done()
		}(b)
		b = !b
	}

	n.Wait()
	fmt.Printf("Saldo final: %d\n", bank3.Balance())
	return bank3.Valores()
}

func max(x ...int) int {
	i := x[0]
	for j := 1; j < len(x); j++ {
		if x[j] > i {
			i = x[j]
		}
	}
	return i
}
