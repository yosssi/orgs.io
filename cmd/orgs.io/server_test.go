package main

import (
	"flag"
	"net/http"
	"os"
	"testing"
)

// Configuration file paths
const (
	configFilePathNotExist = "not_exist_file_path"
	configFilePath1        = "test/1.yaml"
)

func init() {
	logPanic = func(v ...interface{}) {}
	listenAndServe = func(addr string, handler http.Handler) error {
		return nil
	}
}

func Test_main_newFlagsErr(t *testing.T) {
	main()
}

func Test_main_newConfigErr(t *testing.T) {
	resetForTesting(nil)

	os.Args = []string{os.Args[0]}
	os.Args = append(os.Args, "-c", configFilePathNotExist)

	main()
}

func Test_main(t *testing.T) {
	resetForTesting(nil)

	os.Args = []string{os.Args[0]}
	os.Args = append(os.Args, "-c", configFilePath1)

	main()
}

// resetForTesting clears all flag state and sets the usage function as directed.
// After calling ResetForTesting, parse errors in flag handling will not
// exit the program.
func resetForTesting(usage func()) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	flag.Usage = usage
}
