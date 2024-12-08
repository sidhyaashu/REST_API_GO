package main

import "database/sql" // Importing the sql package for database handling

// Store is an interface that defines the methods that the storage layer (e.g., database interactions) must implement.
// It currently contains a method for creating a user, but this can be expanded as needed for other operations.
type Store interface {
	// CreateUser defines the method signature for creating a user in the database.
	// It returns an error if the operation fails.
	CreateUser() error
}

// Storage struct represents the implementation of the Store interface, specifically interacting with the database.
// It holds a reference to the database connection (`*sql.DB`), which is used for executing queries and interacting with the database.
type Storage struct {
	db *sql.DB // db is a pointer to an open SQL database connection
}

// NewStore is a constructor function that creates and returns a new instance of Storage.
// It takes a pointer to an open database connection (`*sql.DB`) as a parameter.
func NewStore(db *sql.DB) *Storage {
	// Returning a new instance of Storage with the provided database connection
	return &Storage{
		db: db, // Set the db field to the passed database connection
	}
}

// CreateUser is a method on the Storage struct that implements the Store interface's CreateUser method.
// Currently, it simply returns `nil`, indicating that the operation is a placeholder and doesn't perform any actual work.
func (s *Storage) CreateUser() error {
	// Placeholder method for creating a user. It doesn't perform any database operations at the moment.
	// In a real implementation, this method would execute an SQL query to insert a new user into the database.
	return nil
}
