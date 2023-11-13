package handler

import (
	"alexandre/gorest/app/helper"
	"alexandre/gorest/app/model"
	"alexandre/gorest/app/service"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockCharacterService struct{}

type MockCharacterServiceWithErrors struct{}

func (mcs *MockCharacterService) GetCharacters(ctx context.Context, ownerId string, limit int64) ([]*model.Character, error) {
	return []*model.Character{}, nil
}
func (mcs *MockCharacterService) GetCharacter(ctx context.Context, ownerId string, characterId primitive.ObjectID) (*model.Character, error) {
	return &model.Character{}, nil
}
func (mcs *MockCharacterService) CreateCharacter(ctx context.Context, newCharacter model.Character) (*model.Character, error) {
	return &model.Character{}, nil
}
func (mcs *MockCharacterService) UpdateCharacter(ctx context.Context, newCharacter model.Character) (bool, error) {
	return true, nil
}
func (mcs *MockCharacterService) DeleteCharacter(ctx context.Context, ownerId string, characterId primitive.ObjectID) (bool, error) {
	return true, nil
}

func (mcswe *MockCharacterServiceWithErrors) GetCharacters(ctx context.Context, ownerId string, limit int64) ([]*model.Character, error) {
	return []*model.Character{}, helper.ErrorMessageTesting
}
func (mcswe *MockCharacterServiceWithErrors) GetCharacter(ctx context.Context, ownerId string, characterId primitive.ObjectID) (*model.Character, error) {
	return nil, helper.ErrorMessageTesting
}
func (mcswe *MockCharacterServiceWithErrors) CreateCharacter(ctx context.Context, newCharacter model.Character) (*model.Character, error) {
	return nil, helper.ErrorMessageTesting
}
func (mcswe *MockCharacterServiceWithErrors) UpdateCharacter(ctx context.Context, newCharacter model.Character) (bool, error) {
	return false, helper.ErrorMessageTesting
}
func (mcswe *MockCharacterServiceWithErrors) DeleteCharacter(ctx context.Context, ownerId string, characterId primitive.ObjectID) (bool, error) {
	return false, helper.ErrorMessageTesting
}

type MockCharacterTestCaseWrapper struct {
	CharacterService      service.ICharacterService
	IsValidTesting        bool
	IsValidAuthentication bool
	ExpectedStatusCode    int
	Args                  map[string]interface{}
}

func TestCharacterHandlerGetCharacters(t *testing.T) {
	testCases := map[string]MockCharacterTestCaseWrapper{
		"Normal Path Test": {
			CharacterService:      &MockCharacterService{},
			IsValidTesting:        true,
			IsValidAuthentication: true,
			ExpectedStatusCode:    http.StatusOK,
		},
		"Unauthorized User Test": {
			CharacterService:      &MockCharacterService{},
			IsValidTesting:        true,
			IsValidAuthentication: false,
			ExpectedStatusCode:    http.StatusUnauthorized,
		},
		"Character Service Error Test": {
			CharacterService:      &MockCharacterServiceWithErrors{},
			IsValidTesting:        false,
			IsValidAuthentication: true,
			ExpectedStatusCode:    http.StatusInternalServerError,
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			characterHandler := NewCharacterHandler(testCase.CharacterService)

			router := gin.Default()
			router.GET("/character", characterHandler.GetCharacters)

			req := httptest.NewRequest("GET", "/character", http.NoBody)
			tokenString, err := helper.GenerateJWT(primitive.NewObjectID().Hex())
			assert.Nil(t, err)

			if testCase.IsValidAuthentication {
				req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", tokenString))
			} else {
				req.Header.Set("Authorization", "invalid token")
			}
			httpRecorder := httptest.NewRecorder()
			router.ServeHTTP(httpRecorder, req)

			assert.Equal(t, testCase.ExpectedStatusCode, httpRecorder.Code)
		})
	}
}

