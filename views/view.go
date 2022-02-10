package views

import "html/template"

type View struct {
	Template *template.Template
	Layout string
}
// Add footer to the page
func NewView(layout string, files ...string) *View{
	files = append(files, 
		"views/layouts/footer.gohtml",
		"views/layouts/bootstrap.gohtml", 
		"views/layouts/navbar.gohtml")
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		t,
		layout,
	}
}