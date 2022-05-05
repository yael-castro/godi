package model

// ValidationError is returned when occurs a validation error
type ValidationError string

// Error returns the string value of NotFound
func (e ValidationError) Error() string {
	return string(e)
}

// NotFound is returned when occurs an error related to missing resource
type NotFound string

// Error returns the string value of NotFound
func (n NotFound) Error() string {
	return string(n)
}
