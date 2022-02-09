package renderer

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/Mohiiuddiin/GoLang_Web_Programming/tree/main/Booking_App/pkg/config"
	"github.com/Mohiiuddiin/GoLang_Web_Programming/tree/main/Booking_App/pkg/models"
)

var functions = template.FuncMap{}
var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}
func AddDefault(td *models.TemplateData) *models.TemplateData {
	return td
}
func RenderTemplate(rw http.ResponseWriter, temp string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = RenderTemplateTest()
	}
	// tc ,err := RenderTemplateTest()
	// fmt.Println(temp)
	// if err != nil {
	// 	//fmt.Println("Error getting template cache:", err)
	// 	log.Fatal(err)
	// }
	// fmt.Println(tc)

	t, ok := tc[temp]

	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefault(td)
	_ = t.Execute(buf, td)
	_, err := buf.WriteTo(rw)
	if err != nil {
		//fmt.Println("Error getting template cache:", err)
		fmt.Println("Error writing to browser", err)
	}

	// parsedTemplate, _ := template.ParseFiles("./Templates/" + temp)
	// fmt.Println(temp)
	// err = parsedTemplate.Execute(rw, nil)
	// if err != nil {
	// 	fmt.Println("Error parsing template:", err)
	// 	return
	// }
}

func RenderTemplateTest() (map[string]*template.Template, error) {
	myCahhe := map[string]*template.Template{}
	pages, err := filepath.Glob("./Templates/*.page.tmpl")
	if err != nil {
		return myCahhe, err
	}
	//  fmt.Println(pages)
	for _, page := range pages {

		name := filepath.Base(page)
		// fmt.Println("name:"+name+",page:"+page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		//fmt.Println(ts)
		// fmt.Println(*ts)
		// fmt.Println(&ts)
		if err != nil {
			return myCahhe, err
		}
		matches, err := filepath.Glob("./Templates/*.layout.tmpl")
		if err != nil {
			return myCahhe, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./Templates/*.layout.tmpl")
			if err != nil {
				return myCahhe, err
			}
		}
		myCahhe[name] = ts
	}
	return myCahhe, nil
}
