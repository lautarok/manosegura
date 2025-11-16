package dto

type AppError struct {
	StatusCode int    `json:"statusCode" example:"500" validate:"required"`
	Message    string `json:"message" example:"internal server error" validate:"required"`
}
