package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Epoch time:", time.Now().Unix())
	t := time.Now()
	fmt.Println(t, t.Format(time.RFC3339))
	fmt.Println(t.Weekday(), t.Day(), t.Month(), t.Year())

	time.Sleep(time.Second)
	t1 := time.Now()
	fmt.Println("Time difference:", t1.Sub(t))

	//Epoch time: 1566805572
	//2019-08-26 16:46:12.671992866 +0900 KST m=+0.000088733 2019-08-26T16:46:12+09:00
	//Monday 26 August 2019
	//Time difference: 1.000246964s

	formatT := t.Format("01 January 2006")
	fmt.Println(formatT)
	loc, _ := time.LoadLocation("Europe/Paris")
	londonTime := t.In(loc)
	fmt.Println("Paris", londonTime)

	//08 August 2019
	//Paris 2019-08-26 09:48:26.229622198 +0200 CEST
}
