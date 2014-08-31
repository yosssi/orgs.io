package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/yosssi/ace"
	"github.com/yosssi/ace-proxy"
	"github.com/yosssi/orgs.io/app/models"
)

func TestTop_Index_loadErr(t *testing.T) {
	config := &models.Config{
		App: models.AppConfig{
			Env: "test",
		},
		Server: models.ServerConfig{
			Port: "8080",
		},
	}

	aceProxy := proxy.New(nil)

	top := NewTop(config, aceProxy)

	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf("err %q occurred", err.Error())
	}

	w := httptest.NewRecorder()

	top.Index(w, r, nil)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("w.Code should be %d [actual: %d]", http.StatusInternalServerError, w.Code)
	}

	expectedErrMsg := "open layout/base.ace: no such file or directory\n"
	if w.Body == nil || w.Body.String() != expectedErrMsg {
		t.Errorf("w.Body should be %q [actual: %q]", expectedErrMsg, w.Body)
	}
}

func TestTop_Index_executeErr(t *testing.T) {
	config := &models.Config{
		App: models.AppConfig{
			Env: "test",
		},
		Server: models.ServerConfig{
			Port: "8080",
		},
	}

	aceProxy := proxy.New(&ace.Options{
		BaseDir:       "test/top_1",
		DynamicReload: true,
	})

	top := NewTop(config, aceProxy)

	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf("err %q occurred", err.Error())
	}

	w := httptest.NewRecorder()

	top.Index(w, r, nil)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("w.Code should be %d [actual: %d]", http.StatusInternalServerError, w.Code)
	}

	expectedErrMsg := `template: layout/base:top/index:1:17: executing "T1" at <.A>: can't evaluate field A in type string` + "\n"
	if w.Body == nil || w.Body.String() != expectedErrMsg {
		t.Errorf("w.Body should be %q [actual: %q]", expectedErrMsg, w.Body)
	}
}

func TestTop_Index(t *testing.T) {
	config := &models.Config{
		App: models.AppConfig{
			Env: "test",
		},
		Server: models.ServerConfig{
			Port: "8080",
		},
	}

	aceProxy := proxy.New(&ace.Options{
		BaseDir:       "test/top_2",
		DynamicReload: true,
	})

	top := NewTop(config, aceProxy)

	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf("err %q occurred", err.Error())
	}

	w := httptest.NewRecorder()

	top.Index(w, r, nil)

	if w.Code != http.StatusOK {
		t.Errorf("w.Code should be %d [actual: %d]", http.StatusOK, w.Code)
	}
}

func TestNewTop(t *testing.T) {
	config := &models.Config{
		App: models.AppConfig{
			Env: "test",
		},
		Server: models.ServerConfig{
			Port: "8080",
		},
	}

	aceProxy := proxy.New(nil)

	top := NewTop(config, aceProxy)

	if top.Config != config {
		t.Errorf("top.Config shoud be %+v [actual: %+v]", config, top.Config)
	}

	if top.AceProxy != aceProxy {
		t.Errorf("top.AceProxy shoud be %+v [actual: %+v]", aceProxy, top.AceProxy)
	}
}
