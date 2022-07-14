package main

import (
	"fmt"
	"net/http"

	"github.com/stwnnnnnnnnnnnn/website-go/pkg/config"
	"github.com/stwnnnnnnnnnnnn/website-go/pkg/handlers"
	"github.com/stwnnnnnnnnnnnn/website-go/pkg/render"
)

const portNumber = ":8000"

func main() {
	//route handler
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	// w.Write([]byte("hello world"))
	// 	n, err := fmt.Fprintf(w, "hello wolrd")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(fmt.Sprintf("number of bytes written : %d", n))
	// })

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		fmt.Println(err)
	}

	app.TemplateCache = tc
	app.Usecache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	// http.HandleFunc("/divide", Divide)

	fmt.Printf("Starting application on port %s.", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
