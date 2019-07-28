// Simple example of creating and using a template in Go
package main

import (
	"html/template"
	"log"
	"net/http"
)

// Type representing a gopher
type Gopher struct {
	Name string
}

// Handler for the hello-gopher route
func helloGopherHandler(w http.ResponseWriter, r *http.Request) {

	var gophername string
	gophername = r.URL.Query().Get("gophername")
	if gophername == "" {
		gophername = "Gopher"
	}
	gopher := Gopher{Name: gophername}
	//d, err := os.Getwd()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//tempFile := d + "/GoBackEnd/section1/templatedemo/templates/greeting.html"
	renderTemplate(w, "./templates/greeting.html", gopher)
	//renderTemplate(w, tempFile, gopher)

}

// Template rendering function
func renderTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Fatal("Error encountered while parsing the template: ", err)
	}
	t.Execute(w, templateData)
}

func main() {
	http.HandleFunc("/hello-gopher", helloGopherHandler)
	http.ListenAndServe(":8080", nil)
}
