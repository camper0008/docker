package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func incrementRequestHandler(rw http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != "POST" {
		return
	}
	query := "INSERT INTO clicks VALUES()"
	_, err := db.Exec(query)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, "500 Internal Server Error")
		return
	}
	rw.Write([]byte("200 OK"))
}

func amountRequestHandler(rw http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != "GET" {
		return
	}
	var amount string
	query := "SELECT COUNT(*) AS amount FROM clicks"
	err := db.QueryRow(query).Scan(&amount)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, "500 Internal Server Error")
		return
	}
	rw.Write([]byte(amount))
}

func mySqlConnection() *sql.DB {
	mariadb_user := "example_user"
	mariadb_pass := "example_password"
	mariadb_host := "db" // this refers to the mariadb container from docker
	mariadb_db := "example_db"

	connection := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", mariadb_user, mariadb_pass, mariadb_host, mariadb_db)

	log.Printf("attempting connection to %q:", connection)

	db, err := sql.Open("mysql", connection)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func main() {
	db := mySqlConnection()
	http.HandleFunc("/increment", func(rw http.ResponseWriter, r *http.Request) { incrementRequestHandler(rw, r, db) })
	http.HandleFunc("/amount", func(rw http.ResponseWriter, r *http.Request) { amountRequestHandler(rw, r, db) })
	defer db.Close()

	if os.Getenv("PORT") == "" {
		log.Fatal(fmt.Errorf("PORT environment variable undefined"))
	}

	portString := fmt.Sprintf(":%s", os.Getenv("PORT"))

	log.Fatal(http.ListenAndServe(portString, nil))
}
