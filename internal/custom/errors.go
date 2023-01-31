package custom

// type CustomError struct {
// 	Text string
// }

// func (e CustomError) Error() string {
// 	return e.Text
// }

// type DuplicateError CustomError
// type NotFoundError CustomError
// type NotFoundError CustomError

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
