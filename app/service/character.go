package service

import (
	"alexandre/gorest/app/model"
	"alexandre/gorest/app/repository"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CharacterService struct {
	CharacterRepository repository.CharacterRepository
}

func NewCharacterService(characterRepository repository.CharacterRepository) *CharacterService {
	return &CharacterService{
		CharacterRepository: characterRepository,
	}
}

func (s *CharacterService) GetCharacters(ctx context.Context, ownerId string, limit int64) ([]*model.Character, error) {
	return s.CharacterRepository.FindAllByUser(ctx, ownerId, limit)
}

func (s *CharacterService) GetCharacter(ctx context.Context, ownerId string, characterId primitive.ObjectID) (*model.Character, error) {
	return s.CharacterRepository.FindByCharacterId(ctx, ownerId, characterId)
}

func (s *CharacterService) CreateCharacter(ctx context.Context, newCharacter model.Character) (bool, error) {
	return s.CharacterRepository.CreateCharacter(ctx, newCharacter)
}

func (s *CharacterService) UpdateCharacter(ctx context.Context, newCharacter model.Character) (bool, error) {
	return s.CharacterRepository.UpdateCharacter(ctx, newCharacter)
}

func (s *CharacterService) DeleteCharacter(ctx context.Context, ownerId string, characterId primitive.ObjectID) (bool, error) {
	return s.CharacterRepository.DeleteCharacter(ctx, ownerId, characterId)
}
