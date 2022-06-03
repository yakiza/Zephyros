package repository

import "github.com/yakiza/Zephyros"

type ProductRepository interface {
	// Add a new Product to the respective database given the product entity
	Add(product Zephyros.Product) error

	// Get returns teh product as well as a boolen
	Get(productCharacteristics Zephyros.ProductCharacteristics) (bool, *Zephyros.Product)
}
