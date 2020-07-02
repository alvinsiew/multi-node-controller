package main

import (
	"fmt"
	"multi-node-controller/internal/options"
	"os"
)

func main() {
	// awsinternal.FilterInstances()
	// sshcmd.RemoteCommand()
	if err := options.Root(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
