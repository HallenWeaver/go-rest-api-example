package handler

import (
	"alexandre/gorest/app/model"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticationHandler(t *testing.T) {
	testUserOK := model.TokenRequest{
		Username: "Andraus23",
		Password: "Subvers10N",
	}

	testCases := map[string]MockUserTestCaseWrapper{
		"Normal Path Test": {
			userService:        &MockUserService{},
			isValidTesting:     true,
			userPayload:        &testUserOK,
			ExpectedStatusCode: http.StatusOK,
		},
		"Invalid Login Payload Test": {
			userService:        &MockUserService{},
			isValidTesting:     true,
			userPayload:        "{",
			ExpectedStatusCode: http.StatusBadRequest,
		},
		"User Service Error Test": {
			userService:        &MockUserServiceWithErrors{},
			isValidTesting:     false,
			userPayload:        &testUserOK,
			ExpectedStatusCode: http.StatusInternalServerError,
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			authenticationHandler := NewAuthenticationHandler(testCase.userService)

			router := gin.Default()
			router.POST("/authenticate", authenticationHandler.LoginUser)

			jsonValue, err := json.Marshal(testCase.userPayload)
			assert.Nil(t, err)

			req, _ := http.NewRequest("POST", "/authenticate", bytes.NewBuffer(jsonValue))
			httpRecorder := httptest.NewRecorder()
			router.ServeHTTP(httpRecorder, req)

			assert.Equal(t, testCase.ExpectedStatusCode, httpRecorder.Code)
		})
	}
}
