package repository

import "github.com/yakiza/Zephyros"

type ProductRepository interface {
	// Add a new Product to the respective database given the kite entity
	Add(product Zephyros.Product) error

	// Exist check if the product already exists
	Exist(product Zephyros.Product) (bool, error)
}
