package models

import (
	"fmt"
	"testing"
)

// Configuration file paths
const (
	configFilePathNotExist = "not_exist_file_path"
	configFilePath1        = "test/1.yaml"
	configFilePath2        = "test/2.yaml"
)

func TestNewConfig(t *testing.T) {
	configc, errc := NewConfig(&Flags{})
	select {
	case <-configc:
	case <-errc:
	}
}

func Test_newConfig_readFileErr(t *testing.T) {
	flags := &Flags{
		ConfigFilePath: configFilePathNotExist,
	}

	configc := make(chan *Config)
	errc := make(chan error)

	go newConfig(flags, configc, errc)

	select {
	case <-configc:
		t.Error("an error should be returned")
	case err := <-errc:
		expectedErrMsg := fmt.Sprintf("open %s: no such file or directory", flags.ConfigFilePath)
		if err.Error() != expectedErrMsg {
			t.Errorf("err should be %q [actual: %q]", expectedErrMsg, err.Error())
		}
	}
}

func Test_newConfig_yamlUnmarshalErr(t *testing.T) {
	flags := &Flags{
		ConfigFilePath: configFilePath1,
	}

	configc := make(chan *Config)
	errc := make(chan error)

	go newConfig(flags, configc, errc)

	select {
	case <-configc:
		t.Error("an error should be returned")
	case err := <-errc:
		expectedErrMsg := "YAML error: control characters are not allowed"
		if err.Error() != expectedErrMsg {
			t.Errorf("err should be %q [actual: %q]", expectedErrMsg, err.Error())
		}
	}
}

func Test_newConfig(t *testing.T) {
	flags := &Flags{
		ConfigFilePath: configFilePath2,
	}

	configc := make(chan *Config)
	errc := make(chan error)

	go newConfig(flags, configc, errc)

	select {
	case config := <-configc:
		expectedEnv := "development"
		if config.App.Env != expectedEnv {
			t.Errorf("config.App.Env should be %q [actual %q]", expectedEnv, config.App.Env)
		}

		expectedPort := "8080"
		if config.Server.Port != expectedPort {
			t.Errorf("config.Server.Port should be %q [actual %q]", expectedEnv, config.App.Env)
		}
	case err := <-errc:
		t.Errorf("err should not be returned [actual: %v]", err)
	}

}

func TestAppConfig_Development(t *testing.T) {
	appConfig := &AppConfig{
		Env: envDevelopment,
	}

	if !appConfig.Development() {
		t.Error("appConfig.Development() should be true")
	}

	appConfig.Env = envTest

	if appConfig.Development() {
		t.Error("appConfig.Development() should be false")
	}
}

func TestAppConfig_Test(t *testing.T) {
	appConfig := &AppConfig{
		Env: envTest,
	}

	if !appConfig.Test() {
		t.Error("appConfig.Test() should be true")
	}

	appConfig.Env = envDevelopment

	if appConfig.Test() {
		t.Error("appConfig.Test() should be false")
	}
}

func TestAppConfig_Staging(t *testing.T) {
	appConfig := &AppConfig{
		Env: envStaging,
	}

	if !appConfig.Staging() {
		t.Error("appConfig.Staging() should be true")
	}

	appConfig.Env = envDevelopment

	if appConfig.Staging() {
		t.Error("appConfig.Staging() should be false")
	}
}

func TestAppConfig_Production(t *testing.T) {
	appConfig := &AppConfig{
		Env: envProduction,
	}

	if !appConfig.Production() {
		t.Error("appConfig.Production() should be true")
	}

	appConfig.Env = envDevelopment

	if appConfig.Production() {
		t.Error("appConfig.Production() should be false")
	}
}
