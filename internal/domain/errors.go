package domain

type NotFoundError struct {
	Message string
}

type InvalidLengthError struct {
	Message string
}

func (e *InvalidLengthError) Error() string {
	return e.Message
}

func (e *NotFoundError) Error() string {
	return e.Message
}
