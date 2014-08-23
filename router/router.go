package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/yosssi/orgs.io/controllers"
)

// New creates and returns a router.
func New() http.Handler {
	router := httprouter.New()

	router.GET("/", controllers.TopIndex)

	return router
}
