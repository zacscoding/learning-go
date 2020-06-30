package main

import (
	"github.com/jinzhu/gorm"
	"time"
)

type UserType int

func (ut UserType) String() string {
	if ut > 2 {
		return "unknown"
	}
	return [...]string{"admin", "service", "normal"}[ut]
}

type BasicEnumUser struct {
	gorm.Model
	Name string
	Type UserType `gorm:"column:type;type:TINYINT;DEFAULT:0"`
}

func testEnum(db *gorm.DB) error {
	db.DropTableIfExists(&BasicEnumUser{})
	db.CreateTable(&BasicEnumUser{})

	u0 := BasicEnumUser{
		Name: "basicEnumUser0",
		Type: 0,
	}
	u1 := BasicEnumUser{
		Name: "basicEnumUser1",
		Type: 1,
	}
	u2 := BasicEnumUser{
		Name: "basicEnumUser2",
		Type: 2,
	}
	u3 := BasicEnumUser{
		Name: "basicEnumUser3",
		Type: 3,
	}
	u0.DeletedAt = new(time.Time)
	*u0.DeletedAt = time.Now()
	db.Create(u0)
	db.Create(u1)
	db.Create(u2)
	db.Create(u3)

	//var users []BasicEnumUser
	//if err := db.Find(users).Error; err != nil {
	//	return err
	//}
	//for _, u := range users {
	//	fmt.Println(u.Name, "=>", u.Type.String())
	//}
	return nil
}
