package main
import("fmt")

func main(){
	//los mapas nos permiten indicar un indice y un valor
	x:=make(map[string] string) //el primer string es el tipo del indice, y es segundo el tipo del valor

	fmt.Println(x)

	//segunda forma de declara un map
	y:=make(map[string] string, 2)  //el 2 representa la capaciadad del mapas
	fmt.Println(y)

	//agreganodles valores

	x["nombre"]="junior"
	x["edad"]="20"
	fmt.Println(x)

	//accesdiendo a las valores de map
	fmt.Println(x["nombre"])


	//tercer forma de declarar map con valores ya definidos

	//fijarse que este no lleva make
	edades:= map[string] int{ //el indice es de tipo string pero los valores seran enteros
		"ana" : 34,
		"pedro": 23,
		"maria": 18,
	}

	fmt.Println(edades)

	//otro ejemplo con indece int y valor string
	dias:=map[int] string{
		1: "lunes",
		2:"martes",
		3: "miercoles",
	}

	fmt.Println(dias[1])

	//forma de elminar elementos en un mapa con delete
	delete(edades,"ana") //del mapa edades volar el indice ana
	fmt.Println(edades)


	//al poner un elemento que no exista nos devolbera un valor 0 en caso int, o vacion en string
	//en el caso de int podemos hacer operaciones ya que nos devuelve 0
	edades["lucas"]++
	fmt.Println(edades)

	//los mapaas devuel un valor si existe y otro valor booleano
	if edad,ok := edades["maria"]; ok{

		if edad>18{
			fmt.Println("mayor de edad")
		}else{
			fmt.Println("menor de edad")
		}

	}else{
		fmt.Println("no existe registro")
	}


	//viendo todos los elementos del  map con range en edades
	for nombre,edad :=range(edades){
		fmt.Printf("el nombre es : %s y la edad es: %d \n ",nombre,edad)
	}
}