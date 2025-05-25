package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
	}
	fmt.Fprint(w, "Hello")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm () err %v", err)
	}
	fmt.Fprintf(w, "POST Request Successful")
	name := r.FormValue("name")
	password := r.FormValue("password")
	fmt.Fprintf(w, "Name : \n", name)
	fmt.Fprintf(w, "Password : \n", password)

}

func main() {
	fileServer := http.FileServer(http.Dir("./1_BuildWebServer"))

	http.HandleFunc("/", fileServer.ServeHTTP)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Server started at pirt 8080 ")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
