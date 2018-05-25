//Este ejemplo es el mismo que serverhttp solo que aqui le creamo sun mux al handle junior

package main

import(
	"fmt"
	"net/http"
	"log"
)

type User struct{
	name string
}

//creando una funcion con metodo
func (this *User) ServeHTTP (w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"mi nombre es: "+this.name)
}

func main(){
	junior:=&User{name:"Mi nombre es Junior"}

	//he aqui la diferencia con el ejercicio anterios
	mux:=http.NewServeMux() 
	mux.Handle("/junior",junior)
	

	server:=&http.Server{
		Addr: "localhost:3000",
		Handler : mux, //cambiamos junior por el mux creado

	}

	log.Fatal(server.ListenAndServe())

}