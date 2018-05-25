/*para poder ocupar y poder establecer una conexion a msql neceesitamos el driver
go get -u github.com/go-sql-driver/mysql esto lo pegamos en nuestra terminal

el metodo Ping es para saber si nuestra conexion aun sigue vivas
*/
package main
import(
	"log"
	"fmt"

	"database/sql" //importando la libreria para que funcione sql
 _ "github.com/go-sql-driver/mysql" //una vez instalada la libreria debemos importar
 
)

type User struct{
	Id int `json:"id"`
	Username string `json: username`
	Password string `json:password`
}



type Users []User





//crando constantes para poder realizar la conexion en una base de datoa mysql
const username string ="root"
const password string =""
const host string="localhost"
const port int =3306
const database string ="project_go_Web"


//variable global
var db *sql.DB

func CreateConection(){
	//recibe dos parametros gestor de base de datos a utliizar y el segundo la url
	//esta funcion retorna dos valores, un apuntador y un error en dado caso que haya error
	if	connection,err:=sql.Open("mysql",generateUrl()) ; err!=nil{//aqui estamos abriendo la conexion
		panic(err)	
	}else{
				db=connection;//pasando la conexion a nuestra variable global
				fmt.Println("Conexion Exitosa")
		}
	}

	func CloseConnection(){
		db.Close()
	}

	//funcion ping

	func Ping(){
		if err:=db.Ping(); err !=nil{
			log.Println(err)
		}
	}

//<username>:<password>@tcp(<host>:<port>)/<> database
func generateUrl() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",username,password,host,port,database)//esta funcion reotrnara la url
}

//esta funcion nos permitira saber si la tabla que estamos verificando existe en nuestra base de datos
func existeTabla(tablename string)bool{
		//show tables like 'users';
		sql := fmt.Sprintf("show tables like '%s'",tablename);

		//metodo consulta de nuestro db
		//row,err:=db.Query(sql)

		row,_:=Query(sql)

		// if err!=nil{
		// 	panic(err)
		// }

		//row devuelve un bool en caso de que se haya tornado verdadera nuestra consulta y false en caso que no
		return row.Next()
}

//esta funcion es para crear una nueva tabla dentro de go
func CreateTable(nameTable, schema string){

	//llamamos al metodo existe tabla para saber si hay alguna tabla con el mismo nombre si 
	if !existeTabla(nameTable){
			//procedemos a crearla
			// _,err:=db.Exec(schema) //este parametro Exec es el que ejecuta la sentencia obtenida

			// if err!=nil{
			// 	log.Println(err)
			// }
			Exec(schema)
	}

}


//encapsulando los metodos en funciones propias para agregarlas mas funcionalidad
//los metodos son CreateTable, existeTabla

          //query es la consulta   interface: puede recivir n cantidad de parametros de cualquier tipo
func Exec(query string, args ...interface{}) (sql.Result, error){
		result,err:=db.Exec(query, args...) //este parametro Exec es el que ejecuta la sentencia obtenida
		if err !=nil{
			log.Println(err)
		}

	return result,err
}



	func Query(query string, args ...interface{}) (*sql.Rows, error){
			//metodo consulta de nuestro db
			row,err:=db.Query(query, args...)

			if err!=nil{
				panic(err)
			}

		return row,err
	}


func main(){

	const schema string = `create table users(
		id int(6) UNSIGNED AUTO_INCREMENT primary key,
		username varchar(50) not null,
		Pass varchar(50) not null,
		email varchar(50),
		create_date timestamp default CURRENT_TIMESTAMP
	)`
	
	CreateConection()

	resultado:=existeTabla("users")
	fmt.Println(resultado)
	//Ping()



	//creando la tabla
		//CreateTable("users",schema)

	user:=NewUser("Junior","123123")
	user.save()
	fmt.Println(user)
	
	CloseConnection()
	
}



//creando constructor 
func NewUser(usernombre,passwordl string) *User{
	usuario:=&User{Username:usernombre,Password:passwordl}
	return usuario
}




func (this *User) save(){
	//estos campos username y pass son propios de la tabla en mysql
	sql:="Insert Users SET username=?, pass=?"

	//el metodo exec se encarga de que los parametros sean correctos, es por eso que le ponemos ?
	Exec(sql,this.Username, this.Password)
}

