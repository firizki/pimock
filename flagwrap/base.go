package flagwrap

import (
	"flag"
	"os"
)

type BaseFlagWrap struct {
	rootDirectory string
	responseFile  string
	port          string
}

func Initialize() (*BaseFlagWrap, error) {
	flags, err := loadFlags(os.Args[1:])
	if err != nil {
		return nil, err
	}

	return &BaseFlagWrap{
		rootDirectory: *flags[RootDirectoryFlagName],
		port:          *flags[PortFlagName],
		responseFile:  *flags[ResponseFileFlagName],
	}, nil
}

func loadFlags(args []string) (map[string]*string, error) {
	flags := make(map[string]*string)
	flags[RootDirectoryFlagName] = loadRootDirectory()
	flags[PortFlagName] = loadPort()
	flags[ResponseFileFlagName] = loadResponseFile()

	err := flag.CommandLine.Parse(args)
	if err != nil {
		return nil, err
	}

	return flags, nil
}

func (self *BaseFlagWrap) GetPort() *string {
	return &self.port
}

func (self *BaseFlagWrap) GetResponseFile() *string {
	return &self.responseFile
}

func (self *BaseFlagWrap) GetRootDirectory() *string {
	return &self.rootDirectory
}
