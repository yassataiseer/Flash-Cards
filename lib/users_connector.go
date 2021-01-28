package main

import(
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func add_user( Username string,  password string) bool{
	fmt.Println("Starting server")
	db,err := sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/flashcarddb")
	if err != nil {
		panic(err)
	}

	add,err := db.Query("INSERT INTO Users (Username,Password) VALUES (?,?)", (Username),(password))
	if err != nil {
        panic(err)
        return false
    }
    fmt.Println(add)
    defer db.Close()
    return true 
}
func grab_user(){
	fmt.Println("Starting server")
	db,err := sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/flashcarddb")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
func main(){
}
