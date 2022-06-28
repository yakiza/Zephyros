package Zephyros

//Adding this as a safe mechanism, so that the entities are not used directly
type entity struct{}

// Product entity
type Product struct {
	_           entity
	Name        string
	Description string
	Brand       string
	Price       float64
	Currency    string
}

type AddProductService interface {
	AddProduct(product Product) error
}