func TestCharacterHandlerGetCharacter(t *testing.T) {
	testCases := map[string]MockCharacterTestCaseWrapper{
		"Normal Path Test": {
			CharacterService:      &MockCharacterService{},
			IsValidTesting:        true,
			IsValidAuthentication: true,
			ExpectedStatusCode:    http.StatusOK,
			Args:                  map[string]interface{}{"ID": primitive.NewObjectID().Hex()},
		},
		"Unauthorized User Test": {
			CharacterService:      &MockCharacterService{},
			IsValidTesting:        true,
			IsValidAuthentication: false,
			ExpectedStatusCode:    http.StatusUnauthorized,
			Args:                  map[string]interface{}{"ID": primitive.NewObjectID().Hex()},
		},
		"Invalid Character ID Test": {
			CharacterService:      &MockCharacterService{},
			IsValidTesting:        true,
			IsValidAuthentication: true,
			ExpectedStatusCode:    http.StatusInternalServerError,
			Args:                  map[string]interface{}{"ID": "abc"},
		},
		"Character Service Error Test": {
			CharacterService:      &MockCharacterServiceWithErrors{},
			IsValidTesting:        false,
			IsValidAuthentication: true,
			ExpectedStatusCode:    http.StatusInternalServerError,
			Args:                  map[string]interface{}{"ID": primitive.NewObjectID().Hex()},
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			characterHandler := NewCharacterHandler(testCase.CharacterService)

			router := gin.Default()
			router.GET("/character/:id", characterHandler.GetCharacter)

			endpoint := fmt.Sprintf("/character/%v", testCase.Args["ID"])
			req := httptest.NewRequest("GET", endpoint, http.NoBody)
			tokenString, err := helper.GenerateJWT(primitive.NewObjectID().Hex())
			assert.Nil(t, err)

			if testCase.IsValidAuthentication {
				req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", tokenString))
			} else {
				req.Header.Set("Authorization", "invalid token")
			}
			httpRecorder := httptest.NewRecorder()
			router.ServeHTTP(httpRecorder, req)

			assert.Equal(t, testCase.ExpectedStatusCode, httpRecorder.Code)
		})
	}
}

func TestCharacterHandlerCreateCharacter(t *testing.T) {
	testCases := map[string]MockCharacterTestCaseWrapper{
		"Normal Path Test": {
			CharacterService:      &MockCharacterService{},
			IsValidTesting:        true,
			IsValidAuthentication: true,
			ExpectedStatusCode:    http.StatusCreated,
			Args: map[string]interface{}{"Body": model.Character{
				Name: "Hallen Weaver",
				Age:  20,
			}},
		},
		"Unauthorized User Test": {
			CharacterService:      &MockCharacterService{},
			IsValidTesting:        true,
			IsValidAuthentication: false,
			ExpectedStatusCode:    http.StatusUnauthorized,
			Args: map[string]interface{}{"Body": model.Character{
				Name: "Hallen Weaver",
				Age:  20,
			}},
		},
		"Invalid Payload Test": {
			CharacterService:      &MockCharacterServiceWithErrors{},
			IsValidTesting:        false,
			IsValidAuthentication: true,
			ExpectedStatusCode:    http.StatusBadRequest,
			Args:                  map[string]interface{}{"Body": "{"},
		},
		"Character Service Error Test": {
			CharacterService:      &MockCharacterServiceWithErrors{},
			IsValidTesting:        false,
			IsValidAuthentication: true,
			ExpectedStatusCode:    http.StatusBadRequest,
			Args: map[string]interface{}{"Body": model.Character{
				Name: "Hallen Weaver",
				Age:  20,
			}},
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			characterHandler := NewCharacterHandler(testCase.CharacterService)

			router := gin.Default()
			router.POST("/character", characterHandler.CreateCharacter)

			jsonValue, err := json.Marshal(testCase.Args["Body"])
			assert.Nil(t, err)

			req := httptest.NewRequest("POST", "/character", bytes.NewBuffer(jsonValue))
			tokenString, err := helper.GenerateJWT(primitive.NewObjectID().Hex())
			assert.Nil(t, err)

			if testCase.IsValidAuthentication {
				req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", tokenString))
			} else {
				req.Header.Set("Authorization", "invalid token")
			}
			httpRecorder := httptest.NewRecorder()
			router.ServeHTTP(httpRecorder, req)

			assert.Equal(t, testCase.ExpectedStatusCode, httpRecorder.Code)
		})
	}
}

