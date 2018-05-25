package main

import(
    "fmt"
    "net/http"
    "log"
)


func main() {
    http.HandleFunc("/paramet",func(w http.ResponseWriter, r *http.Request){

        //metodo para separar una consulta
       
        fmt.Println(r.URL.Query())

        //crando una variable y guardando la consulta echa en la url
        name := r.URL.Query().Get("name")
        valor := r.URL.Query().Get("valor")

        if len(name)!=0{
            fmt.Println("hola: "+name +" valor: "+valor)
        }else{
            fmt.Fprintf(w,"hola mundo")
        }
        
    })
        
    
    log.Fatal(http.ListenAndServe("localhost:3002",nil))
}
