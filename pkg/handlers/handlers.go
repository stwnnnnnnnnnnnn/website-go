package handlers

import (
	"net/http"

	"github.com/stwnnnnnnnnnnnn/website-go/pkg/render"
)

// function untuk response dan request harus 2 parameter
func Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "this is the homepage")
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
	// sum := addValues(2, 2)
	// _, _ = fmt.Fprintf(w, fmt.Sprintf("this is at the about page and 2 + 2 is %d", sum))
}

// func Divide(w http.ResponseWriter, r *http.Request) {
// 	f, err := divideValues(100.0, 0.0)
// 	if err != nil {
// 		fmt.Fprintf(w, "cannot divide by zero")
// 		return
// 	}

// 	fmt.Fprintf(w, fmt.Sprintf("%f divide by %f is %f", 100.0, 0.0, f))
// }

// func divideValues(x, y float32) (float32, error) {
// 	if y <= 0 {
// 		err := errors.New("cannot divide by zero")
// 		return 0, err
// 	}
// 	result := x / y
// 	return result, nil
// }

// func addValues(x, y int) int {
// 	return x + y
// }
