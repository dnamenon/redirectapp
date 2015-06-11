package main 

import "fmt"
import "net/http"
import "io/ioutil"

var INDEX_HTML []byte

func main(){
	fmt.Println("starting server on http://localhost:3000")
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/redirect", CreateRedirectHandler)
	http.ListenAndServe(":3000", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("creating redirect /")
	w.Write(INDEX_HTML)
}

func CreateRedirectHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println("creating redirect", r.Form)
}


func init(){
	INDEX_HTML, _ = ioutil.ReadFile("./templates/index.html")
	
}