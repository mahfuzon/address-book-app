package request

type AddContactRequest struct {
	Name        string `json:"name" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

func NewAddContactRequest(name, phoneNumber string) AddContactRequest {
	return AddContactRequest{
		Name:        name,
		PhoneNumber: phoneNumber,
	}
}
