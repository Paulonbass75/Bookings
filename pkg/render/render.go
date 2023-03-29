package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Paulonbass75/bookings/pkg/config"
	// "github.com/Paulonbass75/hwgo/pkg/handlers"
	"github.com/Paulonbass75/bookings/pkg/models"
)

var functions = template.FuncMap{

}

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates (a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}


func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template

	if app.UseCache {

		tc = app.TemplateCache

		} else {
		tc, _ = CreateTemplateCache()
	}

	
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
		
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, td)
	
	
	//render template 
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser", err)
	}
}


func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	
	// get all files named *.page.tmpl from templates folder
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
	return myCache, err

	}

	// loop through pages and create template cache for each page
	for _, page := range pages {
		name := filepath.Base(page)		
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		// add any layout files to the template set
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}

return myCache, nil
}
