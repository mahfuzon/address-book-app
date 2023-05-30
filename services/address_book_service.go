package services

import (
	"errors"
	"github.com/gosimple/slug"
	"github.com/mahfuzon/address_book_app/models"
	"github.com/mahfuzon/address_book_app/repositories"
	"github.com/mahfuzon/address_book_app/request"
	"github.com/mahfuzon/address_book_app/response"
)

type AddressBookService interface {
	AddContact(request request.AddContactRequest) (response.ContactResponse, error)
	RemoveContact(request request.RemoveContactRequest) error
	SearchContact(request request.SearchContactRequest) (response.ContactResponse, error)
}

type addressBookService struct {
	addressBookRepository repositories.AddressBookRepository
}

func NewAddressBookService(addressBookRepository repositories.AddressBookRepository) AddressBookService {
	return &addressBookService{addressBookRepository: addressBookRepository}
}

func (addressBookService *addressBookService) AddContact(request request.AddContactRequest) (response.ContactResponse, error) {

	// check if name or phone number already exists
	contact, err := addressBookService.addressBookRepository.FindByPhoneNumberOrName(request.PhoneNumber, request.Name)
	if err != nil && err.Error() != "record not found" {
		return response.ContactResponse{}, err
	}
	if contact.Id > 0 {
		return response.ContactResponse{}, errors.New("name or phone number already exists")
	}

	contact = models.Contact{
		Name:        request.Name,
		PhoneNumber: request.PhoneNumber,
		Slug:        slug.Make(request.Name),
	}

	contact, err = addressBookService.addressBookRepository.Create(contact)
	if err != nil {
		return response.ContactResponse{}, err
	}

	contactResponse := response.ConverseToContactResponse(contact)
	return contactResponse, nil
}

func (addressBookService *addressBookService) RemoveContact(request request.RemoveContactRequest) error {
	contact, err := addressBookService.addressBookRepository.FindBySlug(request.Slug)
	if err != nil {
		return err
	}

	err = addressBookService.addressBookRepository.Delete(contact)
	if err != nil {
		return err
	}

	return nil
}

func (addressBookService *addressBookService) SearchContact(request request.SearchContactRequest) (response.ContactResponse, error) {
	contact, err := addressBookService.addressBookRepository.FindByName(request.Name)
	if err != nil {
		return response.ContactResponse{}, err
	}

	contactResponse := response.ConverseToContactResponse(contact)
	return contactResponse, nil

}
