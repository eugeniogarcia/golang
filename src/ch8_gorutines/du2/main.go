// The du2 command computes the disk usage of the files in a directory.
package main

// The du2 variant uses select and a time.Ticker
// to print the totals periodically if -v is set.

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

//Definimos un flag, v, que sera de tipo Bool, tiene por defecto el valor false. Al llamar al programa podremos pasar -v=true por ejemplo
var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	// Procesa los flags de entrada. En este ejemplo tenemos el flag -v. El puntero verbose tomara valor
	flag.Parse()
	//En roots tenemos un slice de strings donde se guardan todos los argumentos del programa que no sean flags
	roots := flag.Args()
	//Si no hemos especificado nada como argumento, toma el directorio actual por defecto
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse the file tree.
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	// Print the results periodically.
	//Creamos un canal, un canal que emite datos, y los datos son de tipo time.Time
	var tick <-chan time.Time
	//Comprobamos el valor que hemos pasado en el flag -v
	if *verbose {
		//Recibe cada 500ms un Time
		tick = time.Tick(500 * time.Millisecond)
	}

	var nfiles, nbytes int64
loop:
	for {
		//Escucha en varios canales
		select {
		//Con cada update de datos actualizamos los contadores. Demuestra aqui como al escuchar de un canal lo que recibimos en realidad es una dupla. El segundo campo de la dupla nos dice si el canal esta abierto
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles++
			nbytes += size
		//Con cada tick mostramos los datos
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
func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
