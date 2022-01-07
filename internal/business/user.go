// Package business orchestrate business logic
package business

import (
	"strings"

	"github.com/yael-castro/godi/internal/model"
	"github.com/yael-castro/godi/internal/repository"
)

// UserProvider defines the business that provides information about users
type UserProvider interface {
	repository.UserProvider
}

// AccountProvider implements UserProvider
type AccountProvider struct {
	repository.UserProvider
}

// ProvideUser use repository.UserProvider to search and find a user by id but first trim the spaces in the id
func (p AccountProvider) ProvideUser(id string) (model.User, error) {
	return p.UserProvider.ProvideUser(strings.Trim(id, " "))
}
