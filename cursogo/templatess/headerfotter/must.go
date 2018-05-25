/*Trabajar Plantillas de forma sergura con most
la funcion most nos permite genernar template fuera de la funcion y fuera de los metodos*/

package main
import(
	"html/template"
	"net/http"
	
)


//observar que aqui estea fuera porque tiene la funcion most
var templates=template.Must(template.New("T").ParseFiles("./index.html","./footer.html","./head.html","./registro.html"))
//fijarse en la T


func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		//templates es la variable de arriba var templates
		err:=templates.ExecuteTemplate(w,"registro", nil) //resgistro corresponde al define en el archivo registro

		if err!=nil{
			panic(err)
		}
	})

	http.ListenAndServe("localhost:3000",nil)
}