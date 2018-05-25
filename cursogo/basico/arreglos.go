package main
import(
"fmt"
)

func main(){

	 //forma de declarar un arrglo
	 var numeros [10] int

	 //arreglo multidimencionales
	 var matriz[3][2] int

	 matriz[2][1]=50

	 //esta forma ya declara valores en nuestro arreglo
	 arrglos:=[3] int {1,2,3}

	 //para conocer el tamaño de un arreglo basta con poner len(arreglo)

	 fmt.Println(numeros)
	 fmt.Println(arrglos)
	 fmt.Println(matriz)
}