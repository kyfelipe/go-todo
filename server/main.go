package main

import (
	"fmt"
	"github.com/kyfelipe/go-todo/server/router"
	"log"
	"net/http"
)

// https://levelup.gitconnected.com/build-a-todo-app-in-golang-mongodb-and-react-e1357b4690a6
func main() {
	r := router.Router()
	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
