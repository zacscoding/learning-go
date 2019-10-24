package main

import (
	"bytes"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("ssh", "app@192.168.79.130")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
