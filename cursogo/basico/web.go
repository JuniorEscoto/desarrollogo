package main
 import(
	 "net/http"
	 "fmt"
	 "log"
	 "io"
	)

 func main(){

	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		fmt.Println("recibiendo peticion")

		io.WriteString(w,"este es un sitio web") //aqui es para que se muestre el mensaje en el navegador, la respuesta a la peticions

	})
	log.Fatal(http.ListenAndServe("localhost:3000",nil))
 }