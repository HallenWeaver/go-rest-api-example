package model

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

type Character struct {
	Id      string `json:"id"`
	OwnerId string `json:"owner_id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
}

func init() {
	err := ConnectDatabase()
	if err != nil {
		return
	}

	statement, err := DB.Prepare("CREATE TABLE IF NOT EXISTS characters (id TEXT PRIMARY KEY, owner_id TEXT NOT NULL, name VARCHAR(250) NOT NULL, age INTEGER)")
	if err != nil {
		log.Println("Error in creating table")
	} else {
		log.Println("Successfully created table!")
	}
	statement.Exec()
}

func ConnectDatabase() error {
	fmt.Println("Trying to connect to database...")
	db, err := sql.Open("sqlite3", "./characters.db")
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return err
	}

	DB = db
	return nil
}

func GetCharacters(count int) ([]Character, error) {
	rows, err := DB.Query("SELECT id, owner_id, name, age from characters LIMIT " + strconv.Itoa(count))

	if err != nil {
		fmt.Println("Unable to query database, error: " + err.Error())
		return nil, err
	}

	defer rows.Close()

	characters := make([]Character, 0)

	for rows.Next() {
		character := Character{}
		err = rows.Scan(&character.Id, &character.OwnerId, &character.Name, &character.Age)

		if err != nil {
			return nil, err
		}

		characters = append(characters, character)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return characters, err
}

func GetCharacterById(id string) (Character, error) {

	stmt, err := DB.Prepare("SELECT id, owner_id, name, age from characters WHERE id = ?")

	if err != nil {
		return Character{}, err
	}

	character := Character{}

	sqlErr := stmt.QueryRow(id).Scan(&character.Id, &character.OwnerId, &character.Name, &character.Age)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return Character{}, nil
		}
		return Character{}, sqlErr
	}
	return character, nil
}

func AddCharacter(newCharacter Character) (bool, error) {

	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("INSERT INTO characters (id, owner_id, name, age) VALUES (?, ?, ?, ?)")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(newCharacter.Id, newCharacter.OwnerId, newCharacter.Name, newCharacter.Age)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}

func UpdateCharacter(character Character) (bool, error) {

	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("UPDATE characters SET owner_id = $1, name = $2, age = $3 WHERE Id = $4")

	if err != nil {
		fmt.Println("Prepare Error: " + err.Error())
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(character.OwnerId, character.Name, character.Age, character.Id)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}

func DeleteCharacter(characterId string) (bool, error) {

	tx, err := DB.Begin()

	if err != nil {
		return false, err
	}

	stmt, err := DB.Prepare("DELETE from characters where id = ?")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(characterId)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}
