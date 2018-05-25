/*leer archivos en go*/
package main
import(
	"io/ioutil"//nos permite leer los archivos
	"fmt"
)

func main(){

	//leyendo el archivo hola.txt con io/ioutil, tener cuidado con mayusculas y minusculas
	texto,err := ioutil.ReadFile("./hola.txt")

	if err != nil{
		fmt.Println("hubo un error")
	}

	fmt.Println(string(texto))

}