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
	Addr Address `gorm:"EMBEDDED"` // Address :: OneToOne
}

type Category struct {
	Model
	ID   uint   `gorm:"primary_key;column:category_id"`
	Name string `gorm:"size:255"`
	// TODO : need parent, child
	Items []*Item `gorm:"many2many:category_item;"` // Items :: ManyToMany
}

type Item struct {
	Model
	ID            uint   `gorm:"primary_key;column:item_id"`
	Name          string `gorm:"size:255"`
	Price         int
	StockQuantity int
	Categories    []*Category `gorm:"many2many:category_item;"` // Categories :: ManyToMany
}

type Order struct {
	Model
	ID       uint `gorm:"primary_key;column:order_id"`
	MemberID uint
	Member   Member

	OrderItems []OrderItem // OrderItem :: OneToMany

	DeliveryID uint
	Delivery   Delivery // Delivery :: OneToOne

	Status string `gorm:"size:50"`
}

type OrderItem struct {
	Model
	ID         uint `gorm:"primary_key;column:order_item_id"`
	OrderID    uint
	ItemID     uint
	Item       Item
	OrderPrice uint
	Count      uint
}

type Delivery struct {
	Model
	ID             uint    `gorm:"primary_key;column:delivery_id"`
	Addr           Address `gorm:"EMBEDDED"`
	DeliveryStatus string  `gorm:"size:50"`
}

var DB *gorm.DB

func main() {
	setupDatabase()
	defer DB.Close()

	testInitialData()
}

func testInitialData() {
	// 멤버
	m1 := &Member{Name: "memberA",
		Addr: Address{
			City:    "city",
			Street:  "street",
			ZipCode: "1234",
		},
	}
	m2 := &Member{Name: "memberB",
		Addr: Address{
			City:    "city2",
			Street:  "street2",
			ZipCode: "12345",
		},
	}
	DB.Create(m1)
	DB.Create(m2)

	// 아이템
	i1 := &Item{Name: "Zaccoding's funny go", Price: 12000, StockQuantity: 50}
	i2 := &Item{Name: "Zaccoding's easy algorithm", Price: 15000, StockQuantity: 10}

	i3 := &Item{Name: "하버드 상위 1퍼센트의 비밀!", Price: 14000, StockQuantity: 15}
	i4 := &Item{Name: "5년 후 나에게 Q&A", Price: 18000, StockQuantity: 100}

	// 카테고리
	p1 := &Category{Name: "컴퓨터/IT", Items: []*Item{i1, i2}}
	p2 := &Category{Name: "자기계발", Items: []*Item{i1, i3, i4}}

	DB.Create(p1)
	DB.Create(p2)

	// Order
	o1 := &Order{
		Member: *m2,
		Delivery: Delivery{
			Addr:           m1.Addr,
			DeliveryStatus: "delivery",
		},
		OrderItems: []OrderItem{
			{
				Item:       *i1,
				OrderPrice: 100,
				Count:      2,
			},
			{
				Item:       *i3,
				OrderPrice: 50,
				Count:      1,
			},
		},
		Status: "ORDER",
	}
	DB.Create(o1)

	order := Order{Status: "ORDER"}
	DB.First(&order, o1.ID)
	fmt.Println("Read order ==>", order)
	DB.Model(&order).Related(&order.Delivery)
	fmt.Println("Read order's delivery ==>", order.Delivery)
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

	DB.Create(p1)
	DB.Create(p2)

	item := Item{}
	// SELECT * FROM "items"  WHERE ("items"."item_id" = 1) ORDER BY "items"."item_id" ASC LIMIT 1
	DB.First(&item, i1.ID)
	// SELECT "categories".* FROM "categories" INNER JOIN "category_item" ON
	// "category_item"."category_category_id" = "categories"."category_id" WHERE ("category_item"."item_item_id" IN (1))
	DB.Model(&item).Related(&item.Categories, "Categories")
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
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=tester dbname=testdb password=tester sslmode=disable")

	if err != nil {
		panic(err)
	}

	db.DB().SetMaxIdleConns(10)
	db.LogMode(true)
	// Drop all
	db.DropTable(&Member{}, &Category{}, &Item{}, &Order{}, &Delivery{}, &OrderItem{})
	// Migrate all
	db.AutoMigrate(&Member{}, &Category{}, &Item{}, &Order{}, &Delivery{}, &OrderItem{})

	DB = db
}

// References
//
// http://gorm.io/docs/has_many.html
// https://stackoverflow.com/questions/35821810/golang-gorm-one-to-many-with-has-one
