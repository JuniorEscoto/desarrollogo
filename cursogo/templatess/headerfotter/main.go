/*En este archivo estamos recogiendo multiples templates*/

package main
import(
	"html/template"
	"net/http"
	"fmt"
)


func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){


			fmt.Println("Petecion")
			//template,err:=template.New("junior").Parse("Esta es una nueva plantilla")

			
			funciones:=template.FuncMap{
				"simple":simple,
			}

			template,err:=template.New("index.html").ParseFiles("./index.html","./footer.html","./head.html")

			if err!=nil{
				panic(err)
			}
			template.Execute(w,nil)
	})

	http.ListenAndServe("localhost:3000",nil)
}