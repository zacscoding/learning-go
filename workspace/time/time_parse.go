package main

import (
	"fmt"
	"time"
)

func main() {
	parse1()
}

func parse1() {
	given := "00:01:30"
	t, err := time.Parse("15:04:05", given)
	if err != nil {
		panic(err)
	}

	fmt.Println("Hour :", t.Hour())
	fmt.Println("Minute :", t.Minute())
	fmt.Println("Sec :", t.Second())

	d := 0
	d += t.Hour() * int(time.Hour)
	d += t.Minute() * int(time.Minute)
	d += t.Second() * int(time.Second)

	duration := time.Duration(d)
	fmt.Println("Duration :", duration)
}
