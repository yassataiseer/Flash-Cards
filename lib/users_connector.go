package lib

import(
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
type user_data struct {
//User data struct that acts like a template
//This template is the used to add all user-name data
// This is used in the function: Grab_user_data
    Username1 ,Password1 string
}

//Adds user+password into Users table
func Add_user( Username string,  password string) bool{
	fmt.Println("Starting server")
	db,err := sql.Open("mysql","root:new_password@tcp(127.0.0.1:3306)/flashcarddb")
	if err != nil {
		panic(err)
	}
	add,err := db.Query("INSERT INTO Users (Username,Password) VALUES (?,?)", (Username),(password)) 
    //Takes username and passwords and adds them to the db

	if err != nil {
        panic(err)
        return false
    }
    fmt.Println(add)
    defer db.Close()
    return true 
}
//Checks if username+password exists
func Sign_user_in( user string, password string) bool{
	fmt.Println("Starting server")
	db,err := sql.Open("mysql","root:new_password@tcp(127.0.0.1:3306)/flashcarddb")
	if err != nil {panic(err)}
    var exists bool // Boolean that will hold the value of whether a user exists
    var query string
    query = fmt.Sprintf("SELECT EXISTS(SELECT Username FROM Users WHERE Username = '%s' AND Password = '%s')", (user),(password))
    fmt.Println(query)
    row := db.QueryRow(query).Scan(&exists)//Scans the Selected rows to see if user exists
    //This value is then filled in
    if err != nil && err != sql.ErrNoRows {
        panic(err)
    }
    fmt.Println(row)
    defer db.Close()
    return exists//returns the boolean
}
//Checks to see if user already exists
func Existing_user(user string) bool{
    fmt.Println("Starting server")
	db,err := sql.Open("mysql","root:new_password@tcp(127.0.0.1:3306)/flashcarddb")
	if err != nil {panic(err)}
    var exists bool //Boolean that will hold the value of whether a user exists when scanned
    var query string
    query = fmt.Sprintf("SELECT EXISTS(SELECT Username FROM Users WHERE Username = '%s')", (user))
    //checks to see if user is in Users table and in the Username Column
    //The scanned results are kept in the exists boolean
    fmt.Println(query)
    _ = db.QueryRow(query).Scan(&exists)
    if err != nil && err != sql.ErrNoRows {
        panic(err)
    }
    defer db.Close()
    return exists//returns the boolean
}

func Grab_user_data() []user_data{
    fmt.Println("Starting server")
	db,err := sql.Open("mysql","root:new_password@tcp(127.0.0.1:3306)/flashcarddb")
    if err != nil {panic(err)}
    var query string
    query = fmt.Sprintf("SELECT * FROM Users")
    rows,err := db.Query(query)
    if err != nil {panic(err)}
    user := user_data{}
    users := []user_data{}
    var Username string
    var Password string
    var id int
    for rows.Next(){
        err := rows.Scan(&Username,&Password,&id)
        if err != nil {panic(err)}
        user.Username1 = Username
        user.Password1 = Password

        users = append(users, user)
    }
    defer db.Close()
    return (users)

}

func main(){
}
