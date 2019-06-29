package vkapi

import (
	"github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
