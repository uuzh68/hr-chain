package main

import (
	"github.com/uuzh68/hr-chain/info"
)

var (
	Version = "0.3.0"
	Build   string
	Tag     string
)

func init() {
	info.App.Version = Version
	info.App.Build = Build
	info.App.Tag = Tag
}
