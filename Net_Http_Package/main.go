package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", img)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	//http.HandleFunc("/", img)
	http.HandleFunc("/mypic.jpg", myPic)

	http.ListenAndServe(":8888", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Welcome To y first page with Golang`)
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Welcome to my about page`)
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Welcome to my contact page`)
}

func img(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="mypic.jpg">	`)
}
func myPic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "mypic.jpg")
}
