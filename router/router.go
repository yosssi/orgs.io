package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// New creates and returns a router.
func New() http.Handler {
	router := httprouter.New()

	return router
}
