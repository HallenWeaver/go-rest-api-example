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

func (s *CharacterService) GetCharacters(ownerId string, limit int) ([]*model.Character, error) {
	return s.CharacterRepository.FindAllByUser(ownerId, limit)
}

func (s *CharacterService) GetCharacter(ownerId string, characterId string) (*model.Character, error) {
	return s.CharacterRepository.FindByCharacterId(ownerId, characterId)
}

func (s *CharacterService) CreateCharacter(newCharacter model.Character) (bool, error) {
	return s.CharacterRepository.CreateCharacter(newCharacter)
}

func (s *CharacterService) UpdateCharacter(newCharacter model.Character) (bool, error) {
	return s.CharacterRepository.UpdateCharacter(newCharacter)
}

func (s *CharacterService) DeleteCharacter(ownerId string, characterId string) (bool, error) {
	return s.CharacterRepository.DeleteCharacter(ownerId, characterId)
}
