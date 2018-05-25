package main

import(
	"fmt"
	"strconv" //esta libreria se utiliza para convertir de int a string o viceversa
)

func main(){

	edad:=23 //edad de tipo entero
	edad_string:=strconv.Itoa(edad) //Itoa es para convertir de int a string y lo guardaamos en una nueva variable

	peso:="45" //este es un string
	peso_int, _:=strconv.Atoi(peso) //pasando de string a int, pero este retornara dos valores
	//en este caso le pasamos un guin bajo _ con este capturamos el valor y lo desechamos

	fmt.Println("mi edad es: "+edad_string)

	fmt.Println(peso_int+20)

}