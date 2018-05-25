package main

import(
	"log"
	"net/http"
	"html/template"  //libreria de html
)

func main(){
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		
	template,err:=template.New("hola").Parse("hola") //genera un nuevo template y recibe como parametro un string
	//el metodo parse nos devuelve un template y un erro

	if err!=nil{
		panic(err)
	}

	//para que el mensaje pueda ser visualizado por un cliente realizamos lo siguiente

	//1 donde se desea mostrar el template, 2 strucutra
	template.Execute(w,nil)
	})

	log.Fatal(http.ListenAndServe("localhost:3000",nil))

}