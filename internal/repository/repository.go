// Package repository persistence layer.
// Here you will find everything related to data persistence like crud operations and repository conenctions
package repository

import "github.com/yael-castro/godi/internal/model"

// Type defines the available storage types
type Type uint

// Repository types
const (
	// Memory will store the data in memory
	Memory Type = iota
	// Mock defines a mock storage (fake storage)
	Mock
)

// userData pre-load user data used in storage of type Memory
var userData = map[string]model.User{
	"8d6cbb52-6ea5-11ec-90d6-0242ac120003": {
		Id:   "8d6cbb52-6ea5-11ec-90d6-0242ac120003",
		Name: "X",
	},
	"353a7098-aded-4202-b965-5994c01006c3": {
		Id:   "353a7098-aded-4202-b965-5994c01006c3",
		Name: "Y",
	},
}
