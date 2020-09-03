// The du1 command computes the disk usage of the files in a directory.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	// Procesa los flags de entrada. En este ejemplo no hemos definido ningun flag. En la version du2 ya si definimos un flag
	flag.Parse()
	//En roots tenemos un slice de strings donde se guardan todos los argumentos del programa que no sean flags
	roots := flag.Args()
	//Si no hemos especificado nada como argumento, toma el directorio actual por defecto
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Crea un canal que recibe un entero
	fileSizes := make(chan int64)
	//Lanza una go rutina
	go func() {
		for _, root := range roots {
			//Para cada directorio llamamos a esta función pasando el canal como argumento
			walkDir(root, fileSizes)
		}
		//Cerramos el canal
		close(fileSizes)
	}()

	// Print the results.
	var nfiles, nbytes int64
	//Usa range para procesar continuamente el canal, como si fuera un stream. Cuando cerremos el canal con close, se terminara este loop
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
//Estamos especificando que el argumento es un canal, y especificamente un canal al que podremos enviar datos, no recibirlos
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
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}
