//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"belajar-golang-dasar/app"
	"belajar-golang-dasar/controller"
	"belajar-golang-dasar/middleware"
	"belajar-golang-dasar/repository"
	"belajar-golang-dasar/service"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl))

	service.NewCategoryService,
	wire.Bind(new(service.CategoryRepository), new(*service.CategoryRepositoryImpl))

	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl))
)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router))
		middleware.NewAuthMiddleware,
		NewServer,
	)

	return nil
}