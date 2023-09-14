package model

type ErrorCode int

const (
	INVALID_SLOT_ID ErrorCode = iota + 2000
	ERR_DATABASE
	ERR_INVALID_REQUEST
	ERR_NOT_EXIST_SLOT
	ERR_NO_FREE_SPIN_COUNT
	ERR_NOT_SUPPORT_DEBUG_FUNCTION
	ERR_NOT_ENOUGH_CASH
	ERR_NEED_CLAIM
	ERR_NO_NEED_CLAIM
	ERR_UNKNOWN
)

type Error struct {
	Code    ErrorCode
	Message string
}

func (e Error) Error() string {
	return e.Message
}

func (e Error) GetCode() int {
	return int(e.Code)
}
