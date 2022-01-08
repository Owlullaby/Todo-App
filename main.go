package main

import (
	"fmt"
	"todo_app/router"
	"log"
	"net/http"
)

func main() {
	response := router.Router()
	fmt.Println("Starting server on the port 4000...")

	log.Fatal(http.ListenAndServe(":4000", response))
}