package main

import(
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
type card_data struct {
//card data struct that acts like a template
//This template is the used to add all card data
// This is used in the function: 
    Question ,Answer string
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
//Deletes card when given Username,Question, String
// Returns True or False
func Delete_card(Username string, Question string, Answer string) bool { 
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

//Grabs all the cards of a certain user
func Grab_card(Username string) []card_data{
	fmt.Println("Starting server")
	db,err := sql.Open("mysql","root:new_password@tcp(127.0.0.1:3306)/flashcarddb")
	if err != nil {
	panic(err)	}
	var  query string
	query = fmt.Sprintf("SELECT * FROM Users WHERE Username = '%s'",(Username))//Selects everything from user
    rows,err := db.Query(query)
	fmt.Println("HELLLLLLLOOOOOO")
    if err != nil {panic(err)}
    card := card_data{} //Fetches cards from SQL query line by line 
    cards := []card_data{} //will collect all the card data from the "card" array and store the data
	var question string
	var answer string 
	var id int
    for rows.Next(){
        err := rows.Scan(&question,&answer,&id)//Scanning the data
        if err != nil {panic(err)}
        card.Question = question //Add resulted data to card_data struct
        card.Answer = answer
        cards = append(cards, card) // This card_data struct is then added into the final users structures
    }
    defer db.Close()
    return (cards) //Finally returns the structure
}

func main(){
	fmt.Println(Grab_card("Yassa Taiseer"))
}