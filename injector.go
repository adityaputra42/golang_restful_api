//go:build wireinject
// +build wireinject

package main

import (
	"golang_restful_api/app"
	"golang_restful_api/controller"
	"golang_restful_api/middleware"
	"golang_restful_api/repository"
	"golang_restful_api/service"
	"net/http"

	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var CategorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		NewValidator,
		CategorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}
