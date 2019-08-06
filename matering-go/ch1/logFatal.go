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

	// 정말 나쁜 일이 발생해서 상황을 알려주자마자 프로그램을 종료하고 싶을 때
	log.Fatal(sysLog)
	fmt.Println("Will you see this?")

	//zaccoding@zaccoding-ubuntu:~/go/src/github.com/zacscoding/learning-go/matering-go/ch1$ go run logFatal.go
	//exit status 1

	//zaccoding@zaccoding-ubuntu:~/go/src/github.com/zacscoding/learning-go/matering-go/ch1$ grep "Some program" /var/log/mail.log
	//Aug  6 22:33:37 zaccoding-ubuntu Some program![7063]: 2019/08/06 22:33:37 &{17 Some program! zaccoding-ubuntu   {0 0} 0xc00000e100}
}
