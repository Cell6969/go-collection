package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"restful_api/helper"
	"restful_api/model/web"
	"restful_api/service"
	"strconv"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var categoryRequest = web.CategoryCreateRequest{}
	helper.ReadFromRequest(request, &categoryRequest)
	var result web.CategoryResponse = controller.CategoryService.Create(request.Context(), categoryRequest)
	var webResponse web.WebResponse = web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   result,
	}
	helper.WriteToResponse(writer, webResponse)
}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var categoryRequest = web.CategoryUpdateRequest{}
	helper.ReadFromRequest(request, &categoryRequest)

	var categoryId string = params.ByName("id")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)
	categoryRequest.Id = id

	var result web.CategoryResponse = controller.CategoryService.Update(request.Context(), categoryRequest)
	var webResponse web.WebResponse = web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   result,
	}

	helper.WriteToResponse(writer, webResponse)
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var categoryId string = params.ByName("id")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.CategoryService.Delete(request.Context(), id)
	var webResponse web.WebResponse = web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	helper.WriteToResponse(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var categoryId string = params.ByName("id")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	var categoryResponse = controller.CategoryService.FindById(request.Context(), id)
	var webResponse web.WebResponse = web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponse(writer, webResponse)
}

func (controller *CategoryControllerImpl) Find(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var categoryResponses = controller.CategoryService.FindAll(request.Context())
	var webResponse web.WebResponse = web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponses,
	}
	helper.WriteToResponse(writer, webResponse)
}
