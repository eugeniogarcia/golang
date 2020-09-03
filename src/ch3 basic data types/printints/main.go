// Printints demonstrates the use of bytes.Buffer to format a string.
package main

import (
	"bytes"
	"fmt"
)

//!+
// intsToString is like fmt.Sprint(values) but adds commas.
func intsToString(values []int) string {
	//Define un Buffer de bytes
	var buf bytes.Buffer
	//Escribe en el buffer un byte
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			//escribe en el buffer un String
			buf.WriteString(", ")
		}
		//Añadimos un string al buffer
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	//Convertimos el buffer a String
	return buf.String()
}

func main() {
	fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"
}

//!-
