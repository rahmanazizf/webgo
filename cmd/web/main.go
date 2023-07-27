package main

import (
	"fmt"
	"net/http"

	"github.com/rahmanazizf/basicwgo/cmd/pkg/handlers"
)

// use const instead of var if you want to keep a variable constant
const portNumber = ":8081"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting the application on localhost%s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
