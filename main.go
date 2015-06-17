package main 

import "fmt"
import "net/http"
import "io/ioutil"

import(
	 "database/sql"
	 _ "github.com/lib/pq"
 )
import "log"

var INDEX_HTML []byte
var db *sql.DB
var HOST = "localhost"

func main(){
	fmt.Println("starting server on http://localhost:3000")
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/redirect", CreateRedirectHandler)
	http.ListenAndServe(":3000", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request){
	if r.Host == HOST{
	log.Println("creating redirect /")
	w.Write(INDEX_HTML)
	return
	}
	RedirectHandler(w,r)
}

func RedirectHandler (w http.ResponseWriter, r. *http.Request){
	db.QueryRow("SELECT destination FROM redirects WHERE source = $1", r.Host).Scan(&destination)
	http.Redirect(w, r, destination, http.StatusMoved)
}

func CreateRedirectHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	db.Exec("INSERT INTO redirects (source, destination) VALUES($1, $2)", r.Form["source"][0], r.Form["destination"][0])
	log.Println("creating redirect", r.Form)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}


func init(){
	INDEX_HTML, _ = ioutil.ReadFile("./templates/index.html")
	connectToDB()
}

func connectToDB(){
	var err error
	db, err = sql.Open("postgres", "postgres://devmenon:cloud01@localhost:5432/redirectapp_dev?sslmode=disable")
		if err != nil {
			log.Fatal(err)
		}
		
		
	//	age := 21
	//	rows, err := db.Query("SELECT name FROM users WHERE age = $1", age		
	
}