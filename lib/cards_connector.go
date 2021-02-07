package lib

import(
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
type card_data struct {
//card data struct that acts like a template
//This template is the used to add all card data
// This is used in the function: 
    Title ,Value string
}

//Will grab all flashcard's data for the certain user
func Grab_card_data(Username string){
	fmt.Println("Starting server")
	db,err := sql.Open("mysql","root:new_password@tcp(127.0.0.1:3306)/flashcarddb")
	if err != nil {panic(err)}
	fmt.Println(db)
}
