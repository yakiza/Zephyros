package fakedb

import (
	"errors"
	"github.com/yakiza/Zephyros"
	"sync"
)

type FakeProductRepository struct {
	mux         sync.Mutex
	productList map[Zephyros.ProductCharacteristics]*Zephyros.Product
}

func (repo *FakeProductRepository) Get(productCharacteristics Zephyros.ProductCharacteristics) (bool, *Zephyros.Product) {
	repo.mux.Lock()
	defer repo.mux.Unlock()

	if value, ok := repo.productList[productCharacteristics]; ok {
		return ok, value
	}

	return false, nil
}

func (repo *FakeProductRepository) Add(product Zephyros.Product) error {
	repo.mux.Lock()
	defer repo.mux.Unlock()

	productCharacteristics := Zephyros.MakeProductCharacteristics(product.Brand, product.Model, product.Season)
	if _, ok := repo.productList[productCharacteristics]; ok {
		return errors.New("duplicate")
	}

	repo.productList[productCharacteristics] = &product

	return nil
}

func NewFakeProductRepository() *FakeProductRepository {
	return &FakeProductRepository{
		productList: make(map[Zephyros.ProductCharacteristics]*Zephyros.Product),
	}
}
