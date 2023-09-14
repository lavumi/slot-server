package model

import "slot-server/internal/slot/model"

type SpinRequest struct {
	BaseRequest
	model.SpinInput
}

type SpinResponse struct {
	BaseResponse
	model.SpinOutput
}
