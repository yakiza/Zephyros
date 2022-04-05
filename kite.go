package Zephyros

// Kite entity
type Kite struct {
	Brand string
	Model string
	Size  float32
	Color string
	Year  int
}

type Add interface {
	Add(kite Kite) error
}

type Exist interface {
	Exist(kite Kite) (bool, error)
}
