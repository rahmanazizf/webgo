package renderer

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/rahmanazizf/basicwgo/pkg/config"
	"github.com/rahmanazizf/basicwgo/pkg/models"
)

// declare app to access the same app in the main function
var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate write everything to browser
func RenderTemplate(w http.ResponseWriter, filename string, td *models.TemplateData) {
	var tmpl map[string]*template.Template
	var err error
	if app.UseCache {
		tmpl = app.TemplateCache
	} else {
		// create template cache
		tmpl, err = CreateTemplateCache()
		if err != nil {
			log.Println("error while creating template cache")
			fmt.Println(err)
		}
	}
	// get the requested template
	tc, tcExist := tmpl[filename]
	if !tcExist {
		log.Println("the requested template does not exist")
		return
	}
	// execute using buffer
	buf := new(bytes.Buffer)
	data := AddDefaultData(td)
	// data is passed to template
	err = tc.Execute(buf, data)
	if err != nil {
		log.Println("Error while executing template")
		fmt.Println(err)
		return
	}
	// write to w
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Fatal(err)
		return
	}
}

// CreateTemplateCache to create template cache first
func CreateTemplateCache() (map[string]*template.Template, error) {
	tc := map[string]*template.Template{}

	// get all page tmpl files
	pages, err := filepath.Glob(`C:\Users\USER\Desktop\go\webgo\templates\*.page.tmpl`)
	if err != nil {
		log.Println("Error while finding pages")
		fmt.Println(err)
		return nil, err
	}

	// loop through pages
	for _, page := range pages {
		name := filepath.Base(page)
		// parsing content
		t, err := template.New(name).ParseFiles(page)
		if err != nil {
			log.Println("Error while parsing files")
			fmt.Println(err)
			return nil, err // Return the error to the caller
		}

		// find if there are any templates
		// if there are, parse the matching template
		matches, err := filepath.Glob(`C:\Users\USER\Desktop\go\webgo\templates\*.layout.tmpl`)
		if err != nil {
			log.Println("Error while finding layout templates")
			fmt.Println(err)
			return nil, err // Return the error to the caller
		}

		if len(matches) > 0 {
			t, err = t.ParseGlob(`C:\Users\USER\Desktop\go\webgo\templates\*.layout.tmpl`)
			if err != nil {
				log.Println("Error while parsing layout templates")
				log.Println(err)
				return nil, err // Return the error to the caller
			}
		}

		tc[name] = t
	}

	return tc, nil
}
