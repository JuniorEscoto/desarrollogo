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
	Id int
	Nombre string
	Password string
	Email string
}

func NewCliente(nom,pass,email string) *Cliente{
	usuario:=&Cliente{Nombre:nom,Password:pass,Email:email}
	return usuario
}

func(this *Cliente) Guardar(){
	sql:="Insert users SET username=?, pass=?, email=?"
	Execute(sql,this.Nombre,this.Password,this.Email)
	fmt.Println("Datos Guardados")
}

func CreateCliente(nom,pass,email string) *Cliente{
	cliente:=NewCliente(nom,pass,email)
	cliente.Guardar()
	return cliente
}

func (this *Cliente) EliminarCliente(){
	sql:="Delete  from users where id=?"
	Execute(sql,this.Id)
	fmt.Println("Datos Eliminados")
}


//obtener usuario
func seleccionarUsuario(id int) *Cliente{
	usuario:=NewCliente("","","") //pasandole tres valores vacios

	//realizando consulta
	sql:="SELECT id, username, pass, email from users where id=?"

	//se utiliza el metodo query ya que vamos a obtener registros
	rows,_:=Query(sql,id)

	//obteniendo los valores
	for rows.Next(){
		//todos los valores que estamos pasando deben de ser punteros, es por eso el : &
		rows.Scan(&usuario.Id,&usuario.Nombre,&usuario.Password,&usuario.Email) //obsrvar que los valores son los mismos que el struct
	}

	return usuario

}

//aqui estamos seleccionando todos los usuarios que tengamos en nuestra tabla 
   type listaCliente [] Cliente //para eso necesitamos un arreglos de tipo cliente
func totalUsuarios() listaCliente {
	sql:="SELECT id, username, pass, email from users "
	 //creando un objeto para retornar
	 usuarios:=listaCliente{}

	 rows,_:=Query(sql)

	 for rows.Next(){
		 client:=Cliente{}
		 rows.Scan(&client.Id,&client.Nombre,&client.Password,&client.Email) 

		 usuarios=append(usuarios,client) //agregando a nuestra lista
	 }

	 return usuarios;
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
	usuario:=CreateCliente("pjuanedro","123","maria@gmail.com")
	fmt.Println(usuario)

	//dando valor al id y despues mandarlo al metodo EliminarCliente()
	usuario.Id=4
	usuario.EliminarCliente()

	getCliente:=seleccionarUsuario(1) //aqui solo estamos obteniendo un cliente por que solo pasamos el id =1
	fmt.Println(getCliente)



	//mostrando todos los registrso de nuestra tabla
	totalclientes:=totalUsuarios()
	fmt.Println(totalclientes)

	closeconnection()

}