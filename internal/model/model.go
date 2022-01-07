package model

type (
	// Map used to make a unique structure as unique http responses
	Map map[string]interface{}

	// User contains a user data
	User struct {
		Id   string `json:"id"`
		Name string `json:"name,omitempty"`
	}
)
