package request

type RemoveContactRequest struct {
	Slug string `param:"slug" validate:"required"`
}
