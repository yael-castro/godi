package handler

import (
	"net/http"

	"github.com/yael-castro/godi/internal/business"
	"github.com/yael-castro/godi/internal/model"
)

// _ implemetation constraint to User struct
var _ http.Handler = User{}

// User http handler used to handle all requests related to the user
type User struct {
	business.UserProvider
}

// ServeHTTP decides which http.HandlerFunc use based on the http method
func (u User) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		u.provideUser(w, r)

	default:
		MethodNotAllowed(w, r)
	}

}

// provideUser http.HandlerFunc used to handle http requests made with the purpouse to get a user by id
func (u User) provideUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		JSON(w, http.StatusBadRequest, &model.Map{"error": "missing query param 'id'"})
		return
	}

	user, err := u.ProvideUser(id)
	if _, ok := err.(model.NotFound); ok {
		JSON(w, http.StatusNotFound, model.Map{"error": err.Error()})
		return
	}

	if err != nil {
		JSON(w, http.StatusInternalServerError, model.Map{"error": err.Error()})
		return
	}

	JSON(w, http.StatusOK, user)
}
