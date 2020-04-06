package main

import (
	"github.com/zacscoding/learning-go/fullstack/chap06/backend/src/rest"
	"log"
)

func main() {
	log.Println("Main log...")
	log.Fatal(rest.RunAPI("127.0.0.1:8000"))
}
