// Package dependency manages
package dependency

import (
	"fmt"

	"github.com/yael-castro/godi/internal/business"
	"github.com/yael-castro/godi/internal/handler"
	"github.com/yael-castro/godi/internal/repository"
)

// Profile defines options of dependency injection
type Profile uint

// Supported profiles for dependency injection
const (
	// Default defines the production profile
	Default Profile = iota
	// Testing defines the testing profile used to make a unit and integration tests
	Testing
)

// Injector defines a dependency injector
type Injector interface {
	// Inject takes any data type and fill of required dependencies (dependency injection)
	Inject(interface{}) error
}

// InjectorFunc function that implements the Injector interface
type InjectorFunc func(interface{}) error

func (f InjectorFunc) Inject(i interface{}) error {
	return f(i)
}

// NewInjector is an abstract factory to Injector, it builds a instance of Injector interface based on the Profile based as parameter
//
// Supported profiles: Default and Testing
//
// If pass a parameter an invalid profile it panics
func NewInjector(p Profile) Injector {
	switch p {
	case Default:
		return InjectorFunc(routerDefault)

	case Testing:
		return InjectorFunc(routerTesting)
	}

	panic(fmt.Sprintf(`invalid profile: "%d" is not supported`, p))
}

// routerDefault InjectorFunc for *handler.Handler that uses a Default Profile
func routerDefault(i interface{}) error {
	h, ok := i.(*handler.Handler)
	if !ok {
		return fmt.Errorf(`required "%T" not "%T"`, h, i)
	}

	userProvider, err := repository.NewUserProvider(repository.Memory)
	if err != nil {
		return err
	}

	h.User = handler.User{
		UserProvider: business.AccountProvider{
			UserProvider: userProvider,
		},
	}

	return nil
}

// routerTesting InjectorFunc for *handler.Handler that uses a Testing Profile
func routerTesting(i interface{}) error {
	h, ok := i.(*handler.Handler)
	if !ok {
		return fmt.Errorf(`required "%T" not "%T"`, h, i)
	}

	userProvider, err := repository.NewUserProvider(repository.Mock)
	if err != nil {
		return err
	}

	h.User = handler.User{
		UserProvider: business.AccountProvider{
			UserProvider: userProvider,
		},
	}

	return nil
}
