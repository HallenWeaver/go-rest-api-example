package character_repository

import (
	"alexandre/gorest/app/model"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CharacterRepository struct {
	characterCollection *mongo.Collection
}

func NewCharacterRepository(client *mongo.Client) (*CharacterRepository, error) {
	characterRepoDatabase := client.Database("charDB")
	characterCollection := characterRepoDatabase.Collection("songs")
	return &CharacterRepository{characterCollection: characterCollection}, nil
}

func (r *CharacterRepository) FindAllByUser(ctx context.Context, ownerID string, count int64) ([]*model.Character, error) {
	characters := make([]*model.Character, 0)
	//Set the limit of the number of record to find
	findOptions := options.Find()
	findOptions.SetLimit(count)
	cursor, err := r.characterCollection.Find(ctx, bson.D{{Key: "owner_id", Value: ownerID}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(ctx) {
		//Create a value into which the single document can be decoded
		var character *model.Character
		err := cursor.Decode(&character)
		if err != nil {
			log.Fatal(err)
		}

		characters = append(characters, character)

	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	//Close the cursor once finished
	cursor.Close(ctx)

	return characters, nil
}

func (r *CharacterRepository) FindByCharacterId(ctx context.Context, ownerId string, characterId primitive.ObjectID) (*model.Character, error) {
	character := &model.Character{}
	err := r.characterCollection.FindOne(ctx, model.Character{ID: characterId, OwnerId: ownerId}).Decode(&character)
	if err != nil {
		return &model.Character{}, err
	}
	return character, nil
}

func (r *CharacterRepository) CreateCharacter(ctx context.Context, newCharacter model.Character) (bool, error) {
	_, err := r.characterCollection.InsertOne(ctx, newCharacter)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *CharacterRepository) UpdateCharacter(ctx context.Context, newCharacter model.Character) (bool, error) {
	_, err := r.characterCollection.UpdateOne(ctx, bson.D{{Key: "_id", Value: &newCharacter.ID}}, newCharacter)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *CharacterRepository) DeleteCharacter(ctx context.Context, ownerId string, characterId primitive.ObjectID) (bool, error) {
	_, err := r.characterCollection.DeleteOne(ctx, bson.D{{Key: "owner_id", Value: ownerId}, {Key: "_id", Value: characterId}})

	if err != nil {
		return false, err
	}

	return true, nil
}
