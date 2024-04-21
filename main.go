package main

import (
	"kubing/pkg/cmd"
	// get component-base for logging setup
	"k8s.io/component-base/cli"
	"k8s.io/kubectl/pkg/cmd/util"
)

func main() {
	command := cmd.NewKubingCommand()
	if err := cli.RunNoErrOutput(command); err != nil {
		// Pretty-print the error and exit with an error.
		util.CheckErr(err)
	}
}
