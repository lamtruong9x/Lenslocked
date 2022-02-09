package main

import (
	"html/template"
	"os"
)

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := struct {
		Name string
		Course string
		LoI []int
		Map map[string]int
	} {
		Name: "Truong",
		Course: "Web Development in Go",
		LoI: []int{1, 2, 4, 6, 7},
		Map: map[string]int {"Truong": 23, "Yen": 18, "DM": 17},
	}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}