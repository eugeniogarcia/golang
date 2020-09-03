# Resumen

- Bytecounter demonstrates an implementation of io.Writer that counts bytes. Basic example of the use of interfaces
- Tempflag prints the value of its -temp (temperature) flag. Shows how an interface is implemented
- Xmlselect prints the text of selected elements of an XML document.
- Sorting sorts a music playlist into a variety of orders. Very cool example of the use of Interfaces. Shows the common way in which we can sort data
- http. Demuestra como exponer uris y gestionar peticiones
    - http2. Rudimentary implementation. The handler inspects the path requested, and with a switch decides what handling to do 

    - http3a. Uses the built in `http.NewServeMux()`:

    ```go
    mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
    log.Fatal(http.ListenAndServe("localhost:8000", mux))
    ```

    With `mux.Handle` we specify the uri and the handler for that uri. The handler has to be an inteface, `http.HandlerFunc`. A function is nothing more than a type, so we can convert a function from a type to another. That is what we do with `http.HandlerFunc(db.list)`, we are converting the type `db.list` into `http.HandlerFunc`. The signature is the same, so we are in effect converting a function into a interface type

    - http3a. Uses the built in `http.NewServeMux()`, but in a more direct way. The previous example has only a pedagogic value:

    ```go
    mux := http.NewServeMux()
	mux.HandleFunc("/list", db.list)
    mux.HandleFunc("/price", db.price)
    ```

    With `mux.HandleFunc` we specify the handler.
    
    -htt4. We do not need to use the Mux directly, we can: 

    ```go
    	//Especifica para cada uri el handler que debe usarse
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	//Empieza a servir peticiones
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
    ```

- eval. Develops a quite interesting package. Has everything
    - defer
    - recover
    - handle panic, throw panic
    - create different types, with methods, ...
    - and of course, interfaces, a lot of interfaces stuff