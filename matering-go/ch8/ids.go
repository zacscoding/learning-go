package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	fmt.Println("User id:", os.Getuid())

	var u *user.User
	u, _ = user.Current()
	fmt.Print("Group ids: ")
	groupIDs, _ := u.GroupIds()
	for _, i := range groupIDs {
		fmt.Print(i, " ")
	}
	fmt.Println()
	//User id: 1000
	//Group ids: 1000 4 24 27 30 46 116 126 999
}
