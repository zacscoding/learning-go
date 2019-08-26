package main

import (
	"fmt"
	"sort"
)

type aStructure struct {
	person string
	height int
	weigh  int
}

func main() {
	mySlice := make([]aStructure, 0)
	mySlice = append(mySlice, aStructure{"Mihalis", 180, 90})
	mySlice = append(mySlice, aStructure{"Bill", 134, 45})
	mySlice = append(mySlice, aStructure{"Marietta", 155, 45})
	mySlice = append(mySlice, aStructure{"Epifanios", 144, 50})
	mySlice = append(mySlice, aStructure{"Athina", 134, 40})

	fmt.Println("0:", mySlice)

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height < mySlice[j].height
	})
	fmt.Println("<:", mySlice)

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height > mySlice[j].height
	})
	fmt.Println(">:", mySlice)

	// output
	//0: [{Mihalis 180 90} {Bill 134 45} {Marietta 155 45} {Epifanios 144 50} {Athina 134 40}]
	//<: [{Bill 134 45} {Athina 134 40} {Epifanios 144 50} {Marietta 155 45} {Mihalis 180 90}]
	//>: [{Mihalis 180 90} {Marietta 155 45} {Epifanios 144 50} {Bill 134 45} {Athina 134 40}]
}
