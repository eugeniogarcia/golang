// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package memo_test

import (
	"testing"

	memo "ch9/memo1"
	"ch9/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	//Decidimos cachear la llamada al método httpGetBody
	m := memo.New(httpGetBody)
	//Memo implementa una interface que nos permite pasarlo como segundo argumento, y hacer la llamada secuencial
	memotest.Sequential(t, m)
}

// NOTE: not concurrency-safe!  Test fails.
func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	//Memo implementa una interface que nos permite pasarlo como segundo argumento, y hacer la llamada secuencial
	memotest.Concurrent(t, m)
}
