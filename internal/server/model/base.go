package model

type BaseRequest struct {
	Index   int `json:"index"`
	Counter int `json:"counter"`
}

type BaseResponse struct {
	Code    int    `json:"CODE"`
	Message string `json:"message,omitempty"`
}
