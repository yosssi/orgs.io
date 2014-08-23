package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/yosssi/orgs.io/router"
)

func main() {
	configPath := flag.String("c", "", "config file path")

	flag.Parse()

	if *configPath == "" {
		flag.PrintDefaults()
		panic("config file path is not specified")
	}

	log.Panic(http.ListenAndServe(":8080", router.New()))
}
