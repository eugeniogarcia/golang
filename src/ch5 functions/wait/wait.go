// The wait program waits for an HTTP server to start responding.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: wait url\n")
		os.Exit(1)
	}
	url := os.Args[1]

	if err := WaitForServer(url); err != nil {
		fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		os.Exit(1)
	}
}

// WaitForServer attempts to contact the server of a URL. It tries for one minute using exponential back-off. It reports an error if all attempts fail.
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	//Añadimos al ts actual el timeout que queremos fijar
	deadline := time.Now().Add(timeout)
	//Mientras sigamos dentro del deadline, aumentamos los intentos
	for tries := 0; time.Now().Before(deadline); tries++ {
		//Llamamos
		_, err := http.Head(url)
		if err == nil {
			return nil // success
		}
		log.Printf("server not responding (%s); retrying...", err)
		//Interesante. esperamos 1000, 10000, 100000, ...
		time.Sleep(time.Second << uint(tries)) // exponential back-off
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
