package responder

// BaseResponse is a basic json response wrapper
type BaseResponse struct {
	Message    string      `json:"message"`
	Error      string      `json:"error"`
	IsError    bool        `json:"is_error"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
}
