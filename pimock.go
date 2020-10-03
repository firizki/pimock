package main

import (
	"net/http"

	fw "github.com/firizki/pimock/flagwrap"
)

var mapDiscover *Discover
var flagWrap *fw.BaseFlagWrap

func main() {
	flagWrap, err := fw.Initialize()
	if err != nil {
		panic(err)
	}

	mapDiscover = NewDiscover(*flagWrap.GetRootDirectory(), *flagWrap.GetResponseFile())

	http.HandleFunc("/", MockServer(mapDiscover, flagWrap))
	http.ListenAndServe(":"+*flagWrap.GetPort(), nil)
}
