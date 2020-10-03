package flagwrap

import (
	"flag"
	"os"
	"testing"
)

func tearDown() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
}

func TestLoadFlagsGivenAllArgumentsPresentShouldReturnAssignedBaseFlagWrap(t *testing.T) {
	os.Args = []string{"", "-root_directory=root_dir", "-port=3000", "-response_file=response_json"}

	actualBfw, err := Initialize()

	expectedBfw := &BaseFlagWrap{
		rootDirectory: "root_dir",
		port:          "3000",
		responseFile:  "response_json",
	}
	if *actualBfw != *expectedBfw {
		t.Errorf("Initialize() = %v, want %v", *actualBfw, *expectedBfw)
	}
	if err != nil {
		t.Error(err)
	}

	tearDown()
}

func TestLoadFlagsGivenNoArgumentsShouldReturnDefaultBaseFlagWrap(t *testing.T) {
	os.Args = []string{""}

	actualBfw, err := Initialize()

	expectedBfw := &BaseFlagWrap{
		port:          "8080",
		responseFile:  "response",
		rootDirectory: "responses",
	}
	if *actualBfw != *expectedBfw {
		t.Errorf("Initialize() = %v, want %v", *actualBfw, *expectedBfw)
	}
	if err != nil {
		t.Error(err)
	}

	tearDown()
}
