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
	
	defer db.Close()
	fmt.Println("Connected!")
}