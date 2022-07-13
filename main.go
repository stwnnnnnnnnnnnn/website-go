package main

import (
	"fmt"
	"net/http"
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

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	// http.HandleFunc("/divide", Divide)

	fmt.Printf("Starting application on port %s.", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
