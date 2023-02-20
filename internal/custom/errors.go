package custom

type (
	DuplicateError struct {
		Err error
	}
	NotFoundError struct {
		Err error
	}
	RequestError struct {
		Err error
	}
	UnauthorizedError struct {
		Err error
	}
)

func NewDuplicateError(err error) DuplicateError {
	return DuplicateError{Err: err}
}
func NewNotFoundError(err error) NotFoundError {
	return NotFoundError{Err: err}
}
func NewRequestError(err error) RequestError {
	return RequestError{Err: err}
}
func NewUnauthorizedError(err error) UnauthorizedError {
	return UnauthorizedError{Err: err}
}

func (e DuplicateError) Error() string {
	return e.Err.Error()
}

func (e NotFoundError) Error() string {
	return e.Err.Error()
}

func (e RequestError) Error() string {
	return e.Err.Error()
}

func (e UnauthorizedError) Error() string {
	return e.Err.Error()
}
