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
func Add_card(Username string, Question string, Answer string) bool{
	fmt.Println("Starting server")
	db,err := sql.Open("mysql","root:new_password@tcp(127.0.0.1:3306)/flashcarddb")
	if err != nil {
	panic(err)
	return false
	}
	add,err := db.Query("INSERT INTO Cards (Username,Question,Answer) VALUES (?,?,?)", (Username),(Question),(Answer)) 
    //Takes username and passwords and adds them to the db
	if err != nil {
        panic(err)
        return false
    }
    fmt.Println(add)
    defer db.Close()
    return true 
}
//Deletes order when given Username,Question, String
// Returns True or False
func Delete_order(Username string, Question string, Answer string) bool { 
	fmt.Println("Starting server")
	db,err := sql.Open("mysql","root:new_password@tcp(127.0.0.1:3306)/flashcarddb")
	if err != nil {
	panic(err)
	return false
	}
	delete1,err := db.Exec("DELETE FROM Cards WHERE Username = ? AND Question = ? AND Answer = ? ",(Username),(Question),(Answer))
	//Deletes query where the Username,Question and answer match
	if err != nil{
		panic(err)
		return false
	}
	fmt.Println(delete1)
	defer db.Close()
	return true
}

func main(){
}