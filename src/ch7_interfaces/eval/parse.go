package eval

import (
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
)

// Definimos un tipo como struct de dos elementos, un scaner y un rune
type lexer struct {
	scan  scanner.Scanner
	token rune // current lookahead token
}

//añadimos al tipo varios metodos:
//next
//text
//describe

//Actualiza el lexer con el siguiente o EOF
func (lex *lexer) next() { lex.token = lex.scan.Scan() }

//Obtiene el valor de la posicion actual
func (lex *lexer) text() string { return lex.scan.TokenText() }

// describe returns a string describing the current token, for use in errors.
// Nos devuelve el estado de nuestro lexer en un texto
func (lex *lexer) describe() string {
	switch lex.token {
	//Fin de archivo
	case scanner.EOF:
		return "end of file"
	case scanner.Ident:
		return fmt.Sprintf("identifier %s", lex.text())
	case scanner.Int, scanner.Float:
		return fmt.Sprintf("number %s", lex.text())
	}
	return fmt.Sprintf("%q", rune(lex.token)) // any other rune
}

type lexPanic string

func precedence(op rune) int {
	switch op {
	case '*', '/':
		return 2
	case '+', '-':
		return 1
	}
	return 0
}

// ---- parser ----

func Parse(input string) (_ Expr, err error) {
	//Ejecuta la funcion al terminar el Parseo
	//Como en la funcion deferida queremos actualizar alguno de los datos de respuesta, en la salida de la funcion no solo especificamos el tipo, sino también le asignamos una variable. Como no trabajaremos con Expr, definidmos una variable anónima, _
	defer func() {
		//Con recover la funcion se ejecuta solo en caso de pánico.
		/*Recover is a built-in function that regains control of a panicking goroutine. Recover is only useful inside deferred functions. During normal execution, a call to recover will return nil and have no other effect. If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.*/
		//Comprobamos si estamos en panic, y si lo estamos mira que tipo de panic es
		switch x := recover().(type) {
		case nil:
			// no panic
		case lexPanic:
			err = fmt.Errorf("%s", x)
		default:
			//Relanzamos el panic
			// unexpected panic: resume state of panic.
			panic(x)
		}
	}()

	//Crea un puntero a lexer. Con esto tenemos un scanner y un token
	lex := new(lexer)

	//Inicializa el scanner, con un reader que se arma con el string. Con esto podemos ya escanear el string
	lex.scan.Init(strings.NewReader(input))

	//Indicamos que queremos escanear. Los tipos de token que vamos a tratar. Vamos a scanear enteros, float e identifiers
	lex.scan.Mode = scanner.ScanIdents | scanner.ScanInts | scanner.ScanFloats
	//Pues procesamos el primer token del string
	lex.next() // initial lookahead

	e := parseExpr(lex)

	if lex.token != scanner.EOF {
		return nil, fmt.Errorf("unexpected %s", lex.describe())
	}
	return e, nil
}

func parseExpr(lex *lexer) Expr { return parseBinary(lex, 1) }

// binary = unary ('+' binary)*
// parseBinary stops when it encounters an
// operator of lower precedence than prec1.
func parseBinary(lex *lexer, prec1 int) Expr {
	lhs := parseUnary(lex)
	for prec := precedence(lex.token); prec >= prec1; prec-- {
		for precedence(lex.token) == prec {
			op := lex.token
			lex.next() // consume operator
			rhs := parseBinary(lex, prec+1)
			lhs = binary{op, lhs, rhs}
		}
	}
	return lhs
}

// unary = '+' expr | primary
//Procesa una expresion unary
func parseUnary(lex *lexer) Expr {
	//Si el token es un + o -
	if lex.token == '+' || lex.token == '-' {
		op := lex.token
		//Toma el siguiente token
		lex.next()
		//retorna una expresion unaria
		return unary{op, parseUnary(lex)}
	}
	//Si no se trata de un unario
	return parsePrimary(lex)
}

func parsePrimary(lex *lexer) Expr {
	switch lex.token {
	//Si es un identificador
	case scanner.Ident:
		id := lex.text()
		lex.next() // consume Ident
		if lex.token != '(' {
			return Var(id)
		}
		lex.next() // consume '('
		var args []Expr
		if lex.token != ')' {
			for {
				args = append(args, parseExpr(lex))
				if lex.token != ',' {
					break
				}
				lex.next() // consume ','
			}
			if lex.token != ')' {
				msg := fmt.Sprintf("got %s, want ')'", lex.describe())
				panic(lexPanic(msg))
			}
		}
		lex.next() // consume ')'
		return call{id, args}

	case scanner.Int, scanner.Float:
		f, err := strconv.ParseFloat(lex.text(), 64)
		if err != nil {
			panic(lexPanic(err.Error()))
		}
		lex.next() // consume number
		return literal(f)

	case '(':
		lex.next() // consume '('
		e := parseExpr(lex)
		if lex.token != ')' {
			msg := fmt.Sprintf("got %s, want ')'", lex.describe())
			panic(lexPanic(msg))
		}
		lex.next() // consume ')'
		return e
	}
	msg := fmt.Sprintf("unexpected %s", lex.describe())
	panic(lexPanic(msg))
}
