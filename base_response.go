package freeforex

// BaseResponse is base structure for all responses.
type BaseResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
