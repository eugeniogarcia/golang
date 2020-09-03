package eval

import (
	"fmt"
	"strings"
)

// Implementa el método Check del interface, para los cinco tipos
// La funcion toma un mapa que tiene como clave Var, y guarda un boleano. La funcion devuelve un error
// Vamos a implementar la funcion en todos los tipos que hemos creado en ast.go

//Aplicada sobre un Var informa el mapa, añade la Var y la pone como true. Al final tendremos el mapa actualizado
func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}

//Aplicada sobre un literal siempre retorna nil
func (literal) Check(vars map[Var]bool) error {
	return nil
}

//Aplicada sobre un unary, si el operando del receiver no es + o - no pasa la validación. Si el operando es un + o un -, tiene que pasar la validación de x. x es cualquier expresion - Var, literal, unary, ...
func (u unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("unexpected unary op %q", u.op)
	}
	return u.x.Check(vars)
}

//Aplicada sobre un binary, si el operando del receiver no es +, -, * o / ono pasa la validación. Si el operando pasa la validación, se tiene que pasar también la validación de los argumentos. Los argumentos son también expresiones
func (b binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/", b.op) {
		return fmt.Errorf("unexpected binary op %q", b.op)
	}
	if err := b.x.Check(vars); err != nil {
		return err
	}
	return b.y.Check(vars)
}

var numParams = map[string]int{"pow": 2, "sin": 1, "sqrt": 1}

//Aplicada sobre una funcion comprobamos que la funcion sea bien pow, sin o sqrt
//Comprueba si el tampaño del slice de argumentos coincide con lo definido para la funcion
//Comprueba que cada argumento sea una expresión válida
func (c call) Check(vars map[Var]bool) error {
	arity, ok := numParams[c.fn]
	//Esta la funcion en nuestro mapa?
	if !ok {
		return fmt.Errorf("unknown function %q", c.fn)
	}
	//Comprueba si el tampaño del slice de argumentos coincide con lo definido para la funcion
	if len(c.args) != arity {
		return fmt.Errorf("call to %s has %d args, want %d",
			c.fn, len(c.args), arity)
	}
	//Procesa el slice, comprobando que cada argumento sea una expresión válida
	for _, arg := range c.args {
		if err := arg.Check(vars); err != nil {
			return err
		}
	}
	return nil
}
