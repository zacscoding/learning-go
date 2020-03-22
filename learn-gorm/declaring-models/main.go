package main

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

// https://gorm.io/docs/models.html
type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        // set field size to 255
	MemberNumber *string `gorm:"unique;not null"` // set member number to unique and not null
	Num          int     `gorm:"AUTO_INCREMENT"`  // set num to auto incrementable
	Address      string  `gorm:"index:addr"`      // create index with name `addr` for address
	IgnoreMe     int     `gorm:"-"`               // ignore this field
}

func main() {
	db, err := gorm.Open("postgres", "host=192.168.79.130 port=5432 user=tester dbname=testdb password=tester sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{})
	// generated script
	// create table users (
	//    id            serial not null constraint users_pkey primary key,
	//    created_at    timestamp with time zone,
	//    updated_at    timestamp with time zone,
	//    deleted_at    timestamp with time zone,
	//    name          text,
	//    age           bigint,
	//    birthday      timestamp with time zone,
	//    email         varchar(100),
	//    role          varchar(255),
	//    member_number text   not null constraint users_member_number_key unique,
	//    num           serial not null,
	//    address       text
	// );
	// alter table users owner to tester;
	// create index idx_users_deleted_at on users(deleted_at);
	// create index addr on users(address);
	// create unique index uix_users_email on users(email);
}
