package main
import(
	"fmt"
	"net/url"
	"net/http"
	"io/ioutil"
)

func CreateUrl() string{
	u,err := url.Parse("/params") 
	if err != nil{
		panic(err)
	}
 
	
	u.Host="localhost:3000" 
 
	
	u.Scheme="http"
 
	
 
	  query := u.Query() 
	  query.Add("nombre","junior") 
 
	
	  u.RawQuery=query.Encode()
 
	  return u.String()
}

func main(){

	url:=CreateUrl()

	//generando peticion al servidor  RequesteSet.go
	request,err:=http.NewRequest("GET",url, nil) //recibe tres parametros, sobre que metodo(get,post,put,delete) segundo: url , nil

	if err !=nil{
		panic(err)
	}


	//una vex teniendo nuestro request, generamos un cliente
	cliet :=&http.Client{}
	reponse,err:=cliet.Do(request)  //se encarga de ejecuatr la peticion, como parametro es el request, retrona un response y un error
	if err !=nil{
		panic(err)
	}

	//visualizando que nos trae el response
	fmt.Println("El header es : ",reponse.Header)
	body,err:=ioutil.ReadAll(reponse.Body)

	if err !=nil{
		panic(err)
	}
	fmt.Println("El body es : ",string(body))
	fmt.Println("El Estatus es : ",reponse.Status)
	//fmt.Println("la url es:",url)

}