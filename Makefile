compile:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build

bundle:
	zip linux-amd64.zip pimock
