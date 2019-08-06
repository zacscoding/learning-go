package main

import (
	"fmt"
	"log"
	"log/syslog"
)

func main() {
	sysLog, err := syslog.New(syslog.LOG_ALERT|syslog.LOG_MAIL, "Some program!")

	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}

	log.Panic(sysLog)
	fmt.Println("Will you see this?")
	//zaccoding@zaccoding-ubuntu:~/go/src/github.com/zacscoding/learning-go/matering-go/ch1$ go run logPanic.go
	//	panic: &{17 Some program! zaccoding-ubuntu   {0 0} 0xc00000e100}
	//
	//goroutine 1 [running]:
	//log.Panic(0xc00005ff78, 0x1, 0x1)
	///usr/local/go/src/log/log.go:333 +0xac
	//main.main()
	///home/zaccoding/go/src/github.com/zacscoding/learning-go/matering-go/ch1/logPanic.go:18 +0xe7
	//exit status 2

}
