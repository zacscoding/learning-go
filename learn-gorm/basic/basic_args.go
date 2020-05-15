package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type BasicArgsUser struct {
	gorm.Model
	Name     []byte
	Nickname []byte
}

func testRawQueryVariable(db *gorm.DB) error {
	db.CreateTable(&BasicArgsUser{})
	db.Create(&BasicArgsUser{
		Name:     []byte("name1"),
		Nickname: []byte("nickname1"),
	})
	db.Create(&BasicArgsUser{
		Name:     []byte("name1"),
		Nickname: []byte("nickname2"),
	})
	db.Create(&BasicArgsUser{
		Name:     []byte("name1"),
		Nickname: []byte("nickname1"),
	})
	db.Create(&BasicArgsUser{
		Name:     []byte("name2"),
		Nickname: []byte("nickname3"),
	})
	db.Create(&BasicArgsUser{
		Name:     []byte("name2"),
		Nickname: []byte("nickname2"),
	})

	users, err := findUsers(db, "name1", "nickname1")
	if err != nil {
		return err
	}
	for i, u := range users {
		fmt.Printf("[%d] User name : %s, nickname : %s\n", i, string(u.Name), string(u.Nickname))
	}
	return nil
}

func findUsers(db *gorm.DB, name, nickname string) ([]BasicArgsUser, error) {
	var users []BasicArgsUser
	db = db.Where("name = ?", toBytes(name))
	if nickname != "" {
		db = db.Where("nickname = ?", toBytes(nickname))
	}

	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func toBytes(s string) []byte {
	if s == "" {
		return nil
	}
	return []byte(s)
}
