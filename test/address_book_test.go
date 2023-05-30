package test

import (
	"encoding/json"
	"fmt"
	"github.com/gosimple/slug"
	"github.com/labstack/echo/v4"
	"github.com/mahfuzon/address_book_app/controllers"
	"github.com/mahfuzon/address_book_app/libraries"
	"github.com/mahfuzon/address_book_app/models"
	"github.com/mahfuzon/address_book_app/repositories"
	"github.com/mahfuzon/address_book_app/response"
	"github.com/mahfuzon/address_book_app/services"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TruncateTableContacts(db *gorm.DB) {
	db.Exec("TRUNCATE TABLE CONTACTS")
}

func TestAddContactSuccess(t *testing.T) {
	// instansiasi db
	db := libraries.SetDbTest()

	//truncate table contacts
	TruncateTableContacts(db)

	// setup controller
	addressBookRepository := repositories.NewAddressBookRepository(db)
	addressBooksService := services.NewAddressBookService(addressBookRepository)
	addressBookController := controllers.NewAddressBookController(addressBooksService)

	// make example request
	requestJsonString := `{
	"name" :"mahfuzon akhiar",
"phone_number" : "081278160990"
}`

	// setup router
	router := libraries.SetRouter()
	router.POST("api/v1/address_book", addressBookController.AddContact)

	// make request test
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/address_book", strings.NewReader(requestJsonString))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// get result recorder
	result := rec.Result()
	assert.Equal(t, 200, result.StatusCode)

	// get result body
	body := result.Body

	// get response body
	responseBody, _ := io.ReadAll(body)
	var apiResponse response.ApiResponse
	err := json.Unmarshal(responseBody, &apiResponse)
	assert.NoError(t, err)

	fmt.Println(apiResponse.Data)
}

func TestAddContactValidationError(t *testing.T) {
	// instansiasi db
	db := libraries.SetDbTest()

	//truncate table contacts
	TruncateTableContacts(db)

	// setup controller
	addressBookRepository := repositories.NewAddressBookRepository(db)
	addressBooksService := services.NewAddressBookService(addressBookRepository)
	addressBookController := controllers.NewAddressBookController(addressBooksService)

	// make example request
	requestJsonString := `{
	
}`

	// setup router
	router := libraries.SetRouter()
	router.POST("api/v1/address_book", addressBookController.AddContact)

	// make request test
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/address_book", strings.NewReader(requestJsonString))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// get result recorder
	result := rec.Result()
	assert.Equal(t, 422, result.StatusCode)

	// get result body
	body := result.Body

	// get response body
	responseBody, _ := io.ReadAll(body)
	var apiResponse response.ApiResponse
	err := json.Unmarshal(responseBody, &apiResponse)
	assert.NoError(t, err)

	fmt.Println(apiResponse.Data)
}

func TestAddContactIfNameOrPhoneNumberAlreadyExists(t *testing.T) {
	// instansiasi db
	db := libraries.SetDbTest()

	//truncate table contacts
	TruncateTableContacts(db)

	// create dummy contacts
	contact := models.Contact{
		Name:        "mahfuzon akhiar",
		PhoneNumber: "081278160991",
		Slug:        slug.Make("mahfuzon akhiar"),
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
	err := db.Create(&contact).Error
	assert.NoError(t, err)

	// setup controller
	addressBookRepository := repositories.NewAddressBookRepository(db)
	addressBooksService := services.NewAddressBookService(addressBookRepository)
	addressBookController := controllers.NewAddressBookController(addressBooksService)

	// make example request
	requestJsonString := `{
	"name" :"mahfuzon akhiar",
"phone_number" : "081278160990"
}`

	// setup router
	router := libraries.SetRouter()
	router.POST("api/v1/address_book", addressBookController.AddContact)

	// make request test
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/address_book", strings.NewReader(requestJsonString))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// get result recorder
	result := rec.Result()
	assert.Equal(t, 400, result.StatusCode)

	// get result body
	body := result.Body

	// get response body
	responseBody, _ := io.ReadAll(body)
	var apiResponse response.ApiResponse
	err = json.Unmarshal(responseBody, &apiResponse)
	assert.NoError(t, err)

	fmt.Println(apiResponse.Data)
}

func TestRemoveContactSuccess(t *testing.T) {
	// instansiasi db
	db := libraries.SetDbTest()

	//truncate table contacts
	TruncateTableContacts(db)

	// create dummy contacts
	contact := models.Contact{
		Name:        "mahfuzon akhiar",
		PhoneNumber: "081278160991",
		Slug:        slug.Make("mahfuzon akhiar"),
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
	err := db.Create(&contact).Error
	assert.NoError(t, err)

	// setup controller
	addressBookRepository := repositories.NewAddressBookRepository(db)
	addressBooksService := services.NewAddressBookService(addressBookRepository)
	addressBookController := controllers.NewAddressBookController(addressBooksService)

	// setup router
	router := libraries.SetRouter()
	router.DELETE("api/v1/address_book/:slug", addressBookController.RemoveContact)

	// make request test
	req := httptest.NewRequest(http.MethodDelete, "http://localhost:8000/api/v1/address_book/mahfuzon-akhiar", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// get result recorder
	result := rec.Result()
	assert.Equal(t, 200, result.StatusCode)

	// get result body
	body := result.Body

	// get response body
	responseBody, _ := io.ReadAll(body)
	var apiResponse response.ApiResponse
	err = json.Unmarshal(responseBody, &apiResponse)
	assert.NoError(t, err)

	fmt.Println(apiResponse.Data)
}

func TestRemoveContactSlugNotExists(t *testing.T) {
	// instansiasi db
	db := libraries.SetDbTest()

	//truncate table contacts
	TruncateTableContacts(db)

	// setup controller
	addressBookRepository := repositories.NewAddressBookRepository(db)
	addressBooksService := services.NewAddressBookService(addressBookRepository)
	addressBookController := controllers.NewAddressBookController(addressBooksService)

	// setup router
	router := libraries.SetRouter()
	router.DELETE("api/v1/address_book/:slug", addressBookController.RemoveContact)

	// make request test
	req := httptest.NewRequest(http.MethodDelete, "http://localhost:8000/api/v1/address_book/mahfuzon-akhiar", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// get result recorder
	result := rec.Result()
	assert.Equal(t, 400, result.StatusCode)

	// get result body
	body := result.Body

	// get response body
	responseBody, _ := io.ReadAll(body)
	var apiResponse response.ApiResponse
	err := json.Unmarshal(responseBody, &apiResponse)
	assert.NoError(t, err)

	fmt.Println(apiResponse.Data)
}

func TestSearchContactSuccess(t *testing.T) {
	// instansiasi db
	db := libraries.SetDbTest()

	//truncate table contacts
	TruncateTableContacts(db)

	// create dummy contacts
	contact := models.Contact{
		Name:        "mahfuzon akhiar",
		PhoneNumber: "081278160991",
		Slug:        slug.Make("mahfuzon akhiar"),
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
	err := db.Create(&contact).Error
	assert.NoError(t, err)

	// setup controller
	addressBookRepository := repositories.NewAddressBookRepository(db)
	addressBooksService := services.NewAddressBookService(addressBookRepository)
	addressBookController := controllers.NewAddressBookController(addressBooksService)

	// setup router
	router := libraries.SetRouter()
	router.GET("api/v1/address_book", addressBookController.SearchContact)

	// make request test
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8000/api/v1/address_book?name=mahfuzon+akhiar", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// get result recorder
	result := rec.Result()
	assert.Equal(t, 200, result.StatusCode)

	// get result body
	body := result.Body

	// get response body
	responseBody, _ := io.ReadAll(body)
	var apiResponse response.ApiResponse
	err = json.Unmarshal(responseBody, &apiResponse)
	assert.NoError(t, err)

	fmt.Println(apiResponse.Data)
}
