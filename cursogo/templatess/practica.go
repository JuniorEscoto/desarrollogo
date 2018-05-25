package main

import(
	"html/template"
	"net/http"
	"fmt"
)


func simple()string{
	return "hola desde la funcion"
}



func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){


			fmt.Println("Peticion")
			//template,err:=template.New("junior").Parse("Esta es una nueva plantilla")

			
			funciones:=template.FuncMap{
				"simple":simple,
			}

			template,err:=template.New("practica.html").Funcs(funciones).ParseFiles("./practica.html")

			if err!=nil{
				panic(err)
			}



			template.Execute(w,nil)
	})

	http.ListenAndServe("localhost:3000",nil)
}