//Responder archivos html

package main
import(
	"log"
	"net/http"
	"html/template"

)

//metodo
func (this Usuario) TienePermisos() string { //recordar que el nombre de la funcion tiene que empezar por mayuscula
	
	return this.User
}

func (this Usuario) Esadministrador(llave string) bool { //recordar que el nombre de la funcion tiene que empezar por mayuscula
	
	return this.Activo && llave =="si"
}




type Curso struct{
	Name string
	Duracion int
}

type Usuario struct{
	User string //recordar que tiene que ser en mayusculas para que funciones
	Edad int
	Activo bool
	Administrador bool

	Lenguaje []string

	Cursos []Curso
}

func main(){

	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){

		//PaseFiles nos trae un archivo en este caso el index.html
		template,err:=template.ParseFiles("./metodo.html")

		if err !=nil{
			panic(err)
		}

		//arreglo con lenguajes de programacion
		tags :=[] string{"java","c#","pyton","php"}


		//objeto curso
		miCurso :=[] Curso{Curso{"go",2}, Curso{"angular",3}}


		usu:=Usuario{User: "Pedro", Edad:24, Activo:true,Administrador:true,Lenguaje:tags, Cursos :miCurso,}

		//ejecutar
		template.Execute(w,usu) //el segundo parametro es una estructura la cual nos ayudara a generar paginas dinamicasS
		//pasando usu como estructra entonces en nuestro html ya los podemos pasar como atributo.. ver index.html
	})
	log.Fatal(http.ListenAndServe("localhost:3000",nil))

}