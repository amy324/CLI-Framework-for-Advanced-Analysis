// dbinteraction.go
package dbinteraction

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func init() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

// ConnectDB connects to the PostgreSQL database.
func ConnectDB() (*sql.DB, error) {
	// Retrieve PostgreSQL details from environment variables
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", user, dbname, password)

	log.Printf("Connecting to PostgreSQL with connection string: %s", connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	

	return db, nil
}
