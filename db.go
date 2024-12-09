package main

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

type MySQLStorage struct {
	db *sql.DB
}

func NewMySQLStorage(cfg mysql.Config) *MySQLStorage {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MYSQL")
	return &MySQLStorage{db: db}
}

func (s *MySQLStorage) Init() (*sql.DB, error) {
	if err := s.createProjectTable(); err != nil {
		return nil, err
	}
	if err := s.createUserTable(); err != nil {
		return nil, err
	}
	if err := s.createTasksTable(); err != nil {
		return nil, err
	}

	return s.db, nil
}

func (s *MySQLStorage) createProjectTable() error {
	_, err := s.db.Exec(`
	CREATE TABLE IF NOT EXISTS projects (
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		name VARCHAR(255) NOT NULL,
		createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (id)
	) ENGINE = InnoDB DEFAULT CHARSET = utf8
	`)
	return err
}

func (s *MySQLStorage) createUserTable() error {
	_, err := s.db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		email VARCHAR(255) NOT NULL,
		firstName VARCHAR(255) NOT NULL,
		lastName VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,
		createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (id),
		UNIQUE KEY (email)
	) ENGINE = InnoDB DEFAULT CHARSET = utf8
	`)
	return err
}

func (s *MySQLStorage) createTasksTable() error {
	_, err := s.db.Exec(`
	CREATE TABLE IF NOT EXISTS tasks (
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		name VARCHAR(255) NOT NULL,
		status ENUM('TODO','IN_PROGRESS','IN_TESTING','DONE') NOT NULL DEFAULT 'TODO',
		projectID INT UNSIGNED NOT NULL,
		assignedToID INT UNSIGNED,
		createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (id),
		FOREIGN KEY (assignedToID) REFERENCES users(id),
		FOREIGN KEY (projectID) REFERENCES projects(id)
	) ENGINE = InnoDB DEFAULT CHARSET = utf8
	`)
	return err
}
