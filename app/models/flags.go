package models

import (
	"errors"
	"flag"
)

// Flag messages
const (
	flagMsgConfigFilePath = "configuration file path"
)

// Errors
var (
	ErrConfigFilePathNotSpecified = errors.New("config file path is not specified")
)

// Flags presents command line flags.
type Flags struct {
	// ConfigFilePath represents a configuration file path.
	ConfigFilePath string
}

// NewFlags parses the command-line flags,
// creates and returns flags.
func NewFlags() (*Flags, error) {
	configFilePath := flag.String("c", "", flagMsgConfigFilePath)

	flag.Parse()

	if *configFilePath == "" {
		flag.PrintDefaults()
		return nil, ErrConfigFilePathNotSpecified
	}

	return &Flags{
		ConfigFilePath: *configFilePath,
	}, nil
}
