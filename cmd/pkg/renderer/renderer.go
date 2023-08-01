package renderer

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, filename string) {
	// create templates cache
	tmpl, err := createRenderedTemplate()
	if err != nil {
		log.Println(err)
	}
	// get requested template from cache
	t := tmpl[filename]
	buf := new(bytes.Buffer)
	// write the template into buf first so rendering process can run smoothly and
	// we can modify buf as we go on
	err = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}
	// render the template cache
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func createRenderedTemplate() (map[string]*template.Template, error) {
	tmpCache := map[string]*template.Template{}
	// get all the template pages
	pages, _ := filepath.Glob("./templates/*.page.tmpl")
	// range through pages
	for _, page := range pages {
		// set name to the last element of the page, which means the file name
		name := filepath.Base(page)
		// create template
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return tmpCache, err
		}
		// search for layout
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		// check if ther is layout found
		if len(matches) > 0 {
			// incorporate the layout into template
			// override
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return tmpCache, err
			}
		}
		// map ts to name
		tmpCache[name] = ts
	}
	return tmpCache, nil
}
