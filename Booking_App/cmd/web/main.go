package main

import (
	"GOLANG_WEB_DEV/HELLO_WORLD/pkg/config"
	handler "GOLANG_WEB_DEV/HELLO_WORLD/pkg/handlers"
	renderer "GOLANG_WEB_DEV/HELLO_WORLD/pkg/renderers"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const port = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false

	session = scs.New()

	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session
	
	tc,err := renderer.RenderTemplateTest()
	if err!=nil {
		log.Fatal("Cant create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo:=handler.NewRepo(&app)
	handler.NewHandlers(repo)
	
	renderer.NewTemplates(&app)


	// http.HandleFunc("/Home",handler.Repo.Home)
	// http.HandleFunc("/About",handler.Repo.About)

	fmt.Println("Application started on port : ",port)
	// _ = http.ListenAndServe(port,nil)

	srv := &http.Server{
		Addr: port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}