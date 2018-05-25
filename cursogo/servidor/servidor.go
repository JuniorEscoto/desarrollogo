
package main
import(
	"fmt"
	"net/http"
	"log"
)

	func hola(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"Hola Mundo")
	}

	func holados(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"Este es el mundo dos")
	}


func main(){

	//creando nuestro propio server mux
	mux:=http.NewServeMux()
	mux.HandleFunc("/", holados) //podemos hacer mas de una direccion
	mux.HandleFunc("/direcciondos",hola)


	http.HandleFunc("/",hola) //default servermux
		//creando nuestro propio servidor
		sever:=&http.Server{

			Addr:"localhost:3000", //añadiendo direccion
		
			Handler:mux, //si el parametro es nil quiere decir que es un defaultservemux, y nos devolvera la peticion http

			//cambai nil por el mux que creamos y veras que ahora redirecciona a nuestra funcion 
		}

		log.Fatal(sever.ListenAndServe())
		
	
	//http.ListenAndServe("localhost:3000",nil)
}
