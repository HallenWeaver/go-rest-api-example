package service

import (
	"alexandre/gorest/app/model"
	"alexandre/gorest/app/repository"

	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Defining Mock Character Repositories

type MockCharacterRepository struct{}

type MockCharacterRepositoryWithErrors struct{}

func (mcr *MockCharacterRepository) FindAllByUser(ctx context.Context, ownerID string, count int64) ([]*model.Character, error) {
	return []*model.Character{}, nil
}
func (mcr *MockCharacterRepository) FindByCharacterId(ctx context.Context, ownerId string, characterId primitive.ObjectID) (*model.Character, error) {
	return &model.Character{}, nil
}

func (mcr *MockCharacterRepository) CreateCharacter(ctx context.Context, newCharacter model.Character) (*model.Character, error) {
	return &newCharacter, nil
}

func (mcr *MockCharacterRepository) UpdateCharacter(ctx context.Context, newCharacter model.Character) (bool, error) {
	return true, nil
}

func (mcr *MockCharacterRepository) DeleteCharacter(ctx context.Context, ownerId string, characterId primitive.ObjectID) (bool, error) {
	return true, nil
}

func (mcrwe *MockCharacterRepositoryWithErrors) FindAllByUser(ctx context.Context, ownerID string, count int64) ([]*model.Character, error) {
	return nil, ErrorMessageTesting
}

func (mcrwe *MockCharacterRepositoryWithErrors) FindByCharacterId(ctx context.Context, ownerId string, characterId primitive.ObjectID) (*model.Character, error) {
	return nil, ErrorMessageTesting
}

func (mcrwe *MockCharacterRepositoryWithErrors) CreateCharacter(ctx context.Context, newCharacter model.Character) (*model.Character, error) {
	return nil, ErrorMessageTesting
}

func (mcrwe *MockCharacterRepositoryWithErrors) UpdateCharacter(ctx context.Context, newCharacter model.Character) (bool, error) {
	return false, ErrorMessageTesting
}

func (mcrwe *MockCharacterRepositoryWithErrors) DeleteCharacter(ctx context.Context, ownerId string, characterId primitive.ObjectID) (bool, error) {
	return false, ErrorMessageTesting
}

type MockCharacterRepositoryWrapper struct {
	characterRepo  repository.ICharacterRepository
	isValidTesting bool
}

// Building Actual Tests
func TestCharacterService(t *testing.T) {
	characterRepos := []MockCharacterRepositoryWrapper{
		{characterRepo: &MockCharacterRepository{}, isValidTesting: true},
		{characterRepo: &MockCharacterRepositoryWithErrors{}, isValidTesting: false},
	}

	for _, characterRepo := range characterRepos {
		characterService := NewCharacterService(characterRepo.characterRepo)

		returnedCharacterList, err := characterService.GetCharacters(context.Background(), "", 0)
		AssertValidityCondition(t, returnedCharacterList, err, characterRepo.isValidTesting)

		returnedCharacter, err := characterService.GetCharacter(context.Background(), "", primitive.NewObjectID())
		AssertValidityCondition(t, returnedCharacter, err, characterRepo.isValidTesting)

		createdCharacter, err := characterService.CreateCharacter(context.Background(), model.Character{})
		AssertValidityCondition(t, createdCharacter, err, characterRepo.isValidTesting)

		updatedCharacter, err := characterService.UpdateCharacter(context.Background(), model.Character{})
		AssertValidityConditionBoolean(t, updatedCharacter, err, characterRepo.isValidTesting)

		deletedCharacter, err := characterService.DeleteCharacter(context.Background(), "", primitive.NewObjectID())
		AssertValidityConditionBoolean(t, deletedCharacter, err, characterRepo.isValidTesting)
	}
}
