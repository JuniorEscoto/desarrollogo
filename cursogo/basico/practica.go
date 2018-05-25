package main

import(
	"fmt"
	"time"	
)

func enviar(c chan string){
	for{
		c <- "enviando datos"
	}
}

func imprimir(c chan string){
	var i int
	for{
		i++
		fmt.Println(<-c," ",i)
		time.Sleep(time.Second*1)
	}
}







func main(){

	canal:=make(chan string)

	go enviar(canal)
	go imprimir(canal)

	var wait string
	fmt.Scanln(&wait)

	fmt.Println("fin..")

}