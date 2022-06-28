package Zephyros

// Product entity
type Product struct {
	Name   string
	Brand  string
	Model  string
	Type   string
	Color  string
	Size   float32
	Season int
}

type AddService interface {
	Add(product Product) error
}

type GetService interface {
	Get(productCharacteristics ProductCharacteristics) (*Product, error)
}
