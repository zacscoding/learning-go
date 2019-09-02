package main

import (
	"flag"
	"fmt"
	"strings"
)

type Value interface {
	String() string
	Set(string) error
}

type NamesFlag struct {
	Names []string
}

func (s *NamesFlag) GetNames() []string {
	return s.Names
}

func (s *NamesFlag) String() string {
	return fmt.Sprint(s.Names)
}

func (s *NamesFlag) Set(v string) error {
	if len(s.Names) > 0 {
		return fmt.Errorf("cannot use names flag more than once")
	}

	names := strings.Split(v, ",")
	for _, item := range names {
		s.Names = append(s.Names, item)
	}
	return nil
}

func main() {
	var manyNames NamesFlag
	minusK := flag.Int("k", 0, "An int")
	minusO := flag.String("o", "Mihalis", "The name")
	flag.Var(&manyNames, "names", "Comma-separated list")

	flag.Parse()
	fmt.Println("-k:", *minusK)
	fmt.Println("-o:", *minusO)

	for i, item := range manyNames.GetNames() {
		fmt.Println(i, item)
	}

	fmt.Println("Remaining command-line arguments:")
	for index, val := range flag.Args() {
		fmt.Println(index, ":", val)
	}

	//$ go run fun_with_flag.go -names=Mihalis,Jim,Athina 1 two tree
	//-k: 0
	//-o: Mihalis
	//0 Mihalis
	//1 Jim
	//2 Athina
	//Remaining command-line arguments:
	//0 : 1
	//1 : two
	//2 : tree

	//$ go run fun_with_flag.go -Invalid=Matrietta 1 two three
	//flag provided but not defined: -Invalid
	//Usage of /tmp/go-build617106256/b001/exe/fun_with_flag:
	//-k int
	//An int
	//-names value
	//Comma-separated list
	//-o string
	//The name (default "Mihalis")
	//exit status 2
}
