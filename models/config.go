package models

// Config represents a configuration for Orgs.io.
type Config struct {
	ServerConfig ServerConfig `yaml:"server"`
}

// ServerConfig represents a configuration for a server.
type ServerConfig struct {
	Port string `yaml:"port"`
}
