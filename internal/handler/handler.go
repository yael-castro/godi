// Package handler handles all http requests made to the app (is the presentation layer)
package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"path"

	"github.com/yael-castro/godi/internal/model"
)

// NotFound handles not found requests
func NotFound(w http.ResponseWriter, r *http.Request) {
	JSON(w, http.StatusNotFound, model.Map{"message": "not found"})
}

// MethodNotAllowed handles invalid requests using an illegal method
func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write(nil)
}

// Healthcheck handles requests made to check the status server
func Healthcheck(w http.ResponseWriter, r *http.Request) {
	JSON(w, http.StatusOK, model.Map{"message": "ok"})
}

// JSON function used to send serialized data in json more easier
func JSON(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}

// New constructs an empty Handler
func New() *Handler {
	return &Handler{}
}

// Handler main handler used in the ListeAndServe
type Handler struct {
	User http.Handler
}

// ServeHTTP decides which http.HandlerFunc use based on the http method
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	log.Println(path.Join(r.URL.Path, "/"))

	switch p := path.Join(r.URL.Path, "/"); p {

	case "/godi/v1/user":
		h.User.ServeHTTP(w, r)

	case "/godi/v1/healthcheck":
		Healthcheck(w, r)

	default:
		NotFound(w, r)
	}
}
