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




//funcion dentro del template
func funcionsimple()string{
	return "Hola desde la funcion"
}

func suma(valor1,valor2 int) int{
	return valor1+valor2
}

type Usuario struct{
	User string //recordar que tiene que ser en mayusculas para que funciones
	Edad int
	Activo bool
	Administrador bool
}

func main(){

	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){

		//PaseFiles nos trae un archivo en este caso el index.html
		//template,err:=template.ParseFiles("./metodo.html")



		//creacion de un funcmap, que no es mas que un mapa de funciones
		funciones:=template.FuncMap{
			//recuerda llave--- valor
			"funcionsimple": funcionsimple,
			"suma": suma,
		}

		template,err:=template.New("metodo.html").Funcs(funciones).ParseFiles("./metodo.html")//en new el parametro tiene que ser el mismo nombre del template


		if err !=nil{
			panic(err)
		}



		usu:=Usuario{User: "Pedro", Edad:24, Activo:true,Administrador:true,}

	
		//ejecutar
		template.Execute(w,usu) //el segundo parametro es una estructura la cual nos ayudara a generar paginas dinamicasS
		//pasando usu como estructra entonces en nuestro html ya los podemos pasar como atributo.. ver index.html
	})
	log.Fatal(http.ListenAndServe("localhost:3000",nil))

}