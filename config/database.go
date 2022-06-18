package config

import (
	"fmt"
    "database/sql"
    _ "github.com/lib/pq"
	"os"
)

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
 }

// DB set up
func DbInit() *sql.DB {
	DB_USER     := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME     := os.Getenv("DB_NAME")

    dbinfo := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		fmt.Println(fmt.Sprintf("DB error :%v", err))
    }

    return db
}