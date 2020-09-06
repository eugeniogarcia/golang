# Resumen

- bank1. Demuestra como serializar el acceso a una variable centralizandolo en una gorutina, `teller()`. Tambien es unteresante el uso de `select`. Estamos usando `select` tanto para recibir del canal, como para enviar. Cuando alguien solicita 
- bank2. Simula con un canal un mutex
- bank3. Usa un mutex para sincronizar el acceso
- memotest. Se trata de un paquete en el que implementamos utilidades para probar la ejecución secuencial y paralela de un tipo M, que implementa la interface de Memo. 
- memo1. Implementa un paquete con un tipo, `Memo`, que implementa una interface `Get` que permite llamar a una funcion implementando una cache
- memo2. Version mejorada. En este caso protegemos el acceso a la cache con un Mutex
- memo3. Version mejorada. En lugar de extender el mutex durante la lecura y la escritura de la cache, lo aplicamos solo a la lectura, y solo cuando haya que añadir datos a la cache, se serializara el acceso a la misma. Usamos un canal para serializar el acceso a la escritrura.  

    - Siempre se crea un mutex para leer la cache, y solo se libera el mutex con una entrada en la cache
    - En el caso de que no hubiera datos en la cache, se crea una entrada, pero una entrada que tiene apenas el canal y nada más. Luego se libera el mutex y se llama a la función
    - Podría darse el caso de tener dos llamadas en las que hayamos encontrado una entrada en la cache, pero sin datos, solo con el canal. En este caso todas menos una llamada quedarian bloqueadas a la espera de que `<-e.ready`
    - Tras la llamada a la función, hacemos `close(e.ready)`, con lo que todas aquellas llamadas que estuvueran esperando a `<-e.ready` continuarían, ya con la respuesta de la función en la cache
    - Total, que la llamada a la función se hará una sola vez 
