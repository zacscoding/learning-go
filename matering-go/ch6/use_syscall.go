package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	pid, _, _ := syscall.Syscall(39, 0, 0, 0)
	fmt.Println("My pid is", pid)
	uid, _, _ := syscall.Syscall(24, 0, 0, 0)
	fmt.Println("User ID:", uid)
	//My pid is 12874
	//User ID: 0

	message := []byte{'H', 'e', 'l', 'l', 'o', '!', '\n'}
	fd := 1
	syscall.Write(fd, message)
	//Hello!

	fmt.Println("Using syscall.Exec()")
	command := "/bin/ls"
	env := os.Environ()
	syscall.Exec(command, []string{"ls", "-a", "-x"}, env)
	//Using syscall.Exec()
	//.  ..  .gitignore  .idea  README.md  ch1  ch2  ch3  ch4  ch5  ch6  go-package.md
}
