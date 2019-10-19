package responder

// BaseResponse is a basic json response wrapper
type BaseResponse struct {
	Message string      `json:"message"`
	IsError bool        `json:"is_error"`
	Data    interface{} `json:"data"`
}
