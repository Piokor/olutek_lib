package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetClient() (*sqlx.DB, error) {
	connStr := fmt.Sprintf("postgresql://%s:%s@localhost:5432/%s?sslmode=disable", os.Getenv("DBUSER"), os.Getenv("DBPASSWORD"), os.Getenv("DBNAME"))

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return db, nil
}
