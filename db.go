package main

import (
	"database/sql" // Importing the database/sql package for SQL database handling
	"github.com/go-sql-driver/mysql" // Importing the MySQL driver for Go to interact with MySQL databases
	"log" // Importing the log package for logging
)

// MySQLStorage struct represents the storage layer that connects to the MySQL database.
// It holds a reference to an open database connection.
type MySQLStorage struct {
	db *sql.DB // db is a pointer to the SQL database connection object
}

// NewMySQLStorage is a constructor function that creates and returns a new instance of MySQLStorage.
// It accepts a MySQL config (mysql.Config) for setting up the database connection.
func NewMySQLStorage(cfg mysql.Config) *MySQLStorage { // function should start with 'func'
	// Opening a connection to the MySQL database using the provided configuration
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		// If there is an error opening the connection, log the error and terminate the program
		log.Fatal(err)
	}

	// Ping the database to ensure that the connection is successful
	err = db.Ping()
	if err != nil {
		// If there is an error pinging the database, log the error and terminate the program
		log.Fatal(err)
	}

	// Log a message indicating that the connection to the MySQL database was successful
	log.Println("Connected to MYSQL")

	// Return a pointer to a new MySQLStorage instance, with the opened database connection
	return &MySQLStorage{db: db} // Return the MySQLStorage object containing the db connection
}

// Init is a method on the MySQLStorage struct that initializes the database connection and prepares it for use.
// Currently, it only returns the existing database connection.
func (s *MySQLStorage) Init() (*sql.DB, error) {
	// Placeholder for future database table initialization or other setup code.
	// Currently, it just returns the existing database connection without any additional setup.
	return s.db, nil
}
