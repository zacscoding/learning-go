package channelmulti

import (
	"fmt"
	"time"
)

func startProducer(messageChannel chan string) {
	for i := 0; i < 5; i++ {
		message := fmt.Sprint("message", i)
		messageChannel <- message
		time.Sleep(2 * time.Second)
	}
	close(messageChannel)
}

func startConsumer(id int, messageChannel chan string, completeChannel chan bool) {
	for {
		read, ok := <-messageChannel

		if !ok {
			break
		}

		fmt.Printf("Read data from %d : %s\n", id, read)
	}
	completeChannel <- true
}

func main() {
	messageChannel := make(chan string)
	complete := make(chan bool)
	go startConsumer(1, messageChannel, complete)
	//go startConsumer(2, messageChannel)
	go startProducer(messageChannel)

	<-complete
	fmt.Println("Terminate main")
}
