package main
import(
	"net/http"
	"log"
)

func main(){

	redirect:=http.RedirectHandler("http://www.youtube.com",200)
	err:=http.NotFoundHandler()


	mux:=http.NewServeMux()
	mux.Handle("/redireccion",redirect)
	mux.Handle("/eror",err)


	server:=&http.Server{
		Addr: "localhost:3000",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}