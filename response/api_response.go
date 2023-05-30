package response

type ApiResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewApiResponse(status, message string, data interface{}) ApiResponse {
	return ApiResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
}
