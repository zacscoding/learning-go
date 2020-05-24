package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

type BasicArgsUser struct {
	ID       uint `gorm:"primary_key"`
	Name     []byte
	Nickname []byte
}

func (u BasicArgsUser) String() string {
	return fmt.Sprintf("BasicArgsUser{ID:%d, Name: %s, Nickname : %s}", u.ID, u.Name, u.Nickname)
}

type nameArg []byte

func testWhereIn(db *gorm.DB) error {
	db.DropTableIfExists(&BasicArgsUser{})
	db.CreateTable(&BasicArgsUser{})

	_ = saveUser(db, "user1", "user2") // 1 exclude
	_ = saveUser(db, "user1", "user3") // 2 include
	_ = saveUser(db, "user3", "user1") // 3 include
	_ = saveUser(db, "user5", "user1") // 4 include
	_ = saveUser(db, "user3", "user6") // 5 include
	_ = saveUser(db, "user2", "user1") // 6 exclude
	_ = saveUser(db, "user3", "user2") // 7 exclude
	_ = saveUser(db, "user2", "user5") // 8 exclude

	include := []nameArg{toBytes("user1"), toBytes("user3")}
	exclude := []nameArg{toBytes("user2")}

	users, err := findUsersWithFilter(db, include, exclude)
	if err != nil {
		return err
	}

	fmt.Println("Len :", len(users))
	for i, user := range users {
		names := string(user.Name) + " " + string(user.Nickname)
		if !(strings.Contains(names, "user1") || strings.Contains(names, "user3")) {
			fmt.Println("Not contains user1 or user3")
		}

		if strings.Contains(names, "user2") {
			fmt.Println("Contains user2")
		}

		fmt.Printf("[%d] => %s\n", i, user.String())
	}
	return nil
}

func findUsersWithFilter(db *gorm.DB, includes, excludes []nameArg) ([]BasicArgsUser, error) {
	var users []BasicArgsUser
	chain := db
	if len(includes) != 0 {
		chain = chain.Where("name IN (?) OR nickname IN (?)", includes, includes)
	}

	if len(excludes) != 0 {
		chain = chain.Where("name NOT IN (?)", excludes).Where("nickname NOT IN (?)", excludes)
	}

	return users, chain.Find(&users).Error
}

func testRawQueryVariable(db *gorm.DB) error {
	db.CreateTable(&BasicArgsUser{})
	_ = saveUser(db, "name1", "nickname1")
	_ = saveUser(db, "name1", "nickname2")
	_ = saveUser(db, "name1", "nickname1")
	_ = saveUser(db, "name2", "nickname3")
	_ = saveUser(db, "name2", "nickname2")

	users, err := findUsers(db, "name1", "nickname1")
	if err != nil {
		return err
	}
	displayUsers(users)
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

func displayUsers(users []BasicArgsUser) {
	for i, u := range users {
		fmt.Printf("[%d] %s\n", i, u.String())
	}
}

func saveUser(db *gorm.DB, name, nickname string) error {
	return db.Create(&BasicArgsUser{
		Name:     []byte(name),
		Nickname: []byte(nickname),
	}).Error
}

func toBytes(s string) []byte {
	if s == "" {
		return nil
	}
	return []byte(s)
}
