package dblayer

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/zacscoding/learning-go/fullstack/chap07/backend/src/models"
	"golang.org/x/crypto/bcrypt"
)

type DBORM struct {
	*gorm.DB
}

// NewORM returns a new gorm.DB given dbname and connection info.
func NewORM(dbname, con string) (*DBORM, error) {
	db, err := gorm.Open(dbname, con)
	if err != nil {
		return nil, err
	}
	return &DBORM{
		DB: db,
	}, nil
}

// GetAllProducts returns all products in "products" table.
func (db *DBORM) GetAllProducts() ([]models.Product, error) {
	// select * from products
	var products []models.Product
	return products, db.Find(&products).Error
}

// GetPromos returns products that not null promotion value in "products" table.
func (db *DBORM) GetPromos() ([]models.Product, error) {
	// select * from products where promotion IS NOT NULL
	var products []models.Product
	return products, db.Where("promotion IS NOT NULL").Find(&products).Error
}

// GetCustomerByName returns a customer given first name and last name in "customers" table.
func (db *DBORM) GetCustomerByName(firstname, lastname string) (models.Customer, error) {
	// select * from customers where fistname = '..' and lastname = '..'
	var customer models.Customer
	return customer, db.Where(&models.Customer{FirstName: firstname, LastName: lastname}).Find(&customer).Error
}

// GetCustomerByID returns a customer given id in "customers" table
func (db *DBORM) GetCustomerByID(id int) (models.Customer, error) {
	var customer models.Customer
	return customer, db.First(&customer, id).Error
}

// GetProduct returns a product given id in "products" table.
func (db *DBORM) GetProduct(id int) (models.Product, error) {
	var product models.Product
	return product, db.First(&product, id).Error
}

//AddUser save a given customer in "customers" table.
func (db *DBORM) AddUser(customer models.Customer) (models.Customer, error) {
	_ = hashPassword(&customer.Pass)
	customer.LoggedIn = true
	err := db.Create(&customer).Error
	customer.Pass = ""
	return customer, err
}

//SignInUser signin given email and password
//i.e update true(1) to customers.loggedin column if valid user.
func (db *DBORM) SignInUser(email, password string) (models.Customer, error) {
	// 사용자 행을 나타내는 *gorm.DB 타입 할당
	var customer models.Customer
	result := db.Table("Customers").Where(&models.Customer{Email: email})
	// 입력된 이메일로 사용자 정보 조회
	err := result.First(&customer).Error
	if err != nil {
		return customer, err
	}
	// 패스워드 문자열과 해시 값 비교
	if !checkPassword(customer.Pass, password) {
		return customer, ErrInvalidPassword
	}

	// 공유되지 않도록 패스워드 문자열은 지운다
	customer.Pass = ""

	// loggedin 필드 업데이트
	err = result.Update("loggedin", 1).Error
	if err != nil {
		return customer, err
	}

	// 사용자 행 반환
	return customer, result.Find(&customer).Error
}

//SignOutUserById sign out given customer's id
//i.e update false(0) to customers.loggedin.
func (db *DBORM) SignOutUserById(id int) error {
	// ID에 해당하는 사용자 구조체 생성
	customer := models.Customer{
		Model: gorm.Model{
			ID: uint(id),
		},
	}

	return db.Table("Customers").Where(&customer).Update("loggedin", 0).Error
}

//GetCustomerOrdersByID returns orders given customer's id
func (db *DBORM) GetCustomerOrdersByID(id int) ([]models.Order, error) {
	var orders []models.Order
	return orders, db.Table("orders").Select("*").Joins("join customers on customers.id = customer_id").Joins("join products on product_id = product_id").Where("customer_id = ?", id).Scan(&orders).Error
}

func hashPassword(s *string) error {
	if s == nil {
		return errors.New("reference povided for hashing password is nil")
	}

	sBytes := []byte(*s)
	hashedBytes, err := bcrypt.GenerateFromPassword(sBytes, bcrypt.DefaultCost)
	if err != nil {
		return nil
	}
	*s = string(hashedBytes[:])
	return nil
}

func checkPassword(existingHash, incomingPass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(existingHash), []byte(incomingPass)) == nil
}
