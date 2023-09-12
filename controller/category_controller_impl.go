package controller

import (
	"encoding/json"
	"golang_restful_api/helper"
	"golang_restful_api/model/web"
	"golang_restful_api/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

// Create implements CategoryController.
func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)

	categoryCreateRequest := web.CategoryCreateRequest{}
	err := decoder.Decode(&categoryCreateRequest)
	helper.PanicIfError(err)

	categoryResponses := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponses,
	}
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	errorEncode := encoder.Encode(webResponse)
	helper.PanicIfError(errorEncode)

}

// Delete implements CategoryController.
func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("unimplemented")
}

// FindAll implements CategoryController.
func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("unimplemented")
}

// FindById implements CategoryController.
func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("unimplemented")
}

// Update implements CategoryController.
func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("unimplemented")
}

func NewCategoryController() CategoryController {
	return &CategoryControllerImpl{}
}
