// Package memo provides a concurrency-safe memoization a function of
// a function.  Requests for different keys proceed in parallel.
// Concurrent requests for the same key block until the first completes.
// This implementation uses a Mutex.
package memo

import "sync"

// Func is the type of the function to memoize.
type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

//Creamos un nuevo tipo. El tipo aumenta lo que hasta ahora era la respuesta de la función, para añadir un canal
type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]*entry
}

func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		//Cuando el valor no esta en la cache, creamos una entrada en la cache, una entrada que apenas tendrá el canal, no hay valor
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()
		//Hacemos la llamada a la función
		e.res.value, e.res.err = memo.f(key)
		//Cerramos el canal. INTERESANTE, PORQUE TODOS LOS QUE ESTUVIERAN BLOQUEADOS ESPERANDO LA LLEGADA DE DATOS POR ESTE CANAL SE DESBLOQUEAN
		close(e.ready) // broadcast ready condition
	} else {
		// This is a repeat request for this key.
		memo.mu.Unlock()
		//Aunque hemos encontrado una entrada en la cache, puede ser la entrada que solo tiene el canal. Nos aseguramos que este relleno
		//SI EL CANAL NO ESTUVIERA CERRADO, ESPERAMOS. SI EL CANAL YA SE HUBIERA CERRADO, CONTINUA. QUE SE HAYA CERRADO SIGNIFICA QUE LA LLAMADA A LA FUNCION HA CONCLUIDO
		<-e.ready // wait for ready condition
	}
	return e.res.value, e.res.err
}
