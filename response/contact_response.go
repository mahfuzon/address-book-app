package response

import "github.com/mahfuzon/address_book_app/models"

type ContactResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Slug        string `json:"slug"`
}

func ConverseToContactResponse(contact models.Contact) ContactResponse {
	return ContactResponse{
		Id:          contact.Id,
		Name:        contact.Name,
		PhoneNumber: contact.PhoneNumber,
		Slug:        contact.Slug,
	}
}
