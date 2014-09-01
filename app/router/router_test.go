package router

import (
	"testing"

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
