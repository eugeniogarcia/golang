// Package memo provides a concurrency-unsafe
// memoization of a function of type Func.
package memo

// A Memo caches the results of calling a Func.
//Tenemos la función y la cache
type Memo struct {
	f     Func
	cache map[string]result
}

// Func is the type of the function to memoize.
//Define una función. Retorna un interface sin métodos, vamos, cualquier cosa
type Func func(key string) (interface{}, error)

//La respuesta que se cachea
type result struct {
	value interface{}
	err   error
}

//New crea un *Memo. Funcion que toma una funcion como argumento y retorna un puntero a Memo
func New(f Func) *Memo {
	//Iniciamos un Memo, y retorna su dirección, el puntero
	//El Memo se crea con la función, y con un cache - en blanco
	return &Memo{f: f, cache: make(map[string]result)}
}

// NOTE: not concurrency-safe!
//Busca en la cache la respuesta
//Extiende el tipo Memo con este método, que nos permite recuperar una key de la cache del Memo
func (memo *Memo) Get(key string) (interface{}, error) {
	//Buscamos el key en la cache
	res, ok := memo.cache[key]
	if !ok {
		//Si el valor no está en la cache, hacemos la llamada...
		res.value, res.err = memo.f(key)
		//...y actualizamos la cache
		memo.cache[key] = res
	}
	return res.value, res.err
}
