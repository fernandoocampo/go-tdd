package web

import "net/http"

// RestHandler Defines behavior for any kind of handler in a REST mode.
type RestHandler interface {
	// Create creates a new record
	Create(w http.ResponseWriter, r *http.Request)
}
