package main

import(
	"fmt"
)

//creando una estructura
type User struct{
	 edad int
	 nombre,apellido string
}

func main(){

	var usuario User

	//forma de pasarle datos a nuestro struct
	pedro:=User{edad:40,nombre:"pedro",apellido:"Lopez"}

	juan := new (User)
	juan.nombre="mi nombre es juan"
	juan.apellido="hernandez"

	fmt.Println(usuario)

	//otra manera de pasar datos a nuestro struct
	fmt.Println(User{nombre:"junior",apellido:"Escoto"})

	fmt.Println(pedro)

	fmt.Println(juan.nombre+" "+juan.apellido)
}