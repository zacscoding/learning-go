package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
)

func workByWord(file string) error {
	var err error
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("error reading file %s", err)
			return err
		}
		r := regexp.MustCompile("[^\\s]+")
		words := r.FindAllString(line, -1)
		for _, word := range words {
			fmt.Println(word)
		}
	}
	return nil
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Printf("usage: by_word <file1> [<file2> ...]\n")
		return
	}

	for _, file := range flag.Args() {
		err := workByWord(file)
		if err != nil {
			fmt.Println(err)
		}
	}

	//$ go run by_word.go /var/log/syslog | wc
	//98633   98633  933359
	//$ wc /var/log/syslog
	//7463  98633 946606 /var/log/syslog
}
