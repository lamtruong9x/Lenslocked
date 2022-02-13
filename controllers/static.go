package controllers

import (
	"lenslocked.com/views"
)

type Static struct {
	Home *views.View
	Contact *views.View
	Fqa *views.View
}

func NewStatic() *Static {
	return &Static{
		Home: views.NewView("bootstrap", "static/home"),
		Fqa: views.NewView("bootstrap", "static/fqa"),
		Contact: views.NewView("bootstrap", "static/contact"),
	}
}
