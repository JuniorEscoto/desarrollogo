package main
import(
	"fmt"
	"net/http"
	"log"
)
func main(){
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"Hola Mundo")
	})

	
	http.ListenAndServe("localhost:3000",nil)
}