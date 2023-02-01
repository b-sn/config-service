package custom

type DuplicateError struct {
	Text string
}

func (e DuplicateError) Error() string {
	return e.Text
}

type NotFoundError struct {
	Text string
}

func (e NotFoundError) Error() string {
	return e.Text
}

type RequestError struct {
	Text string
}

func (e RequestError) Error() string {
	return e.Text
}
