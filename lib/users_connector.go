package main

import(
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
type user_data struct {
    Username1 string
    Password1 string
}







//Adds user+password into Users table
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
//Checks if username+password exists
func sign_user_in( user string, password string) bool{
	fmt.Println("Starting server")
	db,err := sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/flashcarddb")
	if err != nil {panic(err)}
    var exists bool
    var query string
    query = fmt.Sprintf("SELECT EXISTS(SELECT Username FROM Users WHERE Username = '%s' AND Password = '%s')", (user),(password))
    fmt.Println(query)
    row := db.QueryRow(query).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
        panic(err)
    }
    fmt.Println(row)
    defer db.Close()
    return exists
}
//Checks to see if user already exists
func existing_user(user string) bool{
    fmt.Println("Starting server")
	db,err := sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/flashcarddb")
	if err != nil {panic(err)}
    var exists bool
    var query string
    query = fmt.Sprintf("SELECT EXISTS(SELECT Username FROM Users WHERE Username = '%s')", (user))
    fmt.Println(query)
    _ = db.QueryRow(query).Scan(&exists)
    if err != nil && err != sql.ErrNoRows {
        panic(err)
    }
    defer db.Close()
    return exists
}

func main(){
//fmt.Println(existing_user("Yassa Taiseer"))
}
