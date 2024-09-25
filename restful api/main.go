package main

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"restful_api/app"
	"restful_api/controller"
	"restful_api/repository"
	"restful_api/service"
)

func main() {
	var db *sql.DB = app.NewDB()
	var validate *validator.Validate = validator.New()
	var categoryRepository repository.CategoryRepository = repository.NewCategoryRepository()
	var categoryService service.CategoryService = service.NewCategoryService(categoryRepository, db, validate)
	var categoryController controller.CategoryController = controller.NewCategoryController(categoryService)

	var router *httprouter.Router = httprouter.New()

	// set router
	router.GET("/api/categories", categoryController.Find)
	router.GET("/api/categories/:id", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:id", categoryController.Update)
	router.DELETE("/api/categories/:id", categoryController.Delete)
}
