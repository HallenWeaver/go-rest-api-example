package service

import (
	"alexandre/gorest/app/helper"
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
	return nil, helper.ErrorMessageTesting
}

func (mcrwe *MockCharacterRepositoryWithErrors) FindByCharacterId(ctx context.Context, ownerId string, characterId primitive.ObjectID) (*model.Character, error) {
	return nil, helper.ErrorMessageTesting
}

func (mcrwe *MockCharacterRepositoryWithErrors) CreateCharacter(ctx context.Context, newCharacter model.Character) (*model.Character, error) {
	return nil, helper.ErrorMessageTesting
}

func (mcrwe *MockCharacterRepositoryWithErrors) UpdateCharacter(ctx context.Context, newCharacter model.Character) (bool, error) {
	return false, helper.ErrorMessageTesting
}

func (mcrwe *MockCharacterRepositoryWithErrors) DeleteCharacter(ctx context.Context, ownerId string, characterId primitive.ObjectID) (bool, error) {
	return false, helper.ErrorMessageTesting
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
		helper.AssertValidityCondition(t, returnedCharacterList, err, characterRepo.isValidTesting)

		returnedCharacter, err := characterService.GetCharacter(context.Background(), "", primitive.NewObjectID())
		helper.AssertValidityCondition(t, returnedCharacter, err, characterRepo.isValidTesting)

		createdCharacter, err := characterService.CreateCharacter(context.Background(), model.Character{})
		helper.AssertValidityCondition(t, createdCharacter, err, characterRepo.isValidTesting)

		updatedCharacter, err := characterService.UpdateCharacter(context.Background(), model.Character{})
		helper.AssertValidityConditionBoolean(t, updatedCharacter, err, characterRepo.isValidTesting)

		deletedCharacter, err := characterService.DeleteCharacter(context.Background(), "", primitive.NewObjectID())
		helper.AssertValidityConditionBoolean(t, deletedCharacter, err, characterRepo.isValidTesting)
	}
}
