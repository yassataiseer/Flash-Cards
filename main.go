package main
import(
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
	L "./lib"
	"strings"
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
		var Data = L.Grab_card(username[0])
		// Grabs card data from lib file(cards_connector.go)
		user := http.Cookie{
			Name: "Username", Value: username[0],
		}
		http.SetCookie(w, &user)
		//var c = user.Value
		fmt.Println(user)
		tpl.Execute(w,Data)
	}else{
		var tplt = template.Must(template.ParseFiles("templates/error.gohtml"))
		tplt.Execute(w,nil)	
	}
}
func route_flashcard(w http.ResponseWriter, r *http.Request){
	var tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
	var cookie , _ = r.Cookie("Username")
	//Grab cookie value which is the User's Username
	fmt.Println(cookie.Value)
	var Data = L.Grab_card(cookie.Value)
	//Takes the cookie value and passes it into Lib File (cards_connector.go)
	//Grab_data gets all card data for the curent user
	tpl.Execute(w,Data)//Passes it into index.gohtml
}
func delete(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	var card_info string =  r.FormValue("card_info")
	var sorted_data = strings.Split(card_info,"~")
	//fmt.Println(sorted_data[0])
	var cookie , _ = r.Cookie("Username")
	var delete_card = L.Delete_card(cookie.Value,sorted_data[0],sorted_data[1])
	if delete_card == true{
		var tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
		var cookie , _ = r.Cookie("Username")
		//Grab cookie value which is the User's Username
		fmt.Println(cookie.Value)
		var Data = L.Grab_card(cookie.Value)
		//Takes the cookie value and passes it into Lib File (cards_connector.go)
		//Grab_data gets all card data for the curent user
		tpl.Execute(w,Data)//Passes it into index.gohtml
	} else{
		var tpl = template.Must(template.ParseFiles("templates/error.gohtml"))
		tpl.Execute(w,nil)
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
	http.HandleFunc("/login_user", login_query)
	http.HandleFunc("/newuser", create_user)
	http.HandleFunc("/cards",route_flashcard)
	http.HandleFunc("/delete_order",delete)
	http.ListenAndServe(":8000",nil)
}