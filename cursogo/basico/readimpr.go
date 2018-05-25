//Leer e imprimir datos

package main
import(
	"fmt"
)


func main(){

	//fmt.Print("hola") este imprimira todo pero en una sola linea

	edad:=22
	nombre:="junior"
	bandera:=true

	precio:=12.34
	//fijarse en Printf
	fmt.Printf("mi edad es %d\n",edad) // %d es como concatenar valores en este caso edad en c# sera {0}

	fmt.Printf("mi nombres es %v\n",nombre) //cuando es es string cambia el verbo a %v tambien se puede con %s

	fmt.Printf("bandera : %t\n",bandera) //aqui es booleano %t

	fmt.Printf("precio es : %f\n",precio) //aqui es punto flotante %f



	//permite capturar datos desde la terminal, como en la consola de C#
	var ed int
	fmt.Println("coloca tu edad:")
	fmt.Scanf("%d\n",&ed) //con scanf podemos capturar el valor, %d por que es entero el dato a capturar

	fmt.Printf("la edad que puesistes es: %d",ed)


}