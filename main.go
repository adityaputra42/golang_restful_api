package main

import (
	"golang_restful_api/helper"
	"golang_restful_api/middleware"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}
func NewValidator() *validator.Validate {
	return validator.New()
}
func main() {
	// db := app.NewDB()
	// validate := validator.New()

	// categoryRepository := repository.NewCategoryRepository()
	// categoryService := service.NewCategoryService(categoryRepository, db, validate)
	// categoryController := controller.NewCategoryController(categoryService)
	// router := app.NewRouter(categoryController)
	// authMiddlewar := middleware.NewAuthMiddleware(router)
	// server := NewServer(authMiddlewar)
	// err := server.ListenAndServe()
	// helper.PanicIfError(err)

	// menggunakan dependency injection
	server := InitializedServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
