package main

import(
	"fmt"
	"net/url" //importar la libreria url
)

//creacion de funcion, que retornara un tipo string
func CreateURl() string{
 //para generar la url nesecitamos el metodo parse y solo ponemos la uri, en lineas posteriares creamos lo demas
   u,err := url.Parse("/params") //nos devolvera la url exitosa u, o con un error err
   if err != nil{
	   panic(err)//panic es para terminar el programa
   }

   //agregandole atrributos a nuestra url, la url tiene un atributo llamado host
   u.Host="localhost:3002" //colocalndo nuestro host

   //colocando que protocolo vamos a usar
   u.Scheme="http"

   // agregandole parametros a nuestra url

	 query := u.Query() //query nos regresa un mapa
	 query.Add("nombre","valor") //ya podemos utilizar el metodo Add

	 //indexando nuestros atributos a nuestra url 
	 u.RawQuery=query.Encode()

	 return u.String() //retronando el string tiene que estar en mayusculas
}
func main(){

	url:=CreateURl()
	
	fmt.Println("la url final es: "+url)

	

}