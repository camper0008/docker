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
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	database := os.Getenv("DB_DATABASE")

	connection := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, database)

	db, err := sql.Open("mysql", connection)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func requiredEnvVars() {
	env_vars := []string{"DB_HOST", "DB_USER", "DB_PASSWORD", "DB_DATABASE"}
	for _, env_var := range env_vars {
		if os.Getenv(env_var) == "" {
			log.Fatal(fmt.Errorf("%s environment variable undefined", env_var))
		}
	}
}

func main() {
	db := mySqlConnection()
	http.HandleFunc("/increment", func(rw http.ResponseWriter, r *http.Request) { incrementRequestHandler(rw, r, db) })
	http.HandleFunc("/amount", func(rw http.ResponseWriter, r *http.Request) { amountRequestHandler(rw, r, db) })
	defer db.Close()

	requiredEnvVars()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
