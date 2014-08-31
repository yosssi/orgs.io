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
