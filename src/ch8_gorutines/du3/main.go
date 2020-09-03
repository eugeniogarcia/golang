// The du3 command computes the disk usage of the files in a directory.
package main

// The du3 variant traverses all directories in parallel.
// It uses a concurrency-limiting counting semaphore
// to avoid opening too many files at once.

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var vFlag = flag.Bool("v", false, "show verbose progress messages")

//!+
func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse each root of the file tree in parallel.
	fileSizes := make(chan int64)

	//Este tipo nos permite sincronizar varias go rutinas
	var n sync.WaitGroup

	for _, root := range roots {
		//Incrementamos el contador para indicar que hay una go rutina más en ejecución
		n.Add(1)
		//Pasamos una referencia al semaforo
		go walkDir(root, &n, fileSizes)
	}
	go func() {
		//Esperamos a que el contador llegue a cero para seguir. Sera señal de que no haya más go rutinas ejecutandose, y que podemos cerrar el canal
		n.Wait()
		close(fileSizes)
	}()

	// Print the results periodically.
	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}

	printDiskUsage(nfiles, nbytes) // final totals
}

//!-

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			//Incrementa el contador para indicar que tenemos otra go rutina trabajando
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			//Es recursivo
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// sema is a counting semaphore for limiting concurrency in dirents.
//Vamos a establecer un límite en el número de go rutinas trabajando a la vez. Usamos un buffered channel con tamaño 20
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	//Añadimos un token al buffer. Si ya hubiera 20, se bloqueara hasta que se libere un slot
	sema <- struct{}{} // acquire token
	//Con el defer nos aseguramos que siempre, incluso si hay un error, se saque un token del buffer, liberando así un slot
	defer func() { <-sema }() // release token

	//Hace el trabajo
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
