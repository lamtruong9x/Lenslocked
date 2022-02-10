package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"lenslocked.com/views"
)

//Global homeTemplate varibale
var homeView 	*views.View
var contactView *views.View
var fqaView 	*views.View
// Handle home "/" path
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}
// Handle contact path "/contact"
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))
}
// Handle fqa path "/fqa"
func fqa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(fqaView.Render(w, nil))
}
// Handle every not defined path
func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "<h1>Content not found</h1>")
}
// A helper function that handle err
func must(err error) {
	if err != nil {
		panic(err)
	}
}
// Using gorrilla/mux
func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	fqaView = views.NewView("bootstrap", "views/fqa.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")

	nF := http.HandlerFunc(notFound)
	r := mux.NewRouter()
	r.NotFoundHandler = nF
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/fqa", fqa)
	http.ListenAndServe(":3000", r)
}

// Using Gin router
// func main() {
// 	// http.HandleFunc("/home", home)
// 	// http.HandleFunc("/contact", contact)
// 	// http.HandleFunc("/", notFound)
// 	nF := http.HandlerFunc(notFound)
// 	r := chi.NewRouter()
// 	// fmt.Printf("%T\n", r)
// 	// fmt.Printf("%+v\n", r)
// 	r.NotFound(nF)
// 	r.HandleFunc("/", home)
// 	r.HandleFunc("/contact", contact)
// 	r.HandleFunc("/fqa", fqa)
// 	http.ListenAndServe(":3000", r)
// }