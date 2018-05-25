package main
import("fmt")

type User struct{
	name string
}

type Custom struct{
	User //aqui basicamente estamos heredando
	dos  //aquie estamos haciendo otro heredacion a dos
}

type dos struct{
	name string
}

//funcion hablar de user
func (this User) hablar() string{
	return "bla bla bla"
}

func main(){

	nombre:=Custom{User{"junior"},dos{"pedro"}}

	//es este caso es necesario declara a quien nos estamos referiendo
	//fmt.Println(nombre.name) //esto daria error

	fmt.Println(nombre.dos.name) //aquie estamos haciendo a la referencia dos.name

	
	fmt.Println(nombre.hablar())

}