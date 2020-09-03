package eval

// An Expr is an arithmetic expression.
//Definimos un tipo como interface. Una expresión se define como un tipo que implementa estos dos métodos
type Expr interface {
	// Eval returns the value of this Expr in the environment env.
	Eval(env Env) float64
	// Check reports errors in this Expr and adds its Vars to the set.
	Check(vars map[Var]bool) error
}

//Definimos cinco tipos que expondremos en nuestro paquete
// Var
// literal
// unary
// binary
// call

// A Var identifies a variable, e.g., x.
type Var string

// A literal is a numeric constant, e.g., 3.141.
type literal float64

// A unary represents a unary operator expression, e.g., -x.
//El tipo incluye una expresión
type unary struct {
	op rune // one of '+', '-'
	x  Expr
}

// A binary represents a binary operator expression, e.g., x+y.
//El tipo incluye dos expresiones
type binary struct {
	op   rune // one of '+', '-', '*', '/'
	x, y Expr
}

// A call represents a function call expression, e.g., sin(x).
//El tipo incluye un slice de expresiones
type call struct {
	fn   string // one of "pow", "sin", "sqrt"
	args []Expr
}
