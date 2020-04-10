package dblayer

import "github.com/zacscoding/learning-go/fullstack/chap07/backend/src/models"

type DBLayer interface {
	GetAllProducts() ([]models.Product, error)
	GetPromos() ([]models.Product, error)
	GetCustomerByName(firstname, lastname string) (models.Customer, error)
	GetCustomerByID(id int) (models.Customer, error)
	GetProduct(id int) (models.Product, error)
	AddUser(customer models.Customer) (models.Customer, error)
	SignInUser(email, password string) (models.Customer, error)
	SignOutUserById(id int) error
	GetCustomerOrdersByID(id int) ([]models.Order, error)
}
