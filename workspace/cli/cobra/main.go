package main

import (
	"github.com/zacscoding/learning-go/workspace/cli/cobra/cmd"
	"os"
)

func main() {
	os.Args = append(os.Args,
		"-h")
	cmd.Execute()
}
