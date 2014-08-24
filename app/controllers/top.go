package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/yosssi/ace-proxy"
	"github.com/yosssi/orgs.io/app/models"
)

// Top represents a top controller.
type Top struct {
	Config   *models.Config
	AceProxy *proxy.Proxy
}

// Index represents an index action.
func (ctrl *Top) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tpl, err := ctrl.AceProxy.Load("layout/base", "top/index", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// NewTop creates and returns a top controller.
func NewTop(config *models.Config, aceProxy *proxy.Proxy) *Top {
	return &Top{
		Config:   config,
		AceProxy: aceProxy,
	}
}
