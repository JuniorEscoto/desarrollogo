package main
import(
	"./model"
	"fmt"
)

func main(){
	CreateConection()

	resultado:=existeTabla("users")
	fmt.Println(resultado)
	//Ping()



	//creando la tabla
	CreateTable("users",schema)

	user:=model.NewUser("Junior","123123")
	
	fmt.Println(user)

	user.save()
	CloseConnection()
}