package errors

type ErrMissingField string

func (e ErrMissingField) Error() string {
	return string(e) + " is required"
}
