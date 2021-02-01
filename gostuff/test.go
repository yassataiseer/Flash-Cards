package main

import ("fmt"
	"html/template"
	"net/http")
var tpl = template.Must(template.ParseFiles("templates/about.html"))
var arg = "Hello there -obi wan"

type Book struct {
    Quote  string
    Author string
}

func index(w http.ResponseWriter, r*http.Request){
	fmt.Fprintf(w, "Hello there")
 }
func about_page(w http.ResponseWriter, r*http.Request){
	book := Book{"Hello there", "Obi-Wan"}
	tpl.Execute(w, book)

}
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about/", about_page)
	http.ListenAndServe(":8000",nil)
}