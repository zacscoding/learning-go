package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"regexp"
)

func findIP(input string) string {
	partIP := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	grammar := partIP + "\\." + partIP + "\\." + partIP + "\\." + partIP
	matchMe := regexp.MustCompile(grammar)
	return matchMe.FindString(input)
}

func main() {
	// go run find_ipv4.go auth.log | sort -rn | uniq -c | sort -rn
	// sort -rn : 내림차순 정렬
	// uniq -c : 여러 번 반복되는 줄을 삭제하고 하나로 만든 후 카운팅
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Printf("usage: %s logFile\n", filepath.Base(arguments[0]))
		os.Exit(1)
	}

	for _, filename := range arguments[1:] {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Printf("error opening file%s\n", err)
			os.Exit(-1)
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
				break
			}

			ip := findIP(line)
			//fmt.Println(line, ">>", ip)
			trial := net.ParseIP(ip)

			if trial.To4() == nil {
				continue
			}

			fmt.Println(ip)
		}
	}

}
