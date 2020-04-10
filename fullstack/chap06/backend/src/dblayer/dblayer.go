package dblayer

import (
	models2 "github.com/zacscoding/learning-go/fullstack/chap07/backend/src/models"
)

type DBLayer interface {
	GetAllProducts() ([]models2.Product, error)
	GetPromos() ([]models2.Product, error)
	GetCustomerByName(string, string) (models2.Customer, error)
	GetCustomerByID(int) (models2.Customer, error)
	GetProduct(uint) (models2.Product, error)
	AddUser(models2.Customer) (models2.Customer, error)
	SignInUser(username, password string) (models2.Customer, error)
	SignOutUserById(int) error
	GetCustomerOrdersByID(int) ([]models2.Order, error)
}
