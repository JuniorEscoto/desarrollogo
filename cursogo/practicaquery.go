package main

import(
	
	"fmt"
	"net/http"
	"log"
)

func main(){

	http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){

		fmt.Println(r.URL)

		//añadiendo valores a nuestra variable con query para separalas
		values := r.URL.Query()


		//este metodo es para eliminar parametros de nuestra consulta
		values.Del("otro") //este tiene que coidicir con el parametro que hemos puesto en la url

		
		values.Add("nombre","junior") //añadiendo parametros a nuestra query

		//anexando los parametros a nuestra url
		r.URL.RawQuery=values.Encode()

		fmt.Println(r.URL)

	})

	log.Fatal(http.ListenAndServe("localhost:3002",nil))
}