package utils

const (
	// StatusOk : All good
	StatusOk = "ok"
	// StatusNOk : Bad response
	StatusNOk = "nok"
)

// CustomError : The error format in api response
type CustomError struct {
	Code    int      `json:"code"`
	Details []string `json:"details"`
}

// Response : The api response format
type Response struct {
	Status string       `json:"status"`
	Error  *CustomError `json:"error,omitempty"`
	Result *interface{} `json:"result,omitempty"`
}

// Send : General function to send api response
func Send(payload interface{}) *Response {
	return &Response{
		Status: StatusOk,
		Result: &payload,
	}
}

// Fail : General function to send api response
func Fail(code int, details ...string) *Response {
	return &Response{
		Status: StatusNOk,
		Error: &CustomError{
			Code:    code,
			Details: details,
		},
	}
}
