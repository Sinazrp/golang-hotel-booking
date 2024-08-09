package render

import (
	"bytes"
	"fmt"
	"github.com/justinas/nosurf"
	"golang-hotel-booking/internal/config"
	"golang-hotel-booking/internal/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}
func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = app.Sessions.PopString(r.Context(), "flash")
	td.Warning = app.Sessions.PopString(r.Context(), "warning")
	td.Error = app.Sessions.PopString(r.Context(), "error")
	td.CSRFToken = nosurf.Token(r)
	return td
}
func RenderTemplate(res http.ResponseWriter, tmpl string, data *models.TemplateData, r *http.Request) {
	tc := map[string]*template.Template{}
	if app.UseCache {
		tc = app.TemplateCache

	} else {
		tc, _ = CreateTemplateCache()
	}

	// create template
	//tc = app.TemplateCache

	//get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("template not found: ", tmpl)
	}
	buf := new(bytes.Buffer)
	td := AddDefaultData(data, r)
	_ = t.Execute(buf, td)

	//render the template

	_, err := buf.WriteTo(res)
	if err != nil {
		log.Fatal(err)
	}

	//parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl+".page.gohtml", "./templates/base.layout.gohtml")
	//
	//err := parsedTemplate.Execute(res, nil)
	//if err != nil {
	//	_, _ = fmt.Fprintf(res, "Error executing template")
	//}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	//get all pages
	pages, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil {
		return myCache, err
	}

	//parse each temp  and associate layout file and  put in map
	for _, page := range pages {
		pageName := filepath.Base(page)
		fmt.Println(pageName)
		//pars template as ts
		ts, err := template.New(pageName).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.gohtml")
		if err != nil {
			return myCache, err
		}
		// use ts to pars layout
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				return myCache, err
			}
		}
		//put cached template in map with file name key
		myCache[pageName] = ts
	} //end of for loop

	//return all cached temp as map
	return myCache, nil
}

//var tc = make(map[string]*template.Template)
//
//func RenderTemplate(res http.ResponseWriter, t string) {
//	var tmpl *template.Template
//	var err error
//
//	//check to see if we already have template in cache
//	_, inMap := tc[t]
//	if !inMap {
//		// create template
//		log.Println("creating new template")
//		err = createTemplateCache(t)
//		if err != nil {
//			log.Fatal("Error creating template cache")
//		}
//	} else {
//		log.Println("using cached template")
//	}
//	tmpl = tc[t]
//	err = tmpl.Execute(res, nil)
//	if err != nil {
//		log.Fatal("Error creating template cache")
//	}
//
//}
//
//func createTemplateCache(t string) error {
//	templatePath := "templates/" + t + ".page.gohtml"
//	templates := []string{
//		templatePath,
//		"templates/base.layout.gohtml",
//	}
//	tmpl, err := template.ParseFiles(templates...)
//	if err != nil {
//		return err
//	}
//	tc[t] = tmpl
//	return nil
//
//}
