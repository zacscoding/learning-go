package main

import (
	"bytes"
	"fmt"
	"github.com/gorilla/rpc/v2/json"
	"log"
	"net/http"
)

func main() {
	args := map[string]string{
		"number":  "010-123-1234",
		"content": "content",
	}

	url := "http://localhost:3000/rpc"
	request, err := json.EncodeClientRequest("sms.Send", args)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(request))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error in sending request to %s. %s", url, err)
	}
	defer resp.Body.Close()

	var result string

	err = json.DecodeClientResponse(resp.Body, &result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(">> RPC Result ", result)
}
