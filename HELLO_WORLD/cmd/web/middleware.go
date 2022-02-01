package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func (rw http.ResponseWriter,r *http.Request){
		fmt.Println("Hit the page")
		next.ServeHTTP(rw,r)
	})
}

func NoSurf(next http.Handler) http.Handler {
	CSRFTokenHandler := nosurf.New(next)

	CSRFTokenHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: false,
		SameSite: http.SameSiteLaxMode,
	})
	return CSRFTokenHandler
}