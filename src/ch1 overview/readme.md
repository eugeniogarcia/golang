# Resumen

- dup2. Usa `bufio.NewScanner(f)`, `.Scan()`, y `.Text()`
- dup3. Abre y cierra archivos. Usa `string.Split`
- fetch. Usa `net/http` para hacer una petición http GET y leer la respuesta
- fecthall. Demuestra como construir concurrencia. Utiliza la librería `time` además de acceder a http Get. También usa `ioutil.Discard` en `io.Copy`. No nos interesa tanto copiar la respuesta como saber que cantidad de bytes tiene  
- lissajous. Demuestra como construir un gif animado. Usa `image`. También usa `math` para obtener números aleatorios
-server1. Demuestra como crear un servidor http usando `net/http`
-server2. Similar al caso anterior, pero usa también la librería `sync` para definir un mutex y sincronizar el acceso a una variable. Usa `.Lock()` y `.Unlock()`
-server3. Demuestra como acceder a las propiedades de la request: cacebeceras, url, método.