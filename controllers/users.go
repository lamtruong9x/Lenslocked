package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/schema"
	"lenslocked.com/views"
)

func NewUser() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "users/new"),
	}
}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
} 

func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	form := SignupForm{}
	if err := parseForm(r, form); err!=nil {
		panic(err)
	}
	fmt.Fprintf(w, "%+v", form)
}
func parseForm(r *http.Request, dst interface{}) error {
	if err := r.ParseForm(); err != nil {
		fmt.Print(err)
	}
	dec := schema.NewDecoder()
	if err := dec.Decode(dst, r.PostForm); err != nil {
		return err
	}
	return nil
}
type SignupForm struct {
	Email 		string	`schema:"email"`
	Password 	string	`schema:"password"`
}

type Users struct {
	NewView *views.View
}