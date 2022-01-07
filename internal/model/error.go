package model

// ValidationError returns when occurs a validation error
type ValidationError string

func (e ValidationError) Error() string {
	return string(e)
}

// NotFound error implementation
//
// Returns when occurs an error related to missing resource
type NotFound string

func (e NotFound) Error() string {
	return string(e)
}
