# Resumen

- The `wait` program waits for an HTTP server to start responding. WaitForServer attempts to contact the server of a URL. It tries for one minute using exponential back-off. It reports an error if all attempts fail. It is __interesting__ the way we set the time-out de espera
- `Findlinks1` prints the links in an HTML document read from standard input. We have a function that returns an slice, the slice is used with a range in a loop. We recursivelly call the function. We use __html package__ to parse an html document and navigate through its elements
- `Findlinks2` does an HTTP GET on each URL, parses the result as HTML, and prints the links within it. It shows recursive functions
- `Findlinks3` crawls the web, starting with the URLs on the command line. Shows
    - Create a map
    - Function that returns an slice, and the results are passed as variable number of arguments with the `...` operand
- The `sum` program demonstrates a variadic function
- `Outline` prints the outline of an HTML document tree. Uses the __html package__ in a simular fashion used in the previous example
- The `trace` program uses __defer__ to add entry/exit diagnostics to a function. We can also see an __anoymous function__
- Defer1 demonstrates a deferred call being invoked during a panic. We show recurrent functions, and defer, and how the defer functions are run when the function exits, and they are run in the reverse order in which they were initially called


