package fakedb

import "github.com/yakiza/Zephyros/repository"

type DB struct{}

func (db *DB) Kite() repository.ProductRepository {
	return NewFakeKiteRepository()
}
