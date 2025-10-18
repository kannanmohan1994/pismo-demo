package utils

type Meta struct {
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

// Response : The api response format
type Response struct {
	Data any   `json:"data,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
}

// Send : General function to send api response
func Send(payload any, err error, errCode string) *Response {
	var meta *Meta
	if err != nil {
		if len(errCode) == 0 {
			errCode = errCodeMap[err]
		} else if code, ok := errCodeMap[err]; ok {
			errCode = code
		} else {
			errCode = "UNKNOWN_ERROR"
		}
		meta = &Meta{
			Error:   errCode,
			Message: err.Error(),
		}
	}

	return &Response{
		Data: payload,
		Meta: meta,
	}
}
