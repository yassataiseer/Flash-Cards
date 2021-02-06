package main

import(
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
)
func db_connect(){
	fmt.Println("Starting server")
	db,err := sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/flashcarddb")
	if err != nil {
		panic(err)
	}
	
	defer db.Close()
	fmt.Println("Connected!")
}
func sign_up(w http.ResponseWriter, r*http.Request){
	var tpl = template.Must(template.ParseFiles("templates/sign-up.html"))
	tpl.Execute(w, nil)
 }
 func login(w http.ResponseWriter, r*http.Request){
	var tpl = template.Must(template.ParseFiles("templates/login.html"))
	tpl.Execute(w, nil)
 }
func main(){
	http.HandleFunc("/", login)
	http.HandleFunc("/sign-up", sign_up)
	http.ListenAndServe(":8000",nil)
}