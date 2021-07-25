package main

import (
	"html/template"
	"log"
	"net/http"
)

func TemplateHandler(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("template/index.html")
	if err != nil {
		log.Println("Error parsing template")
		return
	}

	data := map[string]string{
		"pesan": "Hello world! Ini dari Golang!",
	}

	err = template.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		log.Println("Error executing template")
		return
	}
}
