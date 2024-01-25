package errors

type ErrorCode int

const (
	InvalidDate ErrorCode = iota
	Unauthenticated
	InvalidRoom
	RoomAlreadyBooked
)

type CustomError struct {
	Code    ErrorCode
	Message string
}

func (e CustomError) Error() string {
	return e.Message
}
