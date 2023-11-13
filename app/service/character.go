package service

import (
	"alexandre/gorest/app/model"
	"alexandre/gorest/app/repository"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ICharacterService interface {
	GetCharacters(ctx context.Context, ownerId string, limit int64) ([]*model.Character, error)
	GetCharacter(ctx context.Context, ownerId string, characterId primitive.ObjectID) (*model.Character, error)
	CreateCharacter(ctx context.Context, newCharacter model.Character) (*model.Character, error)
	UpdateCharacter(ctx context.Context, newCharacter model.Character) (bool, error)
	DeleteCharacter(ctx context.Context, ownerId string, characterId primitive.ObjectID) (bool, error)
}

type CharacterService struct {
	CharacterRepository repository.ICharacterRepository
}

func NewCharacterService(characterRepository repository.ICharacterRepository) *CharacterService {
	return &CharacterService{
		CharacterRepository: characterRepository,
	}
}

func (cs *CharacterService) GetCharacters(ctx context.Context, ownerId string, limit int64) ([]*model.Character, error) {
	return cs.CharacterRepository.FindAllByUser(ctx, ownerId, limit)
}

func (cs *CharacterService) GetCharacter(ctx context.Context, ownerId string, characterId primitive.ObjectID) (*model.Character, error) {
	return cs.CharacterRepository.FindByCharacterId(ctx, ownerId, characterId)
}

func (cs *CharacterService) CreateCharacter(ctx context.Context, newCharacter model.Character) (*model.Character, error) {
	return cs.CharacterRepository.CreateCharacter(ctx, newCharacter)
}

func (cs *CharacterService) UpdateCharacter(ctx context.Context, newCharacter model.Character) (bool, error) {
	return cs.CharacterRepository.UpdateCharacter(ctx, newCharacter)
}

func (cs *CharacterService) DeleteCharacter(ctx context.Context, ownerId string, characterId primitive.ObjectID) (bool, error) {
	return cs.CharacterRepository.DeleteCharacter(ctx, ownerId, characterId)
}
