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

// Handle home "/" path
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil); err != nil {
		panic(err)
	}
}
// Handle contact path "/contact"
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactView.Template.ExecuteTemplate(w, contactView.Layout, nil); err != nil {
		panic(err)
	}
}
func fqa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `<h1>This is FQA page</h1>
	<b>Where all your question will be answered...</b>`)
}
// Handle every not defined path
func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "<h1>Content not found</h1>")
}

// Using gorrilla/mux
func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")

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