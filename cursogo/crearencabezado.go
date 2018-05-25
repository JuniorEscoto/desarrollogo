package main

import(
	"fmt"
	"log"
	"net/http"
)

func main(){

	http.HandleFunc("/uno",func(w http.ResponseWriter,r *http.Request){

		//mostrando un encabezado que nos devuelve agent, accept
		fmt.Println(r.Header)

		accesToken :=r.Header.Get("acces_Token") //obteniendo el valor accesToken en la variable 

		if len(accesToken)!=0{
			fmt.Println(accesToken)
		}

		//agregar encabezados
		r.Header.Set("apellido","lopezz")
		fmt.Println(r.Header)


		/*
		para mandar un encabezado en curl hay que poner -H "valor:123" este nos devuelve un mapa
		*/

	})

	log.Fatal(http.ListenAndServe("localhost:3002",nil))
}