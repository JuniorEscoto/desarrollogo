package main

import(
	"fmt"
)
 func main(){
	//mat:=[3]int{1,2,3}, esto vendria siendo igual que esto:

	slice:=make([]int,3) //le esatamos pasando el primer parametro que es el arreglo vacio,como segundo la longitud

	segundoSlice:=make([]int,3,5)//observar el parametro 5 que le estamos pasando

	//agregando un nuevo elemento a nuestro slide con append
	segundoSlice=append(segundoSlice,2)
	
	fmt.Println(slice)

	//cap nos devuelve la capacidad de nuestro slice que en este caso es tres
	fmt.Println(cap(slice))

	//observar que la capacidad es 5 y no es tres
	fmt.Println(cap(segundoSlice))


	fmt.Println(segundoSlice) //fijarse que nuestro slice era de 3 pero no afecto en nada cuando agregamos el 2
 }