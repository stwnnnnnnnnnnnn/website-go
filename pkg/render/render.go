package render

import (
	"fmt"
	"html/template"
	"net/http"
)

//render template using html
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	//response writer tidak perlu pakai return
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}
}
