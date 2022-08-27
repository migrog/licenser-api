package dto

import (
	"fmt"
	"net/http"
)

type HttpResponse struct {
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
}

func NewHttpResponse(statusCode int, data interface{}) *HttpResponse {
	return &HttpResponse{StatusCode: statusCode, Data: data}
}

func FortmatError(message interface{}) map[string]string {
	mapError := make(map[string]string)
	mapError["error"] = fmt.Sprintf("%v", message)

	return mapError
}

func Ok(data interface{}) *HttpResponse {
	return NewHttpResponse(http.StatusOK, data)
}

func Created(data interface{}) *HttpResponse {
	return NewHttpResponse(http.StatusCreated, data)
}

func NoContent() *HttpResponse {
	return NewHttpResponse(http.StatusNoContent, nil)
}

func BadRequest(data interface{}) *HttpResponse {
	return NewHttpResponse(http.StatusBadRequest, FortmatError(data))
}

func Unauthorized() *HttpResponse {
	return NewHttpResponse(http.StatusUnauthorized, FortmatError("Invalid or expired token"))
}

func NotFound() *HttpResponse {
	return NewHttpResponse(http.StatusNotFound, FortmatError("Resource not found"))
}

func ServerError() *HttpResponse {
	return NewHttpResponse(http.StatusInternalServerError, FortmatError("An unexpected error occurred"))
}
