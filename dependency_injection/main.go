package main

import (
	"database/sql"
	"dependency_injection/app"
	"dependency_injection/controller"
	"dependency_injection/helper"
	"dependency_injection/middleware"
	"dependency_injection/repository"
	"dependency_injection/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	var db *sql.DB = app.NewDB()
	var validate *validator.Validate = validator.New()
	var categoryRepository repository.CategoryRepository = repository.NewCategoryRepository()
	var categoryService service.CategoryService = service.NewCategoryService(categoryRepository, db, validate)
	var categoryController controller.CategoryController = controller.NewCategoryController(categoryService)
	var router = app.NewRouter(categoryController)

	var server http.Server = http.Server{
		Addr:    ":3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	var err error = server.ListenAndServe()
	helper.PanicIfError(err)
}
