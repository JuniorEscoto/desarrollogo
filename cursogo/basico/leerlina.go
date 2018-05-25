package main
import(
	"fmt"
	"bufio"
	"os" 
)

func main(){
	text,err := os.Open("./holas.txt")

	if err != nil{
		//fmt.Println("hubo un erro")
		panic(err) //panic es la impresion de un error a detalle

		//recove, se salta el panic osea muestra las otras ejecuciones sin pasar por panic
	}

	escaner:=bufio.NewScanner(text)

	var i int
	for escaner.Scan(){
		i++
		linea:=escaner.Text()
		fmt.Println(i)
		fmt.Println(linea)
	}

}