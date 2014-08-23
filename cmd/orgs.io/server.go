package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/yosssi/orgs.io/models"
	"github.com/yosssi/orgs.io/router"
	"gopkg.in/yaml.v1"
)

func main() {
	configPath := flag.String("c", "", "config file path")

	flag.Parse()

	if *configPath == "" {
		flag.PrintDefaults()
		panic("config file path is not specified")
	}

	data, err := ioutil.ReadFile(*configPath)
	if err != nil {
		panic(err)
	}

	config := &models.Config{}

	if err := yaml.Unmarshal(data, config); err != nil {
		panic(err)
	}

	log.Panic(http.ListenAndServe(":"+config.Server.Port, router.New(config)))
}
