package test

import (
	"golang_restful_api/simple"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleServiceError(t *testing.T) {

	simpleService, err := simple.InitializeSimple(true)
	assert.NotNil(t, err)
	assert.Nil(t, simpleService)

}

func TestSimpleServiceSucces(t *testing.T) {

	simpleService, err := simple.InitializeSimple(false)
	assert.NotNil(t, simpleService)
	assert.Nil(t, err)

}
