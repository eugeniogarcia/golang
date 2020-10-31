// Package methods provides a function to print the methods of any value.
package methods

import (
	"fmt"
	"reflect"
	"strings"
)

//Print Toma un objeto genérico y devuelve sus métodos
func Print(x interface{}) {
	//Obtiene el reflect.Value
	v := reflect.ValueOf(x)
	//Obtenemos el tipo
	t := v.Type()
	fmt.Printf("type %s\n", t)

	//con NumMethod tenemos el número de métodos. Con Method accedemos al método en si. Podemos ver su tipo, y su nimbre
	for i := 0; i < v.NumMethod(); i++ {
		methType := v.Method(i).Type()
		fmt.Printf("func (%s) %s%s\n", t, t.Method(i).Name,
			strings.TrimPrefix(methType.String(), "func"))
	}
}
