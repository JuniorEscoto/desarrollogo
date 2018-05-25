package main

//importaciones
import(

	"net/http" //poder hacer uso del servidor http
	"fmt" 
	"log" //log para poder ver resultados en la terminal
	
) 

func main() {

	//funcion Handlefunc es para hacer peticiones pasamos una uri y una funcion anonima con solicitud y respuesta a nuestro servidor
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		//Println, Printf,Fprintf
		w.Header().Add("nombre","valor del header") //header

		//redirict te perimite redireccionar uri a otra en esta caso al recargar la pagina / automaitcament se redireccionara a google
		http.Redirect(w,r,"http://www.youtube.com",http.StatusMovedPermanently)
		fmt.Fprintf(w,"hola mundo")
	})	
	
	//creando un segundo solicitud y respuesta con la uri /dos
	http.HandleFunc("/dos",func(w http.ResponseWriter, r *http.Request){
		//fmt.Fprintf(w,"segundo mundo")


		//redireccionamiento pagina no encotrada
		http.NotFound(w,r)

		//envio de errores
		http.Error(w,"Este es un error",http.StatusNotFound)
		
	})	


	log.Fatal(http.ListenAndServe("localhost:3001",nil)) //conectando al servidor	
}