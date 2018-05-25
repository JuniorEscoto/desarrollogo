package main

import(
	"fmt"
)
func main(){

	//seria un slice con dato null
	var matriz []int //fijarse que no se le pone = 

	//este slice ya esta intanciado por lo que le estamos dando valores
	mat:=[3]int{1,2,3}
	slice:=mat[1:2] //esta forma se parte al arreglo desde la posicion indicada en este caso es uno


	

	if(matriz==nil){
		fmt.Println("esta vacio")
	}else{
		fmt.Println(matriz)
	}

	//fmt.Println(slice)

	//divicion del arrelgo de slice igual a arreglodos
	fmt.Println(slice)
	
	//puntero al arreglo
	//Longitud del arreglo que punta
	//capacidad
	

}