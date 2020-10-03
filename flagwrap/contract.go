package flagwrap

type FlagWrap interface {
	GetRootDirectory() *string
	GetResponseFile() *string
	GetPort() *string
}
