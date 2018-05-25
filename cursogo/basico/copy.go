package main

import("fmt")

func main(){
	 primer:= []int{1,2,3,4}

	segundo:=make([]int,2)

	//copy lo que hace es copiar un arreglo del fuente hacia el destino

	copy(segundo,primer)

	fmt.Println(segundo) //imprime solo dos porque la longitud es 2

}