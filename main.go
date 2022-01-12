package main

import (
	"github.com/pro-bitcoin/cirrus-reporter/cmd"
)

var version = "1.0.0"

func main() {
	if err := cmd.Execute(version); err != nil {
		panic(err)
	}
}
