package main

import (
	"github.com/joho/godotenv"
	"github.com/mahfuzon/address_book_app/controllers"
	"github.com/mahfuzon/address_book_app/libraries"
	"github.com/mahfuzon/address_book_app/repositories"
	"github.com/mahfuzon/address_book_app/services"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	db := libraries.SetDb()
	addressBookrRepository := repositories.NewAddressBookRepository(db)
	addressBookrService := services.NewAddressBookService(addressBookrRepository)
	addressBookController := controllers.NewAddressBookController(addressBookrService)
	router := libraries.SetRouter()

	api := router.Group("/api")

	apiV1 := api.Group("/v1")

	apiV1Auth := apiV1.Group("/address-book")
	apiV1Auth.POST("", addressBookController.AddContact)
	apiV1Auth.DELETE("/:slug", addressBookController.RemoveContact)
	apiV1Auth.GET("", addressBookController.SearchContact)

	router.Logger.Fatal(router.Start(":8000"))
}
