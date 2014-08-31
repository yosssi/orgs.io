package main

import (
	"log"
	"net/http"

	"github.com/yosssi/orgs.io/app/models"
	"github.com/yosssi/orgs.io/app/router"
)

var logPanic = log.Panic

func main() {
	// Parse the command-line flags.
	flags, err := models.NewFlags()
	if err != nil {
		logPanic(err)
		return
	}

	// Read and parse the configuration file.
	configc, errc := models.NewConfig(flags)
	var config *models.Config
	select {
	case config = <-configc:
	case err := <-errc:
		logPanic(err)
		return
	}

	logPanic(http.ListenAndServe(":"+config.Server.Port, router.New(config)))
}
