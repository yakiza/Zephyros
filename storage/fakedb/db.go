package fakedb

import "github.com/yakiza/Zephyros/repository"

type DB struct{}

func (db *DB) Product() repository.ProductRepository {
	return NewFakeProductRepository()
}
