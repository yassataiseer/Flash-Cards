package main

import(
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	fmt.Println("Starting server")
	db,err := sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/flashcarddb")
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected!")
	_, err1 := db.Exec("CREATE TABLE IF NOT EXISTS Users (Username VARCHAR(255), Password VARCHAR(255),id INT AUTO_INCREMENT PRIMARY KEY)")
	if err1 != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("Succesfully")
}