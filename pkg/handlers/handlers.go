package handlers

import (
	"net/http"

	"github.com/stwnnnnnnnnnnnn/website-go/pkg/config"
	"github.com/stwnnnnnnnnnnnn/website-go/pkg/models"
	"github.com/stwnnnnnnnnnnnn/website-go/pkg/render"
)

//template data holds data from handlers to template
// type TemplateData struct {
// 	StringMap map[string]string
// 	IntMap    map[string]int
// 	FloatMap  map[string]float32
// 	Data      map[string]interface{}
// 	CSRFToken string
// 	Flash     string
// 	Warning   string
// 	Error     string
// }

var Repo *Repository

// repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// newreport create a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// newhandlers sets the repository for the handlers

func NewHandlers(r *Repository) {
	Repo = r
}

// function untuk response dan request harus 2 parameter
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "this is the homepage")
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "hello"
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
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
