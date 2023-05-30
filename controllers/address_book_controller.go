package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/mahfuzon/address_book_app/helper"
	"github.com/mahfuzon/address_book_app/request"
	"github.com/mahfuzon/address_book_app/response"
	"github.com/mahfuzon/address_book_app/services"
)

type AddressBookController struct {
	AddressBookService services.AddressBookService
}

func NewAddressBookController(addressBookService services.AddressBookService) *AddressBookController {
	return &AddressBookController{AddressBookService: addressBookService}
}

func (addressBookController *AddressBookController) AddContact(ctx echo.Context) error {
	request := request.AddContactRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		apiResponse := response.NewApiResponse("error", "failed add contact", err.Error())
		return ctx.JSON(500, apiResponse)
	}

	err = ctx.Validate(&request)
	if err != nil {
		errorMessage := helper.ConverseToErrorString(err.(validator.ValidationErrors))
		if err != nil {
			apiResponse := response.NewApiResponse("error", "failed add contact", errorMessage)
			return ctx.JSON(422, apiResponse)
		}
	}

	responseService, err := addressBookController.AddressBookService.AddContact(request)
	if err != nil {
		apiResponse := response.NewApiResponse("error", "failed add contact", err.Error())
		return ctx.JSON(400, apiResponse)
	}

	apiResponse := response.NewApiResponse("ok", "success add contact", responseService)
	return ctx.JSON(200, apiResponse)
}

func (addressBookController *AddressBookController) RemoveContact(ctx echo.Context) error {
	request := request.RemoveContactRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		apiResponse := response.NewApiResponse("error", "failed remove contact", err.Error())
		return ctx.JSON(500, apiResponse)
	}

	err = ctx.Validate(&request)
	if err != nil {
		errorMessage := helper.ConverseToErrorString(err.(validator.ValidationErrors))
		if err != nil {
			apiResponse := response.NewApiResponse("error", "failed remove contact", errorMessage)
			return ctx.JSON(422, apiResponse)
		}
	}

	err = addressBookController.AddressBookService.RemoveContact(request)
	if err != nil {
		apiResponse := response.NewApiResponse("error", "failed remove contact", err.Error())
		return ctx.JSON(400, apiResponse)
	}

	apiResponse := response.NewApiResponse("ok", "success remove contact", nil)
	return ctx.JSON(200, apiResponse)
}

func (addressBookController *AddressBookController) SearchContact(ctx echo.Context) error {
	request := request.SearchContactRequest{}
	err := ctx.Bind(&request)
	if err != nil {
		apiResponse := response.NewApiResponse("error", "failed search contact", err.Error())
		return ctx.JSON(500, apiResponse)
	}

	err = ctx.Validate(&request)
	if err != nil {
		errorMessage := helper.ConverseToErrorString(err.(validator.ValidationErrors))
		if err != nil {
			apiResponse := response.NewApiResponse("error", "failed search contact", errorMessage)
			return ctx.JSON(422, apiResponse)
		}
	}

	contact, err := addressBookController.AddressBookService.SearchContact(request)
	if err != nil {
		apiResponse := response.NewApiResponse("error", "contact tidak ditemukan", err.Error())
		return ctx.JSON(400, apiResponse)
	}

	apiResponse := response.NewApiResponse("ok", "success search contact", contact)
	return ctx.JSON(200, apiResponse)
}
