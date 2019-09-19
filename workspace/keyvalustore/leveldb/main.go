package main

import (
	"encoding/json"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"os"
	"sync"
)

type Data struct {
	Number int `json:"number"`
}

var dbPath = "./temp-db"
var key = []byte("key")
var useLock = true
var mutex sync.Mutex

func main() {
	err := removeDbPath()
	if err != nil {
		fmt.Println(err)
	}

	db, err := leveldb.OpenFile(dbPath, nil)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	initData := &Data{
		Number: 0,
	}

	serialized, _ := json.Marshal(initData)
	if db.Put(key, serialized, nil) != nil {
		fmt.Println("Failed to put init data")
	}

	var waitGroup sync.WaitGroup
	for i := 0; i < 1000; i++ {
		waitGroup.Add(1)
		go func(db *leveldb.DB) {
			if useLock {
				mutex.Lock()
				defer mutex.Unlock()
			}
			serialized, err := db.Get(key, nil)
			if err != nil {
				fmt.Println("Failed to get data. key :", string(key))
				return
			}

			var data Data
			if json.Unmarshal(serialized, &data) != nil {
				fmt.Println("failed to unmarshal")
			}

			data.Number++
			serialized, _ = json.Marshal(data)
			if db.Put(key, serialized, nil) != nil {
				fmt.Println("Failed to put data")
			}
			defer waitGroup.Done()
		}(db)
	}

	waitGroup.Wait()

	readValue, _ := db.Get(key, nil)
	if json.Unmarshal(readValue, &initData) != nil {
		fmt.Println("failed to unmarshal after done")
	}
	fmt.Println("key1's value is :", initData.Number)
	err = removeDbPath()
	if err != nil {
		fmt.Println(err)
	}
}

func removeDbPath() error {
	_, err := os.Stat(dbPath)
	if err == nil {
		return os.RemoveAll(dbPath)
	}
	return nil
}
