package app

import (
	"dependency_injection/controller"
	"dependency_injection/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	var router *httprouter.Router = httprouter.New()

	// set router
	router.GET("/api/categories", categoryController.Find)
	router.GET("/api/categories/:id", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:id", categoryController.Update)
	router.DELETE("/api/categories/:id", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
