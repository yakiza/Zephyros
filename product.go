package Zephyros

// Product entity
type Product struct {
	Brand string
	Model string
	Size  float32
	Color string
	Year  int
}

type Add interface {
	Add(product Product) error
}

type Exist interface {
	Exist(product Product) (bool, error)
}
