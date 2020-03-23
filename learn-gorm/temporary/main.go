package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// User has many CreditCards, UserID is the foreign key
type User struct {
	gorm.Model
	CreditCards []CreditCard
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

func main() {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=tester dbname=testdb password=tester sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.LogMode(true)
	db.DropTableIfExists(&User{}, &CreditCard{})
	db.AutoMigrate(&User{}, &CreditCard{})

	c1 := &CreditCard{
		Number: "c1",
	}
	c2 := &CreditCard{
		Number: "c1",
	}

	u1 := &User{
		CreditCards: []CreditCard{*c1, *c2},
	}
	db.Create(u1)
}
