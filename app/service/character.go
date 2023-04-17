package character_service

import (
	"alexandre/gorest/app/model"
	character_repository "alexandre/gorest/app/repository"
)

type CharacterService struct {
	CharacterRepository character_repository.CharacterRepository
}

func NewCharacterService(characterRepository character_repository.CharacterRepository) *CharacterService {
	return &CharacterService{
		CharacterRepository: characterRepository,
	}
}

func (s *CharacterService) GetCharacters(ownerID string, limit int) ([]*model.Character, error) {
	return s.CharacterRepository.FindAllByUser(ownerID, limit)
}
