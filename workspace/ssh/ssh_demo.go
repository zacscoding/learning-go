package main

import (
	"bufio"
	"fmt"
	"github.com/shiena/ansicolor"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
	"os/signal"
)

// https://gist.github.com/Mebus/c3a437e339481de03a98569090c53b08
func main() {
	host := "192.168.79.130:22"
	user := "app"
	/*pemPath := "/home/zaccoding/keys/local-vm.pem"
	pemBytes, err := ioutil.ReadFile(pemPath)
	if err != nil {
		log.Fatal(err)
		return
	}
	key, err := ssh.ParsePrivateKey(pemBytes)
	auth := ssh.PublicKeys(key)*/
	auth := ssh.Password("apppw")

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			auth,
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", host, config)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	defer session.Close()

	session.Stdout = ansicolor.NewAnsiColorWriter(os.Stdout)
	session.Stderr = ansicolor.NewAnsiColorWriter(os.Stderr)
	in, _ := session.StdinPipe()

	modes := ssh.TerminalModes{
		ssh.ECHO:  0, // Disable echoing
		ssh.IGNCR: 1, // Ignore CR on input.
	}

	if err := session.RequestPty("vt100", 80, 40, modes); err != nil {
		log.Fatalf("request for pseudo terminal failed: %s", err)
	}

	// Start remote shell
	if err := session.Shell(); err != nil {
		log.Fatalf("failed to start shell: %s", err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for {
			<-c
			os.Exit(0)
		}
	}()

	// Accepting commands
	for {
		reader := bufio.NewReader(os.Stdin)
		str, _ := reader.ReadString('\n')
		fmt.Fprint(in, str)

	}
}
