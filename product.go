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

type Add interface {
	Add(product Product) error
}

type Exist interface {
	Exist(product Product) (bool, error)
}
