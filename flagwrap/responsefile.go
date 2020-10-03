package flagwrap

import "flag"

const (
	ResponseFileFlagName     = "response_file"
	ResponseFileDefaultValue = "response"
	ResponseFileHelpText     = "Set response filename"
)

func loadResponseFile() *string {
	return flag.String(ResponseFileFlagName, ResponseFileDefaultValue, ResponseFileHelpText)
}
