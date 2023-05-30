package request

type SearchContactRequest struct {
	Name string `query:"name" validate:"required"`
}
