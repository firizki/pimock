package flagwrap

func GetDummyBaseFlagWrap() *BaseFlagWrap {
	return &BaseFlagWrap{}
}

func GetSampleBaseFlagWrap() *BaseFlagWrap {
	return &BaseFlagWrap{
		port:          "8080",
		responseFile:  "response",
		rootDirectory: "responses",
	}
}
