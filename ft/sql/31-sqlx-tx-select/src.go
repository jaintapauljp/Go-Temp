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
	db, err := sqlx.Open("sqlite3", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Beginx()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	people := []Person{}
	err = tx.Select(&people, query)
	if err != nil {
		log.Fatal(err)
	}

	for _, person := range people {
		fmt.Printf("id: %d, name: %s, age: %d\n", person.ID, person.Name, person.Age)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", signupHandler)
	http.ListenAndServe(":8090", nil)
}
