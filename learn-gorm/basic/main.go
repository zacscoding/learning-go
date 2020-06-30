package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type RunTest func(db *gorm.DB) error

func main() {
	tests := []RunTest{
		//testRawQueryVariable,
		// testWhereIn,
		// testDelete,
		// testEnum,
		// testTimestamp,
		testTimeout,
	}

	db, err := gorm.Open("mysql", "root:password@tcp(127.0.0.1:13306)/my_database?charset=utf8&parseTime=True")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.LogMode(true)

	for _, t := range tests {
		err := t(db)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}
}

func toJson(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}
