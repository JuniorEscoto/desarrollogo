package main
import(
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const username string="root"
const password string=""
const host string="localhost"
const port int =3306
const database string ="pruebadatabasego"

var db *sql.DB

func CreateConexion(){
	exito,err:=sql.Open("mysql",generateUrl())
	if err !=nil{
		panic(err)
	}else{
		db=exito;
		fmt.Println("conexion Exitosa")
	}
}
//<username>:<password>@tcp(<host>:<port>)/<> database
func generateUrl() string{
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",username,password,host,port,database)
}

func closeconnection(){
	db.Close()
}


/**************************************************************************************************************/
/**************************************************************************************************************/

func CreateTable(nameTable,schema string){
	if !existeTabla(nameTable){
		Execute(schema)
		fmt.Println("Tabla Creada Exitosamente!!")
	}else{
		fmt.Println("La tabla ya existe!")
	}
		
}

func Execute(query string, args ...interface{}) (sql.Result,error){
	result,err:=db.Exec(query,args...)
	if err !=nil{
		panic(err)
	}

	return result,err
}

func existeTabla(tablename string) bool{
	sql:=fmt.Sprintf("show tables like '%s'",tablename)
	row,_:=Query(sql)
	return row.Next()

}

func Query(query string, args ...interface{}) (*sql.Rows, error){
	row,err:=db.Query(query, args...)
	if err!=nil{
		panic(err)
	}
	return row,err
}

/**************************************************************************************************************/
/**************************************************************************************************************/


type Cliente struct{
	Nombre string
	Password string
}

type Arreglo [] Cliente

func NewCliente(nom,pass string) *Cliente{
	usuario:=&Cliente{Nombre:nom,Password:pass}
	return usuario
}

func(this *Cliente) Guardar(){
	sql:="Insert users SET username=?, pass=?"
	Execute(sql,this.Nombre,this.Password)
	fmt.Println("Datos Guardados")
}

/**************************************************************************************************************/
/**************************************************************************************************************/

func main(){

	//esquema de la tabla
	const schema string = `create table users(
		id int(6) UNSIGNED AUTO_INCREMENT primary key,
		username varchar(50) not null,
		Pass varchar(50) not null,
		email varchar(50),
		create_date timestamp default CURRENT_TIMESTAMP
	)`

	CreateConexion()
	CreateTable("users",schema)
	usuario:=NewCliente("junior","mipassword")
	fmt.Println(usuario)
	usuario.Guardar()
	closeconnection()

}