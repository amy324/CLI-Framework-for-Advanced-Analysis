package dbinteraction

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// FieldType represents the data type of a field.
type FieldType string

const (
	IntType    FieldType = "int"
	StringType FieldType = "string"
	// TODO Add more types as needed
)

// FieldDefinition represents the definition of a field, including its name and type.
type FieldDefinition struct {
	Name string
	Type FieldType
}

// GenericData represents a generic structure for user-defined data.
type GenericData map[string]interface{}

func init() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

// ConnectDB connects to the MySQL database.
func ConnectDB() (*sql.DB, error) {
	// Retrieve MySQL details from environment variables
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DB")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)

	log.Printf("Connecting to MySQL with connection string: %s", connStr)

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// InsertGenericData inserts generic data into the database.
func InsertGenericData(db *sql.DB, tableName string, data GenericData, fieldDefinitions []FieldDefinition) error {
	// Validate data against field definitions
	if err := validateData(data, fieldDefinitions); err != nil {
		return err
	}

	// Convert map keys and values into slices
	var columns []string
	var values []interface{}
	for column, value := range data {
		columns = append(columns, column)
		values = append(values, value)
	}

	// Build the SQL query dynamically
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, join(columns, ", "), placeholders(len(columns)))

	// Execute the query
	_, err := db.Exec(query, values...)
	if err != nil {
		return err
	}

	fmt.Println("Data inserted successfully")
	return nil
}

// validateData checks if the provided data matches the specified field definitions.
func validateData(data GenericData, fieldDefinitions []FieldDefinition) error {
	for _, definition := range fieldDefinitions {
		if value, ok := data[definition.Name]; ok {
			if !isValidType(value, definition.Type) {
				return fmt.Errorf("invalid type for field %s, expected %s", definition.Name, definition.Type)
			}
		} else {
			return fmt.Errorf("missing value for field %s", definition.Name)
		}
	}
	return nil
}

// isValidType checks if the value has the expected type.
func isValidType(value interface{}, expectedType FieldType) bool {
	switch expectedType {
	case IntType:
		_, ok := value.(int)
		return ok
	case StringType:
		_, ok := value.(string)
		return ok
	// Add more cases for other types
	default:
		return false
	}
}

// join concatenates strings with a separator.
func join(parts []string, sep string) string {
	return strings.Join(parts, sep)
}

// placeholders generates placeholders for a given count.
func placeholders(count int) string {
	return strings.Join(make([]string, count), ", ")
}
