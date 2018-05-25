package main
import(
	"net/http"
	"encoding/json" //para poder servir archivos json
)

type curso struct{
	Titulo string //al mostrar en json la primer letra tiene que ser mayuscula en caso de json
	Duracion int //al igual con duracion tiene que se DuracionS
}

func main(){

	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		cur := curso{Titulo:"primertutorail",Duracion:30}
		json.NewEncoder(w).Encode(cur) //le pasamos un string de datos en este w, por asi decirlo manda la respuesta la navegador
	})

	http.ListenAndServe("localhost:3000",nil)

}