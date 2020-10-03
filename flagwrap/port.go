package flagwrap

import "flag"

const (
	PortFlagName     = "port"
	PortDefaultValue = "8080"
	PortHelpText     = "Set server port"
)

func loadPort() *string {
	return flag.String(PortFlagName, PortDefaultValue, PortHelpText)
}
