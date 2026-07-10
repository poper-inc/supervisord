// +build !release
//go:generate go run github.com/UnnoTed/fileb0x@v1.1.4 b0x.yaml

package main

import (
	"net/http"
)

//HTTP auto generated
var HTTP http.FileSystem = http.Dir("./webgui")
