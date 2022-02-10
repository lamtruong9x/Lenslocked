package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

// Take all files that has .gohtml extension on inside the path LayoutDir
var (
	LayoutDir 	string = "views/layouts/"
	TemplateExt string = ".gohtml"
)
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	//fmt.Println(files)
	if err != nil {
		panic(err)
	}
	return files
}

// Create a View struct to simplify the code in main.go
type View struct {
	Template *template.Template
	Layout string
}
func NewView(layout string, files ...string) *View{
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		t,
		layout,
	}
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}