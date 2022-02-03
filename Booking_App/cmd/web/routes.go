package main

import (
	"github.com/Mohiiuddiin/GoLang_Web_Programming/tree/main/Booking_App/pkg/config"
	handler "github.com/Mohiiuddiin/GoLang_Web_Programming/tree/main/Booking_App/pkg/handlers"
	"net/http"

	// "github.com/bmizerany/pat"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {

	//using pat
	// mux := pat.New()
	// mux.Get("/home",http.HandlerFunc(handler.Repo.Home))
	// mux.Get("/about",http.HandlerFunc(handler.Repo.About))

	//using go-chi
	mux := chi.NewRouter()	
	
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	
	mux.Get("/Home",handler.Repo.Home)
	mux.Get("/About",handler.Repo.About)
	
	return mux
}