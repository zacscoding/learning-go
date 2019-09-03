package main

func main() {
	var c chan string
	close(c)
	// panic: close of nil channel
}
