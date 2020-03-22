package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

// Model is a basic model definition
type Model struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Address is a embedded model
type Address struct {
	City    string `gorm:"size:255"`
	Street  string `gorm:"size:255"`
	ZipCode string `gorm:"size:255"`
}

type Member struct {
	Model
	ID   uint    `gorm:"primary_key;column:member_id"`
	Name string  `gorm:"size:255"`
	Addr Address `gorm:"EMBEDDED"`
}

type Category struct {
	Model
	ID   uint   `gorm:"primary_key;column:category_id"`
	Name string `gorm:"size:255"`
	// TODO : need parent, child
	Items []*Item `gorm:"many2many:category_item;"`
}

type Item struct {
	Model
	ID            uint   `gorm:"primary_key;column:item_id"`
	Name          string `gorm:"size:255"`
	Price         int
	StockQuantity int
	Categories    []*Category `gorm:"many2many:category_item;"`
}

// TODO : order, order item, delivery

var db *gorm.DB

func main() {
	setupDatabase()
	defer db.Close()

	testCategoriesItems()
}

func testCategoriesItems() {
	// 아이템
	i1 := &Item{Name: "Zaccoding's funny go", Price: 12000, StockQuantity: 50}
	i2 := &Item{Name: "Zaccoding's easy algorithm", Price: 15000, StockQuantity: 10}

	i3 := &Item{Name: "하버드 상위 1퍼센트의 비밀!", Price: 14000, StockQuantity: 15}
	i4 := &Item{Name: "5년 후 나에게 Q&A", Price: 18000, StockQuantity: 100}

	// 카테고리
	p1 := &Category{Name: "컴퓨터/IT", Items: []*Item{i1, i2}}
	p2 := &Category{Name: "자기계발", Items: []*Item{i1, i3, i4}}

	db.Create(p1)
	db.Create(p2)

	item := Item{}
	// SELECT * FROM "items"  WHERE ("items"."item_id" = 1) ORDER BY "items"."item_id" ASC LIMIT 1
	db.First(&item, i1.ID)
	// SELECT "categories".* FROM "categories" INNER JOIN "category_item" ON
	// "category_item"."category_category_id" = "categories"."category_id" WHERE ("category_item"."item_item_id" IN (1))
	db.Model(&item).Related(&item.Categories, "Categories")
	fmt.Println(item)
	for _, category := range item.Categories {
		fmt.Println(category)
	}
}

// TODO : test add parent and read child
func testCategories() {
	//// 최상위 카테고리 추가
	//p := &Category{Name: "컴퓨터/IT"}
	//db.Create(p)
	//
	//// 하위 카테고리 추가
	//c1 := &Category{Name: "프로그래밍", ParentID: p.ID}
	//db.Create(c1)
	//
	//c2 := &Category{Name: "자료구조", ParentID: p.ID}
	//db.Create(c2)
}

func setupDatabase() {
	var err error
	db, err = gorm.Open("postgres", "host=192.168.79.130 port=5432 user=tester dbname=testdb password=tester sslmode=disable")

	if err != nil {
		panic(err)
	}

	db.LogMode(true)
	// Drop all
	db.DropTable(&Member{}, &Category{}, &Item{})
	// Migrate all
	db.AutoMigrate(&Member{}, &Category{}, &Item{})
}

// References
//
// http://gorm.io/docs/has_many.html
// https://stackoverflow.com/questions/35821810/golang-gorm-one-to-many-with-has-one
