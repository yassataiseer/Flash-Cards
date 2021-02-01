
package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strings"
)
type Data struct {
    Username  string
    password string
}
func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
    // attention: If you do not call ParseForm method, the following data can not be obtained form
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello astaxie!") // write data to response
}

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        t, _ := template.ParseFiles("templates/login.html")
        t.Execute(w, nil)
    } else {
        r.ParseForm()
        // logic part of log in
        fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
        fmt.Println("hello")
        var username  = r.Form["username"]
        var pswd  = r.Form["password"]
        username1 := fmt.Sprint(username)
        username2 := strings.Replace(username1, "[", "", -1)
        pswd1 := fmt.Sprint(pswd)
        pswd2 := strings.Replace(pswd1, "[", "", -1)
        data := Data{username2,pswd2}
        fmt.Println(data)
		x, _ := template.ParseFiles("templates/release.html")
		fmt.Println(pswd2)
        x.Execute(w, data)
		
    }
}
func test(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("templates/other.html")
    t.Execute(w, nil)
}

func main() {
    http.HandleFunc("/", sayhelloName) // setting router rule
    http.HandleFunc("/login", login)
    http.HandleFunc("/other",test )
    err := http.ListenAndServe(":9090", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}