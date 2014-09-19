package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/yosssi/orgs.io/app/models"
)

func TestNew(t *testing.T) {
	config := &models.Config{
		App: models.AppConfig{
			Env: models.EnvDevelopment,
		},
		Server: models.ServerConfig{},
	}

	if rtr := New(config); rtr == nil {
		t.Error("rtr should not be nil")
	}
}

func Test_newAceProxy(t *testing.T) {
	config := &models.Config{
		App: models.AppConfig{
			Env: models.EnvDevelopment,
		},
		Server: models.ServerConfig{},
	}

	aceProxy := newAceProxy(config)

	if aceProxy == nil {
		t.Error("aceProxy should not be nil")
	}

	if aceProxy.Opts.DynamicReload != true {
		t.Error("aceProxy.Opts.DynamicReload should not be true")
	}
}

func Test_aceProxy_funcMap_config(t *testing.T) {
	config := &models.Config{
		App: models.AppConfig{
			Env: models.EnvDevelopment,
		},
		Server: models.ServerConfig{},
	}

	aceProxy := newAceProxy(config)

	if aceProxy == nil {
		t.Error("aceProxy should not be nil")
	}

	if c := aceProxy.Opts.FuncMap["config"].(func() *models.Config)(); c != config {
		t.Errorf("c should be %+v [actual: %+v]", config, c)
	}
}

func Test_serveAssets_Err(t *testing.T) {
	r, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Errorf("error occurred [error: %q]", err.Error())
	}

	w := httptest.NewRecorder()

	serveAssets(w, r, []httprouter.Param{{Key: "filepath", Value: "/stylesheets/application.css"}})

	if w.Code != http.StatusInternalServerError {
		t.Errorf("w.Code should be %d [actual: %d]", http.StatusInternalServerError, w.Code)
	}
}

func Test_serveAssets(t *testing.T) {
	r, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Errorf("error occurred [error: %q]", err.Error())
	}

	w := httptest.NewRecorder()

	serveAssets(w, r, []httprouter.Param{{Key: "filepath", Value: "/not/exist/file"}})

	if w.Code != http.StatusNotFound {
		t.Errorf("w.Code should be %d [actual: %d]", http.StatusNotFound, w.Code)
	}
}
