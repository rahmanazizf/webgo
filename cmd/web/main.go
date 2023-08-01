package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rahmanazizf/basicwgo/pkg/config"
	"github.com/rahmanazizf/basicwgo/pkg/handlers"
	"github.com/rahmanazizf/basicwgo/pkg/renderer"
)

// use const instead of var if you want to keep a variable constant
const portNumber = ":8081"

func main() {
	var app config.AppConfig

	tc, err := renderer.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting the application on localhost%s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
