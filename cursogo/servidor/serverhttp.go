//el handle recibe en la url cualquir uri y siempre nos redireccionara de manera exitosa

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
	junior:=&User{name:"Junior"}

	server:=&http.Server{
		Addr: "localhost:3000",
		Handler : junior,

	}

	log.Fatal(server.ListenAndServe())

}