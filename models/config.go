package models

// Env consts
const (
	envDevelopment = "development"
	envTest        = "test"
	envStaging     = "staging"
	envProduction  = "production"
)

// Config represents a configuration for Orgs.io.
type Config struct {
	App    AppConfig    `yaml:"app"`
	Server ServerConfig `yaml:"server"`
}

// AppConfig represents a configuration for the application
type AppConfig struct {
	Env string `yaml:"env"`
}

// Development returns true if the Env == envDevelopment.
func (app *AppConfig) Development() bool {
	return app.Env == envDevelopment
}

// Test returns true if the Env == envTest.
func (app *AppConfig) Test() bool {
	return app.Env == envTest
}

// Staging returns true if the Env == envStaging.
func (app *AppConfig) Staging() bool {
	return app.Env == envStaging
}

// Production returns true if the Env == envProduction.
func (app *AppConfig) Production() bool {
	return app.Env == envProduction
}

// ServerConfig represents a configuration for the server.
type ServerConfig struct {
	Port string `yaml:"port"`
}
