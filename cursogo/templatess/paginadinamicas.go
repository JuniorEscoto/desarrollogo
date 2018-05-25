//Responder archivos html

package main
import(
	"log"
	"net/http"
	"html/template"

)

type Usuario struct{
	User string //recordar que tiene que ser en mayusculas para que funciones
	Edad int
}

func main(){

	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){

		//PaseFiles nos trae un archivo en este caso el index.html
		template,err:=template.ParseFiles("./index.html")

		if err !=nil{
			panic(err)
		}


		usu:=Usuario{"pedro",23}

		//ejecutar
		template.Execute(w,usu) //el segundo parametro es una estructura la cual nos ayudara a generar paginas dinamicasS
		//pasando usu como estructra entonces en nuestro html ya los podemos pasar como atributo.. ver index.html
	})
	log.Fatal(http.ListenAndServe("localhost:3000",nil))

}