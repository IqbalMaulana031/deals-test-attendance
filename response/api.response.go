package response

import (
	"encoding/json"
	"net/http"
	"strings"
)

// APIResponseList define interface for common api response
type APIResponseList interface {
	GetCode() int
	GetMessage() string
	GetData() interface{}
}

type apiResponseList struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (a apiResponseList) GetCode() int {
	return a.Code
}

func (a apiResponseList) GetMessage() string {
	return a.Message
}

func (a apiResponseList) GetData() interface{} {
	return a.Data
}

// SuccessAPIResponseList returns formatted success api response
func SuccessAPIResponseList(code int, message string, data interface{}) APIResponseList {
	return &apiResponseList{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// SuccessAPIResponse returns formatted success api response
func SuccessAPIResponse(code int, message string, data interface{}) APIResponseList {
	return &apiResponseList{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// ErrorAPIResponse returns formatted error api response
func ErrorAPIResponse(code int, message string) APIResponseList {
	message = strings.ReplaceAll(message, ";", ":")
	return &apiResponseList{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}

// ErrorAPIResponseWithData returns formatted error api response with data
func ErrorAPIResponseWithData(code int, message string, data interface{}) APIResponseList {
	message = strings.ReplaceAll(message, ";", ":")
	return &apiResponseList{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// ExtractErrorResponse extracts the error response from the message
func ExtractErrorResponse(message string) (int, string, interface{}) {
	const internalServerErrorCode = http.StatusInternalServerError
	var errorResponse struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	if err := json.Unmarshal([]byte(message), &errorResponse); err != nil {
		return internalServerErrorCode, "Internal Server Error", nil
	}
	return errorResponse.Code, errorResponse.Message, errorResponse.Data
}
