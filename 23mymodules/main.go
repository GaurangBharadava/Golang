package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("hello mod in golang")

	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/",servHome).Methods("GET")

	//listening server
	log.Fatal(http.ListenAndServe(":4000",r)) // if any error occures then lof.Fatel will log thr error
}

func greeter() {
	fmt.Println("hey there mode users")
}

func servHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to golang lerning</h1>"))
}