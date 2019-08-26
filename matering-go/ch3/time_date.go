package main

import (
	"fmt"
	"regexp"
	"time"
)

func main() {
	logs := []string{
		"127.0.0.1 - - [16/Nov/2017:10:49:46 +0200] 325504",
		"127.0.0.1 - - [16/Nov/2017:10:16:41 +0200] \"GET /CVEN HTTP/1.1\" 200 12531 \"-\" \"Mozilla/5.0 AppleWebKit/537.36",
		"127.0.0.1 200 9412 - - [12/Nov/2017:06:26:05 +0200] \"GET \"http://www.mtsoukalos.eu/taxonomy/term/47\" 1507",
		"[12/Nov/2017:16:27:21 +0300]",
		"[12/Nov/2017:20:88:21 +0200]",
		"[12/Nov/2017:20:21 +0200]",
	}

	for _, logEntry := range logs {
		r := regexp.MustCompile(`.*\[(\d\d\/\w+/\d\d\d\d:\d\d:\d\d:\d\d.*)\].*`)

		if !r.MatchString(logEntry) {
			fmt.Println("Not a match :", logEntry)
			continue
		}
		match := r.FindStringSubmatch(logEntry)
		dt, err := time.Parse("02/Jan/2006:15:04:05 -0700", match[1])

		if err != nil {
			fmt.Println("Not a valid date time format!", logEntry)
			continue
		}

		newFormat := dt.Format(time.RFC850)
		fmt.Println(newFormat)
	}
}

//Thursday, 16-Nov-17 10:49:46 +0200
//Thursday, 16-Nov-17 10:16:41 +0200
//Sunday, 12-Nov-17 06:26:05 +0200
//Sunday, 12-Nov-17 16:27:21 +0300
//Not a valid date time format! [12/Nov/2017:20:88:21 +0200]
//Not a match : [12/Nov/2017:20:21 +0200]

