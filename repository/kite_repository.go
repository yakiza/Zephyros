package repository

import "github.com/yakiza/Zephyros"

type KiteRepository interface {
	// Add a new Kite to the respective database given the kite entity
	Add(kite Zephyros.Kite) error

	// Exist check if the kite already exists
	Exist(kite Zephyros.Kite) (bool, error)
}
