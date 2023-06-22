// Issue 89
// Passing tainted data into sqlx.tx.Select can
// result in sql injection.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Person struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	query := fmt.Sprintf("SELECT * FROM users WHERE username = %s", username)

	connStr := os.Getenv("DbConnStr")
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}

	tx, err := db.Beginx()
	if err != nil {
		log.Fatalln(err)
	}
	defer tx.Rollback()

	person := Person{}
	err = tx.QueryRowx(query).StructScan(&person)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(person)
	err = tx.Commit()
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	http.HandleFunc("/", signupHandler)
	http.ListenAndServe(":8090", nil)
}
