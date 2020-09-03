package eval

import (
	"bytes"
	"fmt"
)

// Format formats an expression as a string.
// It does not attempt to remove unnecessary parens.
func Format(e Expr) string {
	var buf bytes.Buffer
	write(&buf, e)
	return buf.String()
}

func write(buf *bytes.Buffer, e Expr) {
	//Usa el tipo que implementa Expr para imprimir la información específica de cada Expr
	switch e := e.(type) {
	case literal:
		//Solo el float
		fmt.Fprintf(buf, "%g", e)

	case Var:
		//Solo el float
		fmt.Fprintf(buf, "%s", e)

	case unary:
		//Imprimimos el operando, y luego la Expr que corresponde al valor
		fmt.Fprintf(buf, "(%c", e.op)
		write(buf, e.x)
		buf.WriteByte(')')

	case binary:
		buf.WriteByte('(')
		write(buf, e.x)
		fmt.Fprintf(buf, " %c ", e.op)
		write(buf, e.y)
		buf.WriteByte(')')

	case call:
		fmt.Fprintf(buf, "%s(", e.fn)
		for i, arg := range e.args {
			if i > 0 {
				buf.WriteString(", ")
			}
			write(buf, arg)
		}
		buf.WriteByte(')')

	default:
		panic(fmt.Sprintf("unknown Expr: %T", e))
	}
}
