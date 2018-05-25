//Utilizacion de metodos get,put,delete, push


package main

import (
    "net/http"
    "log"
    "github.com/gorilla/mux"
    "fmt"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
    vars:=mux.Vars(r) //Vars es un mapa le pertenece a gorilla mux

    nombre:=vars["nombre"]  //obteniendo los valores del mapa vars
    id:=vars["id"]
    fmt.Fprintf(w,"el nombre:"+nombre+" el id: "+id)

    //w.Write([]byte("Gorilla!\n"))
}

func main() {
    r := mux.NewRouter()
    // Routes consist of a path and a handler function.

    //leyendo parametros de la url
    r.HandleFunc("/usuarios/{nombre}/{id:[0-9]+}", YourHandler).Methods("PUT","DELETE") // al poner id: [0-9]+ es una expresion regular que solo acepte numeros
	/*metodos: get, put, delete, estos metodos siempre estan asociados a una base de datos
	como get, en esta uri que tenemos seria obtener el nombre con el id que mandamos
	*/


    // Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":3000", r))
}