package test

import (
	"fmt"
	"golang_restful_api/simple"
	"testing"
)

func TestSimpleService(t *testing.T) {

	simpleService := simple.InitializeSimple()
	fmt.Println(simpleService.SimpleRepository)

}
