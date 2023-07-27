package renderer

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func RenderTemplateTest(w http.ResponseWriter, filename string) {
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

// map declaration to cache content and template
var tempCache = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, filename string) {
	var tmpl *template.Template
	var err error
	// return the key and value? if the map does not exist then the value is false?
	_, inMap := tempCache[filename]
	//check if cache is already available
	if !inMap {
		log.Println("creating a new template")
		err = createRenderedTemplate(filename)
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Println("using template cache...")
	}
	tmpl = tempCache[filename]
	tmpl.Execute(w, nil)
}

func createRenderedTemplate(filename string) error {
	var templates = []string{fmt.Sprintf("./templates/%s", filename), "./templates/base.layout.tmpl"}
	// parse template and content
	tmpl, err := template.ParseFiles(templates...)
	// if there is an error, return the error
	if err != nil {
		return err
	}
	// assign the parsed template and return nil as an error
	tempCache[filename] = tmpl
	return nil
}
