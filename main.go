package main

import (
	"fmt"
	"log"
	"net/http"
	. "web/routes"
)

func main() {
	Routes() //Calling routes from web/routes
	fmt.Println("Started...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
