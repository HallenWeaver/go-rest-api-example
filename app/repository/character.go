package character_repository

import (
	"alexandre/gorest/app/database"
	"alexandre/gorest/app/model"
	"database/sql"
	"fmt"
)

type CharacterRepository struct {
	db *sql.DB
}

func NewCharacterRepository() (*CharacterRepository, error) {
	db, err := database.ConnectDatabase("sqlite3", "./characters.db")
	if err != nil {
		return nil, err
	}

	tableQuery := "CREATE TABLE IF NOT EXISTS characters (id TEXT PRIMARY KEY, owner_id TEXT NOT NULL, name VARCHAR(250) NOT NULL, age INTEGER)"
	err = database.PrepareAndCreateTable(db, tableQuery)
	if err != nil {
		return nil, err
	}

	return &CharacterRepository{db: db}, nil
}

func (r *CharacterRepository) FindAllByUser(ownerID string, count int) ([]*model.Character, error) {
	findQuery := fmt.Sprintf("SELECT id, owner_id, name, age FROM characters WHERE owner_id='%s' LIMIT %d", ownerID, count)
	rows, err := r.db.Query(findQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	characters := make([]*model.Character, 0)
	for rows.Next() {
		character := &model.Character{}
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

	return characters, nil
}

func (r *CharacterRepository) FindByCharacterId(ownerId string, characterId string) (*model.Character, error) {
	findQuery := fmt.Sprintf("SELECT id, owner_id, name, age FROM characters WHERE owner_id='%s' AND id='%s' LIMIT 1", ownerId, characterId)
	row := r.db.QueryRow(findQuery)

	character := &model.Character{}
	sqlErr := row.Scan(&character.Id, &character.OwnerId, &character.Name, &character.Age)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return &model.Character{}, nil
		}
		return &model.Character{}, sqlErr
	}
	return character, nil
}

func (r *CharacterRepository) CreateCharacter(newCharacter model.Character) (bool, error) {
	insertQuery := "INSERT INTO characters(id, owner_id, name, age) VALUES($1, $2, $3, $4) RETURNING id"
	err := r.db.QueryRow(insertQuery, newCharacter.Id, newCharacter.OwnerId, newCharacter.Name, newCharacter.Age).Scan(&newCharacter.Id)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *CharacterRepository) UpdateCharacter(newCharacter model.Character) (bool, error) {
	updateQuery := "UPDATE characters SET owner_id = $1, name = $2, age = $3 WHERE id = $4 AND owner_id = $5"
	err := r.db.
		QueryRow(updateQuery, newCharacter.Id, newCharacter.OwnerId, newCharacter.Name, newCharacter.Age, newCharacter.OwnerId).
		Scan(&newCharacter.Id)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *CharacterRepository) DeleteCharacter(ownerId string, characterId string) (bool, error) {
	deleteQuery := "DELETE from characters where owner_id = $1 AND id = $2"
	_, err := r.db.Exec(deleteQuery, ownerId, characterId)

	if err != nil {
		return false, err
	}

	return true, nil
}
