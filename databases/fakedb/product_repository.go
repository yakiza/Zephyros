package fakedb

import (
	"github.com/yakiza/Zephyros"
	"sync"
)

type FakeProductRepository struct {
	mux         sync.Mutex
	productList map[Zephyros.ProductCharacteristics]*Zephyros.Product
}

func (repo *FakeProductRepository) Exist(kite Zephyros.Product) (bool, error) {
	//TODO implement me
	panic("implement me")
	return false, nil
}

func (repo *FakeProductRepository) Add(product Zephyros.Product) error {
	repo.mux.Lock()
	defer repo.mux.Unlock()

	productCharacteristics := Zephyros.MakeProductCharacteristics(product.Brand, product.Model, product.Year)
	repo.productList[productCharacteristics] = &product

	return nil
}

func NewFakeKiteRepository() *FakeProductRepository {
	return &FakeProductRepository{
		productList: make(map[Zephyros.ProductCharacteristics]*Zephyros.Product),
	}
}
