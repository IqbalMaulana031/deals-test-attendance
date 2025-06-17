package response

type APIResponseList interface {
	GetRequestID() string
	GetCode() int
	GetMessage() string
	GetData() interface{}
}

type apiResponseList struct {
	RequestID string      `json:"request_id"`
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

func (a apiResponseList) GetRequestID() string {
	return a.RequestID
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

func SuccessAPIResponseList(code int, requestID, message string, data interface{}) APIResponseList {
	return &apiResponseList{
		RequestID: requestID,
		Code:      code,
		Message:   message,
		Data:      data,
	}
}

func ErrorAPIResponse(code int, requestID, message string) APIResponseList {
	return &apiResponseList{
		RequestID: requestID,
		Code:      code,
		Message:   message,
		Data:      nil,
	}
}

type APIResponseWithoutReqID interface {
	GetCode() int
	GetMessage() string
	GetData() []string
}

type apiResponseWithoutReqID struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (a apiResponseWithoutReqID) GetCode() int {
	return a.Code
}

func (a apiResponseWithoutReqID) GetMessage() string {
	return a.Message
}

func (a apiResponseWithoutReqID) GetData() interface{} {
	return a.Data
}

func SuccessAPIResponseWithoutReqID(code int, message string, data interface{}) apiResponseWithoutReqID {
	return apiResponseWithoutReqID{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func ErrorAPIResponseWithoutReqID(code int, message string) apiResponseWithoutReqID {
	return apiResponseWithoutReqID{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}

// Tambahkan interface baru untuk response dengan array of strings
type APIResponseArray interface {
	GetRequestID() string
	GetCode() int
	GetMessage() string
	GetData() []string
}

type apiResponseArray struct {
	RequestID string   `json:"request_id"`
	Code      int      `json:"code"`
	Message   string   `json:"message"`
	Data      []string `json:"data"`
}

func (a apiResponseArray) GetRequestID() string {
	return a.RequestID
}

func (a apiResponseArray) GetCode() int {
	return a.Code
}

func (a apiResponseArray) GetMessage() string {
	return a.Message
}

func (a apiResponseArray) GetData() []string {
	return a.Data
}

// Fungsi baru untuk mengembalikan array of strings sebagai data
func ErrorAPIArrayResponse(code int, requestID, message string, data []string) APIResponseArray {
	return &apiResponseArray{
		RequestID: requestID,
		Code:      code,
		Message:   message,
		Data:      data,
	}
}

func SetRequestID(requestId string) func(*apiResponseList) {
	return func(b *apiResponseList) {
		b.RequestID = requestId
	}
}

func SetRequestIDArray(requestId string) func(*apiResponseArray) {
	return func(b *apiResponseArray) {
		b.RequestID = requestId
	}
}
