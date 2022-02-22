package main

import (
	"net/http"

	"github.com/Mohiiuddiin/GoLang_Web_Programming/tree/main/Booking_App/internal/config"
	handler "github.com/Mohiiuddiin/GoLang_Web_Programming/tree/main/Booking_App/internal/handlers"

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

	mux.Get("/home", handler.Repo.Home)
	mux.Get("/about", handler.Repo.About)
	mux.Get("/generals-suite", handler.Repo.Generals)
	mux.Get("/majors-suite", handler.Repo.Majors)
	mux.Get("/search-availability", handler.Repo.Availability)//search-availability
	mux.Post("/search-availability", handler.Repo.PostAvailability)
	mux.Post("/search-availability-json", handler.Repo.AvailabilityJSON)

	mux.Get("/make-reservation", handler.Repo.Reservation)
	mux.Post("/make-reservation", handler.Repo.PostReservation)


	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static",fileServer))

	return mux
}
