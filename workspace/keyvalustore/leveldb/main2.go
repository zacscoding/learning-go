package main

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"log"
	"os"
)

var path = "./temp-db"

func main() {
	// setup leveldb & key values
	err := clearLocalDatabase()
	if err != nil {
		log.Println(err)
	}

	db, err := leveldb.OpenFile(path, nil)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	kv := make(map[string]string)
	kv["p1.key1"] = "value11"
	kv["p1.key2"] = "value12"
	kv["p1.key3"] = "value13"
	kv["p2.key1"] = "value21"
	kv["p2.key2"] = "value22"
	kv["p2.key3"] = "value23"

	// insert all
	for k, v := range kv {
		bk := []byte(k)
		bv := []byte(v)
		err := db.Put(bk, bv, nil)
		if err != nil {
			log.Fatal("Failed to save ", k)
		}
	}

	// search all
	count := 0
	itr1 := db.NewIterator(new(util.Range), nil)
	for itr1.Next() {
		count++
	}
	itr1.Release()
	fmt.Println("db.NewIterator(new(util.Range), nil) :", count)

	// search with start
	count = 0
	itr2 := db.NewIterator(&util.Range{Start: []byte("p2.key1"), Limit: []byte("p2.key3")}, nil)
	for itr2.Next() {
		count++
	}
	itr2.Release()
	fmt.Println("db.NewIterator(&util.Range{Start: []byte(\"p2.key1\"), Limit: []byte(\"p2.key3\")}, nil) :", count)

	// search with byte prefix
	count = 0
	itr3 := db.NewIterator(util.BytesPrefix([]byte("p1")), nil)
	for itr3.Next() {
		count++
	}
	itr3.Release()
	fmt.Println("db.NewIterator(util.BytesPrefix([]byte(\"p1\")), nil) :", count)

	// test update
	// kv["p1.key1"] = "value11"
	err = db.Put([]byte("p1.key1"), []byte("updated"), nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	read, err := db.Get([]byte("p1.key1"), nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("After put with same key >> ", string(read))

	// clear
	err = clearLocalDatabase()
	if err != nil {
		log.Println(err)
	}

	//outputs
	//db.NewIterator(new(util.Range), nil) : 6
	//db.NewIterator(&util.Range{Start: []byte("p2.key1"), Limit: []byte("p2.key3")}, nil) : 2
	//db.NewIterator(util.BytesPrefix([]byte("p1")), nil) : 3
}

func clearLocalDatabase() error {
	_, err := os.Stat(path)
	if err == nil {
		return os.RemoveAll(path)
	}
	return nil
}
