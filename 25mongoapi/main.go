package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GaurangBharadava/mongoapi/router"
)

func main() {
	fmt.Println("coonection of mongoDB in golang")
	fmt.Println("server is getting started")

	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("listening an port 4000...")
}