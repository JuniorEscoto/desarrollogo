package main

import("fmt")

func main(){

	/*
	1.es una direccion de memoria
	2.en lugar del valor tenemos la direccionen donde esta el valor
	X,Y =>dee3=>5   x,y estan en direccion 5
	X=>dee3=>6 x cambia su valor a 6 por lo tanto Y tambien tendra valor 6
	Y=>dee3=>6
	*/

	var x,y *int //el asterico quiere dar a entender que es un puntero

	entero:=5 //creamos una variable con el valor 5

	x=&entero //aquie con & estamos guardando la direccion en memoria de la variable entero
	y=&entero

	*x=6  //si imprimimos esto nos dara el resultado 6 en valor y en direccon

	//como ves nos dara el mismo resultado, porque los dos estan apuntanto a la misma direccion
	fmt.Println(x)
	fmt.Println(y)

	fmt.Println(*x)
	fmt.Println(*y)
}