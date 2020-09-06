# Resumen

- Clock is a TCP server that periodically writes the time
- Netcat is a simple read/write client for TCP servers. We can use the previous project, `Clock`, with this client
- Reverb is a TCP server that simulates an echo.
- Pipeline demonstrates a finite 3-stage pipeline. We can see how we can reive and send data via a channel, and how that is a synchronous/blocking operation. We can also see how a channel can be closed un blocking any operation. We see how the channel may be used in a range
- Pipeline2. Simuilar to the previous case, but this time the channels are passed as arguments to the functions
- Crawl1 crawls web links starting with the command-line arguments. Shows how we can limit de degree of parallelism using a buffered channel
- Crawl2 crawls web links starting with the command-line arguments. Uses the channel with range, to sort os use the channel as an streaming source
- Countdown implements the countdown for a rocket launch. Uses a method available in Timer, `Tick(1 * time.Second)`, that returns a channel of Timer. The function publishes a value with the frequency we specify in the channel
- Countdown2. Uses another method available in Timer, `After(10 * time.Second)`, that publishes one Timer in the channel - on single time after 10 seconds in this case. This example also introduces the `select` instruction to multiplex listening in several channels
- Countdown3. It is an hybrid of Countdown1 & Countdown2. We see the use of `select` as multiplexer, but with a channel that uses `Tick(1 * time.Second)`. Each time we receive a message in select, the select is processed. If we need to listen for another message, we need to run the select again:

```go
tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			// Do nothing.
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
    }
```

- The du1 command computes the disk usage of the files in a directory. Usa `range` para procesar continuamente un canal, como si fuera un stream:

```go
for size := range fileSizes {
    nfiles++
    nbytes += size
}
```

Este procesamiento terminara cuando cerremos el canal:

```go
close(fileSizes)
```

Otra cosa a destacar es que al definir el argumento de una funcion con un canal, podemos concretar y definir el canal como de recepcion o de envio. En este caso, de envio:

```go
func walkDir(dir string, fileSizes chan<- int64) {
```

- The du2 command computes the disk usage of the files in a directory. Improves du1 by adding a flag that allows us to periodically display the progress of the work. We create a new channel that will receive data every second, and the use a select to multiplex the two channels, the one that receives this pulse, and the one that receives the actual updates in the number of files processed:

```go
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
```

Se muestra tambien que cuando recibimos datos de un canal, en realidad recibimos una dupla. El segundo campo de la dupla nos dice si el canal esta abierto.

- The du3 command computes the disk usage of the files in a directory. In this version we are going to limit the degree of paralellism
    - We can syncronize different go rutines, so that we can, for example, close the channel when all the go rutines have finished. For that we use `sync.WaitGroup`:

    ```go
    var n sync.WaitGroup
    ```

    We can Add one or more to the count:

    ```go
    n.Add(1)
    ```

    We can substract - one by one:

    ```go
    n.Done()
    ```

    And finally we can synchronize the execution. Here we are waiting for all the go rutines to finish before closing the channel:

    ```go
    go func() {
    //Esperamos a que el contador llegue a cero para seguir. Sera señal de que no haya más go rutinas ejecutandose, y que podemos cerrar el canal
    n.Wait()
    close(fileSizes)
    }()
    ```

    - We show in this example how to limit the number of go rutines working in parallel. For that we used a buffered channel:

    ```go
    var sema = make(chan struct{}, 20)
    ```

    We have to wrap the payload processing:

    ```go
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
    ```

    Notice here the use of `defer` and how we request and release a token from the channel before and after completing the processing

- The du4 variant includes cancellation. 

We create the channel we are going to use to notify a cancellation, and a helper that will use it:

```go
//Crea un canal asociado a una struct
var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}
```

We can use the channel to know when we have to abort. In this multiplexer we add the cancelation channel, and we drain the channels:

```go
loop:
	//!+3
	for {
		select {
		case <-done:
			// Drain fileSizes to allow existing goroutines to finish.
			for range fileSizes {
				// Do nothing.
			}
			return
		case size, ok := <-fileSizes:
			// ...
			//!-3
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
    }
```

Here we use the helper:

```go
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	if cancelled() {
		return
    }
```

Finally, when we want to signal a cancelation, we just close the channel:

```go
close(done)
```

- Package cake provides a simulation of a concurrent cake shop with numerous parameters. Shows:

	- Use of a Normal distribution to create a random delay

	```go
	delay := d + time.Duration(rand.NormFloat64()*float64(stddev))
	```

	- Define buffered channels

	```go
	baked := make(chan cake, s.BakeBuf)
	```

	- Use channels as arguments to functions. Specify if we want to use the channel as sender or receiver 

	```go
	func (s *Shop) icer(iced chan<- cake, baked <-chan cake)
	```

- Chat is a server that lets clients chat with each other. Shows nice things, like the fact that the type send or received in a channel, can be another channel. It uses a socket server and go rutines to implemente a service that will broadcast to all the clients, what it is being sent by one of the clients