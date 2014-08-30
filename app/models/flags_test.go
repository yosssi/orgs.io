package models

import (
	"flag"
	"os"
	"testing"
)

var configFilePath = "test.yaml"

func TestNewFlags_configFilePathNotSpecified(t *testing.T) {
	resetForTesting(nil)

	os.Args = []string{os.Args[0]}

	flags, err := NewFlags()

	if flags != nil {
		t.Error("flags should be nil [actual: %v]", flags)
	}

	if err != ErrConfigFilePathNotSpecified {
		t.Error("err should be %v [actual: %v]", ErrConfigFilePathNotSpecified, err)
	}
}

func TestNewFlags_noErrorOccurs(t *testing.T) {
	resetForTesting(nil)

	os.Args = []string{os.Args[0]}
	os.Args = append(os.Args, "-c", configFilePath)

	flags, err := NewFlags()

	if flags == nil {
		t.Error("flags should not be nil")
	}

	if flags.ConfigFilePath != configFilePath {
		t.Error("flags.ConfigFilePath should be %v [actual: %v]", configFilePath, flags.ConfigFilePath)
	}

	if err != nil {
		t.Error("err should be %v [actual: %v]", ErrConfigFilePathNotSpecified, err)
	}
}

// resetForTesting clears all flag state and sets the usage function as directed.
// After calling ResetForTesting, parse errors in flag handling will not
// exit the program.
func resetForTesting(usage func()) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	flag.Usage = usage
}
