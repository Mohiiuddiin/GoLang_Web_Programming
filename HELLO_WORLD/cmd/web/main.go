package main

import (
	"GOLANG_WEB_DEV/HELLO_WORLD/pkg/config"
	handler "GOLANG_WEB_DEV/HELLO_WORLD/pkg/handlers"
	renderer "GOLANG_WEB_DEV/HELLO_WORLD/pkg/renderers"
	"fmt"
	"log"
	"net/http"
)

const port = ":8080"

func main() {

	var app config.AppConfig
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

	fmt.Println(fmt.Sprintf("Application started on port %s",port))
	// _ = http.ListenAndServe(port,nil)

	srv := &http.Server{
		Addr: port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}