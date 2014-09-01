package main

import (
	"log"
	"net/http"
	"runtime"

	"github.com/yosssi/orgs.io/app/models"
	"github.com/yosssi/orgs.io/app/router"
)

var logPanic = log.Panic

var listenAndServe = http.ListenAndServe

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

	// Set the maximum number of CPUs.
	setCPUs(config.Server.CPUs)

	logPanic(listenAndServe(":"+config.Server.Port, router.New(config)))
}

// setCPUs sets the maximum number of CPUs.
func setCPUs(cpus int) int {
	localCPUs := runtime.NumCPU()

	switch {
	case cpus < 1:
		cpus = 1
	case localCPUs < cpus:
		cpus = localCPUs
	}

	return runtime.GOMAXPROCS(cpus)
}
