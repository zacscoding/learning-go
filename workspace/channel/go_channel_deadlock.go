package channelmulti

func main() {

	ch := make(chan int)
	ch <- 5

	//fatal error: all goroutines are asleep - deadlock!
	//
	//goroutine 1 [chan send]:
	//main.main()
	///home/zaccoding/go/src/github.com/zacscoding/learning-go/workspace/channel/go_channel_deadlock.go:6 +0x50
	//
	//Process finished with exit code 2
}
