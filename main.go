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
    Username1 ,Password1 string
} 

func db_connect(){
	fmt.Println("Starting server")
	db,err := sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/flashcarddb")
	if err != nil {
		panic(err)
	}
	
	defer db.Close()
	fmt.Println("Connected!")
}
func sign_up(w http.ResponseWriter, r*http.Request){
	var tpl = template.Must(template.ParseFiles("templates/sign-up.html"))
	tpl.Execute(w, nil)
 }
 func login(w http.ResponseWriter, r*http.Request){
	var tpl = template.Must(template.ParseFiles("templates/login.html"))
	tpl.Execute(w, nil)
 }
 func login_query(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	var username  = r.Form["Username"]
	var pswd  = r.Form["pswrd"]
	fmt.Println(username[0],pswd)
	var proceed bool = L.Sign_user_in(username[0],pswd[0])
	if proceed == true{
		var tpl = template.Must(template.ParseFiles("templates/index.html"))
		tpl.Execute(w,nil)
	}else{
		var tplt = template.Must(template.ParseFiles("templates/error.html"))
		tplt.Execute(w,nil)	
	}
}

func create_user(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	var username = r.Form["Username"]
	var password = r.Form["pswrd"]
	var existing_user bool = L.Existing_user(username[0])
	if existing_user == true{
		var tplt = template.Must(template.ParseFiles("templates/error.html"))
		tplt.Execute(w,nil)	
	}
	var proceed bool = L.Add_user(username[0],password[0])
	if proceed == true{
		var tplt = template.Must(template.ParseFiles("templates/index.html"))
		tplt.Execute(w,nil)	
	}
}

func main(){
	http.HandleFunc("/", login)
	http.HandleFunc("/sign-up", sign_up)
	http.HandleFunc("/login_user", login_query )
	http.HandleFunc("/newuser", create_user)
	http.ListenAndServe(":8000",nil)
}