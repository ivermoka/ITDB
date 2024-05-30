package lib

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)


var db *sql.DB

func Init() *sql.DB{
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
	connect()
	return db
}

func connect() {
	connStr := fmt.Sprintf(os.Getenv("NEON_CONN_STR"))
	fmt.Println("Connection String:", connStr)
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	var version string
	if err := db.QueryRow("select version()").Scan(&version); err != nil {
		panic(err)
	}

	fmt.Printf("version=%s\n", version)
}