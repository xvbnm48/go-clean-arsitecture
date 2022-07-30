package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xvbnm48/go-clean-arsitecture/routes"
)

func main() {
	// fmt.Println("hello world")

	router := mux.NewRouter()
	const port = ":8080"
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	router.HandleFunc("/posts", routes.GetPost).Methods("GET")
	router.HandleFunc("/posts", routes.AddPost).Methods("POST")
	log.Println("Server started on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
