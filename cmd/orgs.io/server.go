package main

import (
	"log"
	"net/http"

	"github.com/yosssi/orgs.io/app/models"
	"github.com/yosssi/orgs.io/app/router"
)

func main() {
	// Parse the command-line flags.
	flags, err := models.NewFlags()
	if err != nil {
		panic(err)
	}

	// Read and parse the configuration file.
	configc, errc := models.NewConfig(flags)
	var config *models.Config
	select {
	case config = <-configc:
	case err := <-errc:
		panic(err)
	}

	log.Panic(http.ListenAndServe(":"+config.Server.Port, router.New(config)))
}
