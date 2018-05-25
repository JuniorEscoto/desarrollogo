package main
import("fmt")

type User struct{
	edad int
	nombre,apellido string
}

//creando metodo---- this es como el identificador (puede llevar cualquier nombre) ---y el string es porque retorna un string
//(this user): seria el metodo, el nombre de la funcion seria:nombre_metodo()
func (this User) nombre_metodo() string {

	return this.nombre+" "+this.apellido
}

//en esta funcion trataremos de pasarle un dato a nombre y ver el resultado

func (this *User) set_nombre (nuevoNombre string){  //como ves sin el no le ponemos el asterisco no cambia nada
	this.nombre=nuevoNombre
}

func main(){

	usuario:=new(User)
	usuario.nombre="junior"
	usuario.apellido="Escoto"

	usuario.set_nombre("pedro")

	fmt.Println(usuario.nombre_metodo())

	fmt.Println(usuario.nombre)

}