package main

//curl -X POST http://localhost:8080/posts -H 'Content-Type: application/json' -d '{"title":"williambook","text":"nicetexthere"}'
// $env:GOOGLE_APPLICATION_CREDENTIALS="C:\Users\PATHTOKEY"
import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	port string = ":8080"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up")
	})
	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", addPosts).Methods("POST")

	log.Println("Listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))

}
