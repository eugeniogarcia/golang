// Package eval provides an expression evaluator.
package eval

import (
	"fmt"
	"math"
)

//Definimos un tipo como un mapa de keys Var, y valores float
type Env map[Var]float64

//Implementamos para nuestros tipos el metodo Eval del interface. Implementando Eval y Check hacemos que nuestro tipo sean considerados como una Expresion, tipos que implementan la interface Expr

//Eval devuelve un float64 a partir de una expresión

//Buscamos en la slice nuestro Var, y retornamos ese float
func (v Var) Eval(env Env) float64 {
	return env[v]
}

//Convierte el literal a float
func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

//Suma o resta uno a las expresiones
func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

//+, -, * o divide
func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

//Llama a la función
func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}
