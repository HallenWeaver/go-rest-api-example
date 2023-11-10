package service

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ErrorMessageTesting = fmt.Errorf("test error")

func AssertValidityCondition(t *testing.T, result interface{}, err error, isValidTesting bool) {
	if isValidTesting {
		assert.NotNil(t, result)
		assert.Nil(t, err)
	} else {
		assert.Nil(t, result)
		assert.NotNil(t, err.Error())
	}
}

func AssertValidityConditionBoolean(t *testing.T, result bool, err error, isValidTesting bool) {
	if isValidTesting {
		assert.True(t, result)
		assert.Nil(t, err)
	} else {
		assert.False(t, result)
		assert.NotNil(t, err.Error())
	}
}
