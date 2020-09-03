# Resumen

- sha256. Obtiene el sha256 de un array de bytes
- Rev. Reverses a slice. Este es bastante completo. Podemos ver
    - slices
        - Se definen de forma parecida a los arrays, pero no tenemos que especifica la dimensión
        - Podemos añadir items
        - Podemos gestionar 
    - Strings. Podemos hacer un split en varios strings, podemos concatenar, podemos convertir a int,...
    - En el loop vemos como manejar más de una variable a la vez, como hacer `continue` a una etiqueta, podemos revertir los valores de variables - lo que especifiquemos a la derecha del = se copia antes de hacer la asignación
- nonempty. Usa slices, demuestra como al pasarlo como argumento a una función, lo que estamos pasando es una referencia, de modo que los cambios que se hagan se trasladan fuera de la función
- append. Este ejemplo es muy interesante. 
    - Demostramos como determinar el tamaño y la capacidad de un slice:

    ```go
        var z []int
        zlen := len(x) + 1
        if zlen <= cap(x) {
    ```

    - Demuestra compo copiar datos. Hemos visto otras funciones como append, que gestionan el contenido del slice, lo amplian cuando se necesita, etc.

    ```go
    copy(z, x)
    ```

    - Nos enseña como poder crear un slice con el tamaño y capacidad deseados:

    ```go 
    z = make([]int, zlen, zcap)
    ```

    - Podemos ver tambien como crear una función en la que haya un número variable de argumentos. Este número variable de argumentos serán accesible desde la función como un slice:

    ```go
    func appendInt(x []int, y int) []int {
    ```

- Charcount. Computes counts of Unicode characters
- Embed. Demonstrates basic struct embedding
- Movie. Prints Movies as JSON
- github. Paquete que demuestra como interactuar con una api usando json
- Uso de Templates
    - Issuesreport. Creamos un listado según un template
        - El template

        ```go
        const templ = `{{.TotalCount}} issues:
        {{range .Items}}----------------------------------------
        Number: {{.Number}}
        User:   {{.User.Login}}
        Title:  {{.Title | printf "%.64s"}}
        Age:    {{.CreatedAt | daysAgo}} days
        {{end}}`
        ```

        - Definimos un listado

        ```go
        {{range .Items}}----------------------------------------
        
        ...
        
        {{end}}`
        ```

        - Aplicamos una lambda a un campo. 

        ```go
        Age:    {{.CreatedAt | daysAgo}} days
        ```

        transformaremos el dato `.CreatedAt` con `daysAgo`. Al construir el template y parsearlo, podemos indicar que funcion correspondera con la transformación `daysAgo`. En este caso la función será `haceDias`:

        ```go
        var report = template.Must(template.New("issuelist").
        Funcs(template.FuncMap{"daysAgo": haceDias}).
        Parse(templ))
        ```

        La función se define:

        ```go
        func haceDias(t time.Time) int {
            return int(time.Since(t).Hours() / 24)
        }
        ```

        - Creamos el template:

        ```go
        var report = template.Must(template.New("issuelist").
        Funcs(template.FuncMap{"daysAgo": haceDias}).
        Parse(templ))
        ```

        - Ejecutamos el template. Al ejecutarlo especificamos los datos que tienen que usarse:

        ```go
        report.Execute(os.Stdout, result)
        ```
    
    - issueshtml. Similar al ejemplo anterior. En este caso el template tiene formato html. Se crea el template como antes:

    ```go
    var issueList = template.Must(template.New("issuelist").Parse(
    ```

    - Autoescape. Demonstrates automatic HTML escaping in html/template. En este caso vamos a asegurarnos que los datos que pasemos al template tengan el escape necesario para poder usar el resultado del template en un payload a enviar via http. El template se define de la misma forma que hemos visto antes:

    ```go
    const templ = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
    t := template.Must(template.New("escape").Parse(templ))
    ```

    La diferencia esta al definir los datos que vamos a utilizar para resolver el template. Con el tipo `template.HTML` vamos a forzar a que los datos que se informen en la variable sean automáticamente "escapados": 

    ```go
    	var data struct {
		A string        
		B template.HTML 
    }
    ```

    El template se resuelve como siempre:

    ```go
    t.Execute(os.Stdout, data)
    ```