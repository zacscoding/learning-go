package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"reflect"
	"time"
)

// https://www.joinc.co.kr/w/man/12/golang/tag

type T1 struct {
	f1     string  "f one"
	f2     int64   `f two`
	f3, f4 float64 `f three and four`
}

func testDefaultGoTag() {
	fmt.Println("====== Test default tag usage ======")
	t := reflect.TypeOf(T1{})
	fmt.Println("len of T1 struct's type :", t.NumField())

	for i := 0; i < t.NumField(); i++ {
		fmt.Println("-----------------------")
		fmt.Println("Check :", i)
		s := t.Field(i)
		fmt.Println("Name :", s.Name)
		fmt.Println("Tag :", s.Tag)
		fmt.Println("Type :", s.Type)
	}
	fmt.Println("=====================================")

	// Output
	//====== Test default tag usage ======
	//len of T1 struct's type : 4
	//-----------------------
	//Check : 0
	//Name : f1
	//Tag : f one
	//Type : string
	//-----------------------
	//Check : 1
	//Name : f2
	//Tag : f two
	//Type : int64
	//-----------------------
	//Check : 2
	//Name : f3
	//Tag : f three and four
	//Type : float64
	//-----------------------
	//Check : 3
	//Name : f4
	//Tag : f three and four
	//Type : float64
	//=====================================
}

type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`
	MemberNumber *string `gorm:"unique;not null"`
	Num          int     `gorm:"AUTO_INCREMENT"`
	Address      string  `gorm:"index:addr"`
	IgnoreMe     int     `gorm:"-"`
}

func testOneTag() {
	fmt.Println("====== Test single tag lookup ======")
	t := reflect.TypeOf(User{})
	for i := 0; i < t.NumField(); i++ {
		v, ok := t.Field(i).Tag.Lookup("gorm")
		if ok {
			fmt.Printf("(%14s) is ORM Field : %s\n", t.Field(i).Name, v)
		}
	}
	fmt.Println("=====================================")
	// Output
	//====== Test tag lookup ======
	//(         Email) is ORM Field : type:varchar(100);unique_index
	//(          Role) is ORM Field : size:255
	//(  MemberNumber) is ORM Field : unique;not null
	//(           Num) is ORM Field : AUTO_INCREMENT
	//(       Address) is ORM Field : index:addr
	//(      IgnoreMe) is ORM Field : -
	//=====================================
}

type User2 struct {
	Name    string `json:"name" xml:"name"`
	Address string `json:"address" xml:"address"`
	Age     string `json:"age"`
}

func testMultiTag() {
	fmt.Println("====== Test multi tag lookup ======")
	t := reflect.TypeOf(User2{})
	for i := 0; i < t.NumField(); i++ {
		if v, ok := t.Field(i).Tag.Lookup("json"); ok {
			fmt.Printf("(%14s) has json tag : %s\n", t.Field(i).Name, v)
		}

		if v, ok := t.Field(i).Tag.Lookup("xml"); ok {
			fmt.Printf("(%14s) has xml tag : %s\n", t.Field(i).Name, v)
		}
	}
	fmt.Println("=====================================")
	// Output
	//====== Test multi tag lookup ======
	//(          Name) has json tag : name
	//(          Name) has xml tag : name
	//(       Address) has json tag : address
	//(       Address) has xml tag : address
	//(           Age) has json tag : age
	//=====================================
}

func main() {
	testDefaultGoTag()
	testOneTag()
	testMultiTag()
}
