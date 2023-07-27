package renderer

import (
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, filename string) {
	parsedTemplate, err := template.ParseFiles("./templates/"+filename, "./templates/base.layout.tmpl")
	if err != nil {
		http.Error(w, "Error while parsing template", http.StatusInternalServerError)
		return
	}

	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error while rendering template", http.StatusInternalServerError)
		return
	}
}
