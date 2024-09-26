package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"restful_api/app"
	"restful_api/controller"
	"restful_api/helper"
	"restful_api/middleware"
	"restful_api/model/domain"
	"restful_api/repository"
	"restful_api/service"
	"strconv"
	"strings"
	"testing"
	"time"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:33066)/go_db")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	var validate *validator.Validate = validator.New()
	var categoryRepository repository.CategoryRepository = repository.NewCategoryRepository()
	var categoryService service.CategoryService = service.NewCategoryService(categoryRepository, db, validate)
	var categoryController controller.CategoryController = controller.NewCategoryController(categoryService)
	var router = app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

func truncateCategory(db *sql.DB) {
	db.Exec("TRUNCATE TABLE categories")
}

func TestCreateCategorySuccess(t *testing.T) {
	var db = setupTestDB()
	truncateCategory(db)
	var router http.Handler = setupRouter(db)
	var requestBody = strings.NewReader(`{"name": "test"}`)
	var request *http.Request = httptest.NewRequest("POST", "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("x-api-key", "SECRET")
	var recorder *httptest.ResponseRecorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	var response *http.Response = recorder.Result()
	assert.Equal(t, response.StatusCode, http.StatusOK)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)

	assert.Equal(t, "test", responseBody["data"].(map[string]interface{})["name"])
}

func TestCreateCategoryFail(t *testing.T) {
	var db = setupTestDB()
	truncateCategory(db)
	var router http.Handler = setupRouter(db)
	var requestBody = strings.NewReader(`{"name": ""}`)
	var request *http.Request = httptest.NewRequest("POST", "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("x-api-key", "SECRET")
	var recorder *httptest.ResponseRecorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	var response *http.Response = recorder.Result()
	assert.Equal(t, response.StatusCode, http.StatusBadRequest)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
}

func TestUpdateCategorySuccess(t *testing.T) {
	var db = setupTestDB()
	truncateCategory(db)
	// create data
	tx, _ := db.Begin()
	var categoryRepository repository.CategoryRepository = repository.NewCategoryRepository()
	var category domain.Category = categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "test",
	})
	tx.Commit()
	var router http.Handler = setupRouter(db)
	var requestBody = strings.NewReader(`{"name": "test update"}`)
	var request *http.Request = httptest.NewRequest("PUT", "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), requestBody)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("x-api-key", "SECRET")
	var recorder *httptest.ResponseRecorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	var response *http.Response = recorder.Result()

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "test update", responseBody["data"].(map[string]interface{})["name"])
}

func TestUpdateCategoryFailed(t *testing.T) {
	var db = setupTestDB()
	truncateCategory(db)
	// create data
	tx, _ := db.Begin()
	var categoryRepository repository.CategoryRepository = repository.NewCategoryRepository()
	var category domain.Category = categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "test",
	})
	tx.Commit()
	var router http.Handler = setupRouter(db)
	var requestBody = strings.NewReader(`{"name": ""}`)
	var request *http.Request = httptest.NewRequest("PUT", "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), requestBody)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("x-api-key", "SECRET")
	var recorder *httptest.ResponseRecorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	var response *http.Response = recorder.Result()

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)

	assert.Equal(t, http.StatusBadRequest, int(responseBody["code"].(float64)))
}

func TestGetCategorySuccess(t *testing.T) {
	var db = setupTestDB()
	truncateCategory(db)
	// create data
	tx, _ := db.Begin()
	var categoryRepository repository.CategoryRepository = repository.NewCategoryRepository()
	var category domain.Category = categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "test",
	})
	tx.Commit()
	var router http.Handler = setupRouter(db)
	var request *http.Request = httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), nil)
	request.Header.Add("x-api-key", "SECRET")
	var recorder *httptest.ResponseRecorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	var response *http.Response = recorder.Result()

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, category.Name, responseBody["data"].(map[string]interface{})["name"])
}

func TestGetCategoryFailed(t *testing.T) {
	var db = setupTestDB()
	truncateCategory(db)
	var router http.Handler = setupRouter(db)

	var request *http.Request = httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/404", nil)
	request.Header.Add("x-api-key", "SECRET")
	var recorder *httptest.ResponseRecorder = httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	var response *http.Response = recorder.Result()

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)

	assert.Equal(t, http.StatusNotFound, int(responseBody["code"].(float64)))
}

func TestDeleteCategorySuccess(t *testing.T) {
	var db = setupTestDB()
	truncateCategory(db)
	// create data
	tx, _ := db.Begin()
	var categoryRepository repository.CategoryRepository = repository.NewCategoryRepository()
	var category domain.Category = categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "test",
	})
	tx.Commit()
	var router http.Handler = setupRouter(db)
	var request *http.Request = httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), nil)
	request.Header.Add("x-api-key", "SECRET")
	var recorder *httptest.ResponseRecorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	var response *http.Response = recorder.Result()

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, nil, responseBody["data"])
}

func TestDeleteCategoryFail(t *testing.T) {
	var db = setupTestDB()
	truncateCategory(db)
	var router http.Handler = setupRouter(db)

	var request *http.Request = httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/404", nil)
	request.Header.Add("x-api-key", "SECRET")
	var recorder *httptest.ResponseRecorder = httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	var response *http.Response = recorder.Result()

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)

	assert.Equal(t, http.StatusNotFound, int(responseBody["code"].(float64)))
}

func TestListCategorySuccess(t *testing.T) {
	var db = setupTestDB()
	truncateCategory(db)
	// create data
	tx, _ := db.Begin()
	var categoryRepository repository.CategoryRepository = repository.NewCategoryRepository()
	_ = categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "test",
	})
	tx.Commit()
	var router http.Handler = setupRouter(db)
	var request *http.Request = httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
	request.Header.Add("x-api-key", "SECRET")
	var recorder *httptest.ResponseRecorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	var response *http.Response = recorder.Result()

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)

	var categories = responseBody["data"].([]interface{})
	fmt.Println(categories)
	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Greater(t, len(categories), 0)
}

func TestUnauthorized(t *testing.T) {
	var db = setupTestDB()
	truncateCategory(db)
	var router http.Handler = setupRouter(db)
	var request *http.Request = httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
	request.Header.Add("x-api-key", "WRONG")
	var recorder *httptest.ResponseRecorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	var response *http.Response = recorder.Result()

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusUnauthorized, int(responseBody["code"].(float64)))
}
