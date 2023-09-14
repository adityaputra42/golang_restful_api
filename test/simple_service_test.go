package test

import (
	"fmt"
	"golang_restful_api/simple"
	"testing"
)

func TestSimpleService(t *testing.T) {

	simpleService, err := simple.InitializeSimple()
	fmt.Println(err)
	fmt.Println(simpleService)

}
