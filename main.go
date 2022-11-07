/*
Copyright © 2022 James Best <jim@justjibba.net>
*/
package main

import (
	"github.com/jim-at-jibba/dev-tools-cli/cmd"
	"golang.design/x/clipboard"
)

func main() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
	cmd.Execute()
}
