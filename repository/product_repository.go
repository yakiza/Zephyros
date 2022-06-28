package repository

import "github.com/yakiza/Zephyros"

type ProductRepository interface {
	// AddProduct a new Product to the respective database given the product entity
	AddProduct(product Zephyros.Product) error
}
