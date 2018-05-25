/*cuando nuestro servidor requeira hacer multiples peticiones a paginas web, tenemos que cambiar
el metodo ParseFiles por parseGlob el cual tendra un unico parametro*/

package main
import(
	"html/template"
	"net/http"
	
)


//observar que aqui estea fuera porque tiene la funcion most
var templates=template.Must(template.New("T").ParseGlob("./*.html"))
//obten todos los archivo con extension html de la carpeta en caso de estar en una sola carpeta los archivos html



/*En caso de tener los archivo separados por folder, digamos que tenemos una carpeta base que se llama template
en ella hay varios folder en donde hay html en ese caso para poder acceder a ellos se realizaria de la siguiente
manera agregando dos asteriscos
--------var templates=template.Must(template.New("T").ParseGlob(".template/**./*.html"))-----
aqui puse un punto porque daba problemas pero el punto no va!!
*/


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