func TestCharacterHandlerUpdateCharacter(t *testing.T) {
	testCases := map[string]MockCharacterTestCaseWrapper{
		"Normal Path Test": {
			CharacterService:      &MockCharacterService{},
			IsValidTesting:        true,
			IsValidAuthentication: true,
			ExpectedStatusCode:    http.StatusOK,
			Args: map[string]interface{}{
				"Body": model.Character{
					Name: "Hallen Weaver",
					Age:  21,
				},
				"ID": primitive.NewObjectID().Hex(),
			},
		},
		"Unauthorized User Test": {
			CharacterService:      &MockCharacterService{},
			IsValidTesting:        true,
			IsValidAuthentication: false,
			ExpectedStatusCode:    http.StatusUnauthorized,
			Args: map[string]interface{}{
				"Body": model.Character{
					Name: "Hallen Weaver",
					Age:  21,
				},
				"ID": primitive.NewObjectID().Hex(),
			},
		},
		"Invalid ID Test": {
			CharacterService:      &MockCharacterService{},
			IsValidTesting:        true,
			IsValidAuthentication: true,
			ExpectedStatusCode:    http.StatusBadRequest,
			Args: map[string]interface{}{
				"Body": model.Character{
					Name: "Hallen Weaver",
					Age:  21,
				},
				"ID": "abc",
			},
		},
		"Invalid Payload Test": {
			CharacterService:      &MockCharacterService{},
			IsValidTesting:        true,
			IsValidAuthentication: true,
			ExpectedStatusCode:    http.StatusBadRequest,
			Args: map[string]interface{}{
				"Body": "{",
				"ID":   primitive.NewObjectID().Hex(),
			},
		},
		"Character Service Error Test": {
			CharacterService:      &MockCharacterServiceWithErrors{},
			IsValidTesting:        false,
			IsValidAuthentication: true,
			ExpectedStatusCode:    http.StatusBadRequest,
			Args: map[string]interface{}{
				"Body": model.Character{
					Name: "Hallen Weaver",
					Age:  21,
				},
				"ID": primitive.NewObjectID().Hex(),
			},
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			characterHandler := NewCharacterHandler(testCase.CharacterService)

			router := gin.Default()
			router.PUT("/character/:id", characterHandler.UpdateCharacter)

			jsonValue, err := json.Marshal(testCase.Args["Body"])
			assert.Nil(t, err)

			endpoint := fmt.Sprintf("/character/%v", testCase.Args["ID"])
			req := httptest.NewRequest("PUT", endpoint, bytes.NewBuffer(jsonValue))
			tokenString, err := helper.GenerateJWT(primitive.NewObjectID().Hex())
			assert.Nil(t, err)

			if testCase.IsValidAuthentication {
				req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", tokenString))
			} else {
				req.Header.Set("Authorization", "invalid token")
			}
			httpRecorder := httptest.NewRecorder()
			router.ServeHTTP(httpRecorder, req)

			assert.Equal(t, testCase.ExpectedStatusCode, httpRecorder.Code)
		})
	}
}

func TestCharacterHandlerDeleteCharacter(t *testing.T) {
	testCases := map[string]MockCharacterTestCaseWrapper{
		"Normal Path Test": {
			CharacterService:      &MockCharacterService{},
			IsValidTesting:        true,
			IsValidAuthentication: true,
			ExpectedStatusCode:    http.StatusNoContent,
			Args: map[string]interface{}{
				"ID": primitive.NewObjectID().Hex(),
			},
		},
		"Unauthorized User Test": {
			CharacterService:      &MockCharacterService{},
			IsValidTesting:        true,
			IsValidAuthentication: false,
			ExpectedStatusCode:    http.StatusUnauthorized,
			Args: map[string]interface{}{
				"ID": primitive.NewObjectID().Hex(),
			},
		},
		"Invalid ID Test": {
			CharacterService:      &MockCharacterService{},
			IsValidTesting:        true,
			IsValidAuthentication: true,
			ExpectedStatusCode:    http.StatusBadRequest,
			Args: map[string]interface{}{
				"ID": "abc",
			},
		},
		"Character Service Error Test": {
			CharacterService:      &MockCharacterServiceWithErrors{},
			IsValidTesting:        false,
			IsValidAuthentication: true,
			ExpectedStatusCode:    http.StatusBadRequest,
			Args: map[string]interface{}{
				"ID": primitive.NewObjectID().Hex(),
			},
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			characterHandler := NewCharacterHandler(testCase.CharacterService)

			router := gin.Default()
			router.DELETE("/character/:id", characterHandler.DeleteCharacter)

			endpoint := fmt.Sprintf("/character/%v", testCase.Args["ID"])
			req := httptest.NewRequest("DELETE", endpoint, http.NoBody)
			tokenString, err := helper.GenerateJWT(primitive.NewObjectID().Hex())
			assert.Nil(t, err)

			if testCase.IsValidAuthentication {
				req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", tokenString))
			} else {
				req.Header.Set("Authorization", "invalid token")
			}
			httpRecorder := httptest.NewRecorder()
			router.ServeHTTP(httpRecorder, req)

			assert.Equal(t, testCase.ExpectedStatusCode, httpRecorder.Code)
		})
	}
}
