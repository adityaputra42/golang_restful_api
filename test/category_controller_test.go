package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"golang_restful_api/app"
	"golang_restful_api/controller"
	"golang_restful_api/helper"
	"golang_restful_api/middleware"
	"golang_restful_api/model/domain"
	"golang_restful_api/repository"
	"golang_restful_api/service"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"

	"github.com/go-playground/validator/v10"
)

func SetupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golang_restful_api_test")
	helper.PanicIfError(err)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
}

func SetupRouter(db *sql.DB) http.Handler {

	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)
	return middleware.NewAuthMiddleware(router)
}

func TruncateCategory(db *sql.DB) {
	db.Exec("TRUNCATE category")

}

func TestCreateCategorySucces(t *testing.T) {
	db := SetupTestDB()
	TruncateCategory(db)
	router := SetupRouter(db)

	requestBody := strings.NewReader(`{"name":"Laptop"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	result := recorder.Result()
	assert.Equal(t, 200, result.StatusCode)

	body, _ := io.ReadAll(result.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Ok", responseBody["status"])
	assert.Equal(t, "Laptop", responseBody["data"].(map[string]interface{})["name"])

}

func TestCreateCategoryFailed(t *testing.T) {
	db := SetupTestDB()
	TruncateCategory(db)
	router := SetupRouter(db)

	requestBody := strings.NewReader(`{"name":""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	result := recorder.Result()
	assert.Equal(t, 400, result.StatusCode)

	body, _ := io.ReadAll(result.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Bad Request", responseBody["status"])

}
func TestUpdateCategorySucces(t *testing.T) {
	db := SetupTestDB()
	TruncateCategory(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{Name: "Komputer"})
	tx.Commit()

	router := SetupRouter(db)

	requestBody := strings.NewReader(`{"name":"Kasur"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	result := recorder.Result()
	assert.Equal(t, 200, result.StatusCode)

	body, _ := io.ReadAll(result.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Ok", responseBody["status"])
	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "Kasur", responseBody["data"].(map[string]interface{})["name"])

}
func TestUpdateCategoryFailed(t *testing.T) {
	db := SetupTestDB()
	TruncateCategory(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{Name: "Komputer"})
	tx.Commit()

	router := SetupRouter(db)

	requestBody := strings.NewReader(`{"name":""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	result := recorder.Result()
	assert.Equal(t, 400, result.StatusCode)

	body, _ := io.ReadAll(result.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Bad Request", responseBody["status"])

}
func TestGetCategorySucces(t *testing.T) {
	db := SetupTestDB()
	TruncateCategory(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{Name: "Komputer"})

	tx.Commit()

	router := SetupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), nil)

	request.Header.Add("X-API-Key", "RAHASIA")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	result := recorder.Result()
	assert.Equal(t, 200, result.StatusCode)

	body, _ := io.ReadAll(result.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Ok", responseBody["status"])
	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, category.Name, responseBody["data"].(map[string]interface{})["name"])

}
func TestGetCategoryFailed(t *testing.T) {
	db := SetupTestDB()
	TruncateCategory(db)
	router := SetupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/404", nil)
	request.Header.Add("X-API-Key", "RAHASIA")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	result := recorder.Result()
	assert.Equal(t, 404, result.StatusCode)

	body, _ := io.ReadAll(result.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "Not Found", responseBody["status"])

}
func TestDeleteCategorySucces(t *testing.T) {
	db := SetupTestDB()
	TruncateCategory(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{Name: "Komputer"})

	tx.Commit()

	router := SetupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	result := recorder.Result()
	assert.Equal(t, 200, result.StatusCode)

	body, _ := io.ReadAll(result.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Ok", responseBody["status"])

}
func TestDeleteCategoryFailed(t *testing.T) {
	db := SetupTestDB()
	TruncateCategory(db)
	router := SetupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/404", nil)
	request.Header.Add("X-API-Key", "RAHASIA")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	result := recorder.Result()
	assert.Equal(t, 404, result.StatusCode)

	body, _ := io.ReadAll(result.Body)
	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)
	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "Not Found", responseBody["status"])

}

func TestGetListCategoriesSucces(t *testing.T) {
	db := SetupTestDB()
	TruncateCategory(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{Name: "Komputer"})

	tx.Commit()

	router := SetupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
	request.Header.Add("X-API-Key", "RAHASIA")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	result := recorder.Result()
	assert.Equal(t, 200, result.StatusCode)

	body, _ := io.ReadAll(result.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Ok", responseBody["status"])
	categories := responseBody["data"].([]interface{})
	categoryResponse1 := categories[0].(map[string]interface{})
	assert.Equal(t, category.Id, int(categoryResponse1["id"].(float64)))
	assert.Equal(t, category.Name, categoryResponse1["name"])

}
func TestUnauhtorized(t *testing.T) {
	db := SetupTestDB()
	TruncateCategory(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{Name: "Komputer"})

	tx.Commit()

	router := SetupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), nil)

	request.Header.Add("X-API-Key", "Karepmu")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	result := recorder.Result()
	assert.Equal(t, 401, result.StatusCode)

	body, _ := io.ReadAll(result.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, "Unauthorized", responseBody["status"])

}
