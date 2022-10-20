package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)
func init() {

    err := godotenv.Load(".env")

    if err != nil {
        log.Fatal("Error loading .env file")
    }
	
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	pswd := os.Getenv("MYSQL_PASSWORD")
	//(driver name, data source name)
	db, err := sql.Open("mysql", "root:"+pswd+"@tcp(localhost:3306)/testaccount")
	if err != nil {
		fmt.Println("error validating sql.Open arguments")
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("error verifying connection with db.Ping")
		panic(err.Error())
	}
	/*
	insert, err := db.Query("INSERT INTO `testaccount`.`user` (`LastName`, `FirstName`, `Hash`) VALUES ('Doe', 'John', '123');")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	*/
	fmt.Println("Successful Connection to Database!")
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/registerauth", registerAuthHandler)
	http.ListenAndServe(":8080", nil)
	
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	err := templates.ExecuteTemplate(w, tmpl+".html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("******rootHandler running ******")
	http.Redirect(w, r, "/register", http.StatusFound)
}

// registerHandler serves form for registering new users
func registerHandler(w http.ResponseWriter, r *http.Request) { 
	fmt.Println("******registerHandler running ******")
	fmt.Println("output: ", os.Getenv("MYSQL_PASSWORD"))
	renderTemplate(w, "register")

}

// registerAuthHandler creates new user in database
func registerAuthHandler(w http.ResponseWriter, r *http.Request) { 
	fmt.Println("******registerAuthHandler running ******")
	/* 
		1. check username criteria.
		2. check passsword criteria.
		3. check first and last name criteria.
		4. check if username already exists in database.
		5. create bcrypt hash from password.

	*/


}