package main
import(
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
	L "./lib"
)
type user_data struct {
//User data struct that acts like a template
//This template is the used to add all user-name data
    Username1 string
	Password1 string
} 
type card_data struct {
	//card data struct that acts like a template
	//This template is the used to add all card data
	// This is used in the function: 
		Question string
		Answer string
}

func db_connect(){
	//Connects to database
	// Not reallyy sure why this is here but is necessary for the functioning of the app :/
	fmt.Println("Starting server")
	db,err := sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/flashcarddb")
	if err != nil {
		panic(err)
	}
	
	defer db.Close()
	fmt.Println("Connected!")
}

func sign_up(w http.ResponseWriter, r*http.Request){
	// redirects the user to sign-up.html
	var tpl = template.Must(template.ParseFiles("templates/sign-up.gohtml"))
	tpl.Execute(w, nil)
 }
 func login(w http.ResponseWriter, r*http.Request){
	// redirects the user to login.html
	var tpl = template.Must(template.ParseFiles("templates/login.gohtml"))
	tpl.Execute(w, nil)
 }
 func login_query(w http.ResponseWriter, r *http.Request){
	// Checks users credentails when loggin in
	r.ParseForm()
	var username  = r.Form["Username"]
	var pswd  = r.Form["pswrd"]
	// grab user data(username&password)
	fmt.Println(username[0],pswd)
	var proceed bool = L.Sign_user_in(username[0],pswd[0])
	//Arguments passed into function Sign_user_in in lib file(users_connector.go)
	// This function returns true or false of the credentials 
	// This will be used to redirect the user
	if proceed == true{
		var tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
		fmt.Println("Starting server")
		db,err := sql.Open("mysql","root:new_password@tcp(127.0.0.1:3306)/flashcarddb")
		if err != nil {
		panic(err)	}
		var  query string
		query = fmt.Sprintf("SELECT * FROM Cards WHERE Username = '%s'",(username[0]))//Selects everything from user
		rows,err := db.Query(query)
		if err != nil {panic(err)}
		var cards []card_data
		//cards := []card_data{} //will collect all the card data from the "card" array and store the data
		var USERNAME string
		var question string
		var answer string 
		var id int
		for rows.Next(){
			err := rows.Scan(&USERNAME,&question,&answer,&id)//Scanning the data
			if err != nil {panic(err)}
			//card.Question = question+//Add resulted data to card_data struct
			//card.Answer = answer
			cards = append(cards, card_data{Question:question,Answer:answer}) // This card_data struct is then added into the final users structures
		}
		defer db.Close()
		fmt.Println(cards)
		tpl.Execute(w,cards)
	}else{
		var tplt = template.Must(template.ParseFiles("templates/error.gohtml"))
		tplt.Execute(w,nil)	
	}
}

func create_user(w http.ResponseWriter, r *http.Request){
	// Takes user data from sign-up.html and makes user
	r.ParseForm()
	var username = r.Form["Username"]
	var password = r.Form["pswrd"]
	//Grabs data from from
	var existing_user bool = L.Existing_user(username[0])
	//Arguments passed into function Existing_User in lib file(users_connector.go)
	// This function checks to see if the username already exists if so will redirect user
	if existing_user == true{
		var tplt = template.Must(template.ParseFiles("templates/error.gohtml"))
		tplt.Execute(w,nil)	
	}else{
		var proceed bool = L.Add_user(username[0],password[0])
		// if the username doesnt exist it will be added 
		//into function Add_user in lib file(users_connector.go)
		//Then will be redirected
		if proceed == true{
			var tplt = template.Must(template.ParseFiles("templates/index.gohtml"))
			tplt.Execute(w,nil)	
		}
	}

}

func main(){
	//Main function simply handles routing
	http.HandleFunc("/", login) 
	http.HandleFunc("/sign-up", sign_up)
	http.HandleFunc("/login_user", login_query )
	http.HandleFunc("/newuser", create_user)
	http.ListenAndServe(":8000",nil)
}