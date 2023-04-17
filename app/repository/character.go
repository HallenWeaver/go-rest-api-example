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
