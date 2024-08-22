package model

type ResponseError struct {
	Message string `json:"message"`
}

type ResponseSuccess struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponseError(message string) *ResponseError {
	return &ResponseError{
		Message: message,
	}
}
