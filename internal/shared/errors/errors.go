package errors

type NotFoundError struct {
	Msg string
}

func (e *NotFoundError) Error() string {
	return e.Msg
}

func NewNotFound(msg string) error {
	return &NotFoundError{Msg: msg}
}

type BadRequestError struct {
	Msg string
}

func (e *BadRequestError) Error() string {
	return e.Msg
}

func NewBadRequest(msg string) error {
	return &BadRequestError{Msg: msg}
}
