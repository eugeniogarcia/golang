// Netflag demonstrates an integer type used as a bit field.
package main

import (
	"fmt"
	. "net"
)

//Toma la constante FlagUp y hace un AND bit a bit y compara el resultado con la constante FlagUp
func IsUp(v Flags) bool { return v&FlagUp == FlagUp }

//No retorna nada, pero actualiza la variable a la que hace referencia el puntero
//El contenido del puntero es igual a aplicar el operando & bit a bit entre el complemento de v y la constante FlagUp
func TurnDown(v *Flags) { *v &^= FlagUp }

//No retorna nada, pero actualiza la variable a la que hace referencia el puntero
//Bit a bit hace un or entre v y la constante FlagBroadcast
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }
func IsCast(v Flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 }

func main() {
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10001 true"
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10000 false"
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))   // "10010 false"
	fmt.Printf("%b %t\n", v, IsCast(v)) // "10010 true"
}

//!-
