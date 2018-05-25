package main

import(
"fmt"
"log"
"net/http"
)

//GET, POST, PUT, DELETE

func main(){

	http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){

		switch r.Method{
		case "GET": 
			fmt.Fprintf(w,"hola mundo con el metodo get")
		case "POST":
			fmt.Fprintf(w,"hola con el metodo post")
		case "PUT":
			fmt.Fprintf(w,"hola con el metodo put")
		case "DELETE":
			fmt.Fprintf(w,"hola con el metodo delete")

		default : http.Error(w,"error",http.StatusBadRequest)
		}
		//fmt.Fprintf(w,"hola mundo")
	})
	log.Fatal(http.ListenAndServe("localhost:3002",nil)) //conectando al servidor

}