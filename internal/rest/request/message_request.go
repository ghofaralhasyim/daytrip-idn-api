package rest_request

type MessageRequest struct {
	Name        string `json:"name" validate:"required"`
	Phone       string `json:"phone" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	PackageName string `json:"packageName" validate:"required"`
	Message     string `json:"message" validate:"required"`
}
