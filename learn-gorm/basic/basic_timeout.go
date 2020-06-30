package main

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type TimeoutUser struct {
	gorm.Model
	Name string
}

// WithContext is not yet
// https://github.com/go-gorm/gorm/pull/1428
// https://github.com/go-gorm/gorm/issues/2960
func testTimeout(db *gorm.DB) error {
	db.DropTableIfExists(&TimeoutUser{})
	db.CreateTable(&TimeoutUser{})

	err := db.Save(&TimeoutUser{
		Name: "user1",
	}).Error
	if err != nil {
		return err
	}

	var u TimeoutUser

	err = db.Raw("SELECT /*+ MAX_EXECUTION_TIME(3000) */  * FROM timeout_users WHERE sleep(2)=0").Scan(&u).Error
	fmt.Println("## ", toJson(u))

	err = db.Raw("SELECT /*+ MAX_EXECUTION_TIME(3000) */ * FROM timeout_users WHERE sleep(5)=0").Scan(&u).Error
	if err != nil {
		switch err.(type) {
		case *mysql.MySQLError:
			mySqlErr := err.(*mysql.MySQLError)
			fmt.Println("MySQLError. code:", mySqlErr.Number, ", message:", mySqlErr.Message)
		default:
			fmt.Println("error with sleep(5)", err)
		}
	}
	return nil
}
