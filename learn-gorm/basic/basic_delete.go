package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
)

type InternalJson struct {
	Field1 string
	Field2 string
}

type DeleteUser struct {
	gorm.Model
	Name     string
	Age      int
	Internal json.RawMessage `sql:"type:json;NOT NULL"`
}

func testDelete(db *gorm.DB) error {
	db.DropTableIfExists(&DeleteUser{})
	db.AutoMigrate(&DeleteUser{})

	i := InternalJson{
		Field1: "f1",
		Field2: "f2",
	}
	b, _ := json.Marshal(i)
	u := DeleteUser{
		Name:     "user1",
		Age:      10,
		Internal: b,
	}
	err := db.Create(&u).Error
	if err != nil {
		return err
	}
	fmt.Println("Success to save")

	db = db.Where("id = ?", u.ID).Unscoped().Delete(&DeleteUser{})
	if db.Error != nil {
		fmt.Println("error :", db.Error)
		return nil
	}
	return nil

	//err = db.Unscoped().Delete(&DeleteUser{Model: gorm.Model{ID: u.ID + 2}}).Error
	//if err != nil {
	//	fmt.Println("failed to delete ", err)
	//	return err
	//}

	//err = db.Where("id = ?", u.ID).Unscoped().Delete(&DeleteUser{}).Error
	//if err != nil {
	//	fmt.Println("failed to delete", err)
	//	return err
	//}

	//err = db.Unscoped().Delete(&u).Error
	//if err != nil {
	//	return err
	//}

	return nil
}
