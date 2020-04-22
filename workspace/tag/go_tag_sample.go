package main

import (
	"fmt"
	"reflect"
)

type T1 struct {
	f1 string "f one"
	f2 int64 `f two`
	f3, f4 float64 `f three and four`
}

func main() {
	t := reflect.TypeOf(T1{})
	f1, _ := t.FieldByName("f1")
	fmt.Println(">>>> Display field f1 <<<<")
	fmt.Println("Tag", f1.Tag)
	fmt.Println("Name", f1.Name)
	fmt.Println("Type", f1.Type)



}
