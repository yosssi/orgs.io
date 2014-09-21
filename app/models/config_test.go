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
	NewConfig(&Flags{})
}

func TestNewConfig_readFileErr(t *testing.T) {
	flags := &Flags{
		ConfigFilePath: configFilePathNotExist,
	}

	_, err := NewConfig(flags)

	if err == nil {
		t.Error("error should occur")
	}

	if expectedErrMsg := fmt.Sprintf("open %s: no such file or directory", flags.ConfigFilePath); err.Error() != expectedErrMsg {
		t.Errorf("err should be %q [actual: %q]", expectedErrMsg, err.Error())
	}
}

func Test_newConfig_yamlUnmarshalErr(t *testing.T) {
	flags := &Flags{
		ConfigFilePath: configFilePath1,
	}

	_, err := NewConfig(flags)

	if err == nil {
		t.Error("error should occur")
	}

	if expectedErrMsg := "YAML error: control characters are not allowed"; err.Error() != expectedErrMsg {
		t.Errorf("err should be %q [actual: %q]", expectedErrMsg, err.Error())
	}
}

func Test_newConfig(t *testing.T) {
	flags := &Flags{
		ConfigFilePath: configFilePath2,
	}

	config, err := NewConfig(flags)

	if err != nil {
		t.Errorf("err should not occur [error: %q]", err.Error())
	}

	expectedEnv := "development"
	if config.App.Env != expectedEnv {
		t.Errorf("config.App.Env should be %q [actual %q]", expectedEnv, config.App.Env)
	}

	expectedPort := "8080"
	if config.Server.Port != expectedPort {
		t.Errorf("config.Server.Port should be %q [actual %q]", expectedEnv, config.App.Env)
	}
}

func TestAppConfig_Development(t *testing.T) {
	appConfig := &AppConfig{
		Env: EnvDevelopment,
	}

	if !appConfig.Development() {
		t.Error("appConfig.Development() should be true")
	}

	appConfig.Env = EnvTest

	if appConfig.Development() {
		t.Error("appConfig.Development() should be false")
	}
}

func TestAppConfig_Test(t *testing.T) {
	appConfig := &AppConfig{
		Env: EnvTest,
	}

	if !appConfig.Test() {
		t.Error("appConfig.Test() should be true")
	}

	appConfig.Env = EnvDevelopment

	if appConfig.Test() {
		t.Error("appConfig.Test() should be false")
	}
}

func TestAppConfig_Staging(t *testing.T) {
	appConfig := &AppConfig{
		Env: EnvStaging,
	}

	if !appConfig.Staging() {
		t.Error("appConfig.Staging() should be true")
	}

	appConfig.Env = EnvDevelopment

	if appConfig.Staging() {
		t.Error("appConfig.Staging() should be false")
	}
}

func TestAppConfig_Production(t *testing.T) {
	appConfig := &AppConfig{
		Env: EnvProduction,
	}

	if !appConfig.Production() {
		t.Error("appConfig.Production() should be true")
	}

	appConfig.Env = EnvDevelopment

	if appConfig.Production() {
		t.Error("appConfig.Production() should be false")
	}
}
