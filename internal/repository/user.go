package repository

import (
	"fmt"

	"github.com/yael-castro/godi/internal/model"
)

// UserProvider defines the functionality of a user provider
type UserProvider interface {
	// ProvideUser search a model.User by an string identifier
	ProvideUser(string) (model.User, error)
}

// NewUserProvider abstract factory to UserProvider based on a Type passed as parameter
//
// Returns an error if exists an error with UserProviderImplementation or a invalid Type is passed as parameter
func NewUserProvider(t Type) (UserProvider, error) {
	switch t {
	case Memory:
		return &memoryProvider{
			data: userData,
		}, nil
	case Mock:
		return &mockProvider{}, nil
	}

	return nil, fmt.Errorf(`type "%d" is not supported by repository.NewUserProvider`, t)
}

// NewUProvider build a UserProvider using the constructor NewUserProvider but it no returns an error instead panics
func NewUProvider(t Type) UserProvider {
	provider, err := NewUserProvider(t)
	if err != nil {
		panic(err)
	}

	return provider
}

// memoryProvider implementation of UserProvider that use a memory storage
type memoryProvider struct {
	data map[string]model.User
}

// ProvideUser search a user id in a hash map (data saved in memory)
//
// If the user id no exists in the hash map as key returns a error of type model.NotFound
func (m memoryProvider) ProvideUser(id string) (user model.User, err error) {
	user, ok := m.data[id]
	if !ok {
		err = model.NotFound(fmt.Sprintf(`not found user with id "%s"`, id))
	}

	return
}

// mockProvider mock for UserProvider
type mockProvider struct{}

// ProvideUser always returns a user by the id passed as parameter and never returns errors
func (m mockProvider) ProvideUser(id string) (user model.User, err error) {
	return model.User{
		Id:   id,
		Name: "name",
	}, nil
}
