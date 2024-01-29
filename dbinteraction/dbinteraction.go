// Interacts with MySQL database
package dbinteraction

import (
	"database/sql"
	"fmt"
	"log"

	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)

	log.Printf("Connecting to MySQL")

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Println("Error opening database connection:", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println("Error pinging database:", err)
		return nil, err
	}

	log.Println("Connected to the database successfully!")

	return db, nil
}

// ...
func InsertAnalysisResults(db *sql.DB, scriptPath string, accuracy float64) error {
	// Prepare the SQL statement
	stmt, err := db.Prepare("INSERT INTO analysis_results (script_name, accuracy) VALUES (?, ?)")
	if err != nil {
		return fmt.Errorf("error preparing SQL statement: %v", err)
	}
	defer stmt.Close()

	// Execute the SQL statement
	_, err = stmt.Exec(scriptPath, accuracy)
	if err != nil {
		return fmt.Errorf("error executing SQL statement: %v", err)
	}

	fmt.Println("Data inserted into the database successfully!")
	return nil
}
