package main

import (
	"github.com/jinzhu/gorm"
	"time"
)

type BasicTimestampUser struct {
	Name      string
	CreatedAt *time.Time
}

func testTimestamp(db *gorm.DB) error {
	db.DropTableIfExists(&BasicTimestampUser{})
	db.AutoMigrate(&BasicTimestampUser{})

	u0 := BasicTimestampUser{
		Name: "user0",
	}
	u0.CreatedAt = new(time.Time)
	*u0.CreatedAt = time.Now()

	return db.Create(u0).Error
}
