package database

import (
	"database/sql"
	"fmt"
	"log"
)

func ConnectDatabase(engine string, path string) (*sql.DB, error) {
	fmt.Println("Trying to connect to database...")
	db, err := sql.Open(engine, path)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return db, err
	}
	return db, nil
}

func PrepareAndCreateTable(dbInstance *sql.DB, query string) error {
	statement, err := dbInstance.Prepare(query)
	if err != nil {
		log.Println("Error in preparing table creation statement")
		return err
	}
	defer statement.Close()

	_, err = statement.Exec()
	if err != nil {
		log.Println("Error in creating table")
		return err
	}

	log.Println("Successfully created table!")
	return nil
}
