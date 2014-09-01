package router

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/yosssi/ace"
	"github.com/yosssi/ace-proxy"
	"github.com/yosssi/orgs.io/app/controllers"
	"github.com/yosssi/orgs.io/app/models"
)

const aceBaseDir = "app/views"

// New creates and returns a router.
func New(config *models.Config) http.Handler {
	// Create an Ace proxy.
	aceProxy := newAceProxy(config)

	// Create controllers.
	top := controllers.NewTop(config, aceProxy)

	// Create a router.
	router := httprouter.New()

	router.GET("/", top.Index)

	if config.App.Development() {
		router.ServeFiles("/bower_components/*filepath", http.Dir("bower_components"))
		router.ServeFiles("/assets/*filepath", http.Dir("app/assets"))
	}

	return router
}

// newAceProxy creates and returns an Ace proxy.
func newAceProxy(config *models.Config) *proxy.Proxy {
	return proxy.New(&ace.Options{
		BaseDir:       aceBaseDir,
		DynamicReload: config.App.Development(),
		FuncMap: template.FuncMap{
			"config": func() *models.Config {
				return config
			},
		},
	})
}
