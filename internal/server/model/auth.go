package model

type LoginRequest struct {
	BaseRequest
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	BaseResponse
	Success bool   `json:"success"`
	Token   string `json:"token,omitempty"`
}
