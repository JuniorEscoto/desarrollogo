package main  //gorutine aplicaciones concurrentes, osea que sean algo asi como asincrono
import(
	"fmt"
	"time" //libreria time
	"strings" //esta libreria podejoms realizar varios metodos fijarse que termina en s al final
)

//funcion para separar palabreas
func separador(nombre string){
	letras:=strings.Split(nombre,"") //recibe dos parametros, el nombre, y la separador ose - o un " " en este caso



	//cliclo for el guion bajo no sirve para nada, let es igual a cada una de las letras y despues imprimimos
	for _, let:=range(letras){
		time.Sleep(1000*time.Millisecond) //esta linea es que por cada vez que entre al ciclo me recoja una letra
		fmt.Println(let)
	}

}

func main(){

	//implementacion de goroutine es anteponer go
	//observar bien que esto pasara esa funcion a un segundo plano por asi decirlo
	go separador("junior")

	fmt.Println("que aburrido") //para que se ejecute esta linea tenemos que esperar que la separador finalize

	var wait string
	//Scanln nos permite dar un enter en la terminal
	fmt.Scanln(&wait) //el programa finalizara hasta que yo de enter, fijarse en poner &

}