//Responder archivos html

package main
import(
	"log"
	"net/http"
	"html/template"
)
func main(){

	http.HandleFunc("/",func (w http.ResponseWriter, r *http.Request){

		//PaseFiles nos trae un archivo en este caso el index.html
		template,err:=template.ParseFiles("./index.html")

		if err !=nil{
			panic(err)
		}

		//ejecutar
		template.Execute(w,nil)
	})

	log.Fatal(http.ListenAndServe("localhost:3000",nil))

}