package repository

import "github.com/yakiza/Zephyros"

type ProductRepository interface {
	// Add a new Product to the respective database given the kite entity
	Add(kite Zephyros.Product) error

	// Exist check if the kite already exists
	Exist(kite Zephyros.Product) (bool, error)
}
