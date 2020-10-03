package flagwrap

import "flag"

const (
	RootDirectoryFlagName     = "root_directory"
	RootDirectoryDefaultValue = "responses"
	RootDirectoryHelpText     = "Set root directory for response files."
)

func loadRootDirectory() *string {
	return flag.String(RootDirectoryFlagName, RootDirectoryDefaultValue, RootDirectoryHelpText)
}
