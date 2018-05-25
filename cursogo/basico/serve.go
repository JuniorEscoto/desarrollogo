package main

import("net/http")

func main(){
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){

		//aqui estamos sirviendo un html desde go
		http.ServeFile(w,r,"index.html")

	})

	http.ListenAndServe("localhost:3000", nil)
}