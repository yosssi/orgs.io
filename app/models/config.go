package models

import (
	"io/ioutil"

	"gopkg.in/yaml.v1"
)

// Env consts
const (
	EnvDevelopment = "development"
	EnvTest        = "test"
	EnvStaging     = "staging"
	EnvProduction  = "production"
)

// Config represents a configuration for Orgs.io.
type Config struct {
	App    AppConfig    `yaml:"app"`
	Server ServerConfig `yaml:"server"`
}

// NewConfig parses the configuration file,
// creates and returns a config.
func NewConfig(flags *Flags) (*Config, error) {
	// Read the configuration file.
	data, err := ioutil.ReadFile(flags.ConfigFilePath)

	if err != nil {
		return nil, err
	}

	// Parse the configuration file.
	config := &Config{}
	if err := yaml.Unmarshal(data, config); err != nil {
		return nil, err
	}

	return config, nil
}

// AppConfig represents a configuration for the application
type AppConfig struct {
	Env string `yaml:"env"`
}

// Development returns true if the Env == envDevelopment.
func (app *AppConfig) Development() bool {
	return app.Env == EnvDevelopment
}

// Test returns true if the Env == envTest.
func (app *AppConfig) Test() bool {
	return app.Env == EnvTest
}

// Staging returns true if the Env == envStaging.
func (app *AppConfig) Staging() bool {
	return app.Env == EnvStaging
}

// Production returns true if the Env == envProduction.
func (app *AppConfig) Production() bool {
	return app.Env == EnvProduction
}

// ServerConfig represents a configuration for the server.
type ServerConfig struct {
	Port string `yaml:"port"`
	CPUs int    `yaml:"cpus"`
}
