# Resumen

- surface. Demuestra como crear un grafico svg de una curva
- mandelbrot. Demuestra como utilizar el package `Image` para crear un grafico. Usa el tipo `complex`, structs como `Color`.
- basename2. Usa métodos de `String` como `strings.LastIndex`. Usa slices. Usa `Scan()` del paquete `bufio`.
- printints. Demuestra como usar un byte Buffer para construir un String. Los Strings son inmutables, pero con Buffer podemos ir componiendo un String
- netflag. Demuestra
    - `dot imports`. Con dot imports establecemos que un determinado paquete que estamos importando puede utilizarse sin especificar el prefijo. No es una práctica recomendada, de echo se esta proponiendo eliminar del leguaje

```go
import (
    "fmt"
    . "net"
)
```

    - Uso de operaciones bit a bit sobre un integer
    - uso de `ioat` para definir constantes

```go
const (
    FlagUp           Flags = 1 << iota // interface is up
    FlagBroadcast                      // interface supports broadcast access capability
    FlagLoopback                       // interface is a loopback interface
    FlagPointToPoint                   // interface belongs to a point-to-point link
    FlagMulticast                      // interface supports multicast access capability
)
``` 

Es equivalente a:

```go
const (
    FlagUp           Flags = 1 << 0 // interface is up
    FlagBroadcast    Flags = 1 << 1 // interface supports broadcast access capability
    FlagLoopback     Flags = 1 << 2 // interface is a loopback interface
    FlagPointToPoint Flags = 1 << 3 // interface belongs to a point-to-point link
    FlagMulticast    Flags = 1 << 4 // interface supports multicast access capability
)
```