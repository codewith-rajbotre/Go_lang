package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Mod Package in GoLang")
	//greeter()

	//Routing System in Go lang gorilla/mux
	r := mux.NewRouter()
	r.HandleFunc("/", servHome).Methods("GET")

	//Running a server in Go lang
	log.Fatal(http.ListenAndServe(":4000", r))
}

// func greeter() {
// 	fmt.Println("Hey Mod Users")
// }

// w --http response writer
// r-- request (this is not and ordinary variable this is a pointer)
func servHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Go Lang</h1>"))
}
