// Sorting sorts a music playlist into a variety of orders.
package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

//Tipo custom que define una estructura
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

//Un slice de punteros a Track. Inicializamos  el slice con cuatro Tracks
var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

//Convierte de string a Duration
func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		//Aborta con panico
		panic(s)
	}
	return d
}

//Toma un slice de punteros a Track
func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	//Crea un puntero a tabwriter.Writer, y llama a un metodo llamado Init que inicializa el puntero y lo devuelve
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

//Define un tipo que corresponde a un slice de punteros a Track
type byArtist []*Track

//Añade tres métodos a nuestro tipo. Los tres tienen como receiver un byArtist
func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//Idem al anterior. Tiene los mismos métodos, pero implementados de forma diferente
type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func main() {
	fmt.Println("byArtist:")
	//Convierte el slice de punteros a track a un byArtist - solo cambia el tipo. byArtist ya implementa el interface, con lo cual podemos usarlo con sort.Sort
	sort.Sort(byArtist(tracks))
	printTracks(tracks)

	fmt.Println("\nReverse(byArtist):")
	sort.Sort(sort.Reverse(byArtist(tracks)))
	printTracks(tracks)

	fmt.Println("\nbyYear:")
	sort.Sort(byYear(tracks))
	printTracks(tracks)

	fmt.Println("\nCustom:")
	//Creamos un objeto del tipo customSort - ver la definición del tipo más abajo. Lo inicializamos con nuestro slice al puntero de Tracks, de modo que usamos los mismos datos que en los ejemplos anteriores, pero además pasamos una función
	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return false
	}})
	printTracks(tracks)
}

//Similar a los tipos anteriores, solo que en este caso añadimos también un miembro con una función. Al instanciar un objeto de este tipo podremos especificar esta funcion, de modo que el interface se comporte de una u otra forma
type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

//De nuevo implementamos en el tipo los métodos que se precisa para el interface. Notese que hacemos uso del método less, con minuscula, definido en el tipo customSort
func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

func init() {
	//!+ints
	values := []int{3, 1, 4, 1}
	fmt.Println(sort.IntsAreSorted(values)) // "false"
	sort.Ints(values)
	fmt.Println(values)                     // "[1 1 3 4]"
	fmt.Println(sort.IntsAreSorted(values)) // "true"
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	fmt.Println(values)                     // "[4 3 1 1]"
	fmt.Println(sort.IntsAreSorted(values)) // "false"
	//!-ints
}
