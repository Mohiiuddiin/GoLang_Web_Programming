package main

import (
	"fmt"
	"net/http"
)

func main() {

	//http.Handle("/", http.FileServer(http.Dir("images")))
	http.Handle("/r/", http.StripPrefix("/r/", http.FileServer(http.Dir("images"))))
	http.HandleFunc("/", home)
	http.HandleFunc("/img", img)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	//http.HandleFunc("/mypic.jpg", myPic)
	//http.HandleFunc("/logemail", email)

	http.ListenAndServe(":8888", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Welcome To your first page with Golang`)
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Welcome to my about page`)
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Welcome to my contact page`)

}

func img(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<img src=\"r/mypic.jpg\" />")
}
func myPic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "mypic.jpg")
}
func email(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr)
	for key, value := range r.Header {
		fmt.Println(key, value)
	}
	http.ServeFile(w, r, "email_open_log_pic.gif")
}
