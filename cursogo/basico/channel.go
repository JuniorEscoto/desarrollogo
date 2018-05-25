/*los go channel nos permite comunicar go rutine unos con otros
las operaciones basicas de los chanel es mandar y recibir
c <- "nombre_datos"   *aqui estamo enviando informacion a un chanel con el operador <-
<- c //aqui estamos recibiendo informacion del channel, el operador esta hacia la izquierda
*/

package main
import(
	"fmt"
	 "time"
)



//funcion

func enviarping(c chan string){
	//ciclo infinito
	for{
		c <-"ping" //enviando informacion a nuestro chanel
	}
}


//imprimir

func imprimirping(c chan string){
	var contador int 
	
	for{
		contador++
		
		fmt.Println(<- c," ",contador)
		time.Sleep(time.Second * 1)
		
	}
}

func main(){

	//para crear un canal es con la funcion make
	chanel:=make(chan string) //chan para saber que es canal y el tipo string



	//fijarse que lo hacemos con gorutine
	go enviarping(chanel)
	go imprimirping(chanel)

	var wait string
	fmt.Scanln(&wait)

	fmt.Println("fin..")

}