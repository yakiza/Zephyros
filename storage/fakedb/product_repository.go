package fakedb

import (
	"errors"
	"github.com/yakiza/Zephyros"
	"sync"
)

type FakeProductRepository struct {
	mux         sync.Mutex
	productList []*Zephyros.Product
}

func (repo *FakeProductRepository) AddProduct(product Zephyros.Product) error {
	repo.mux.Lock()
	defer repo.mux.Unlock()

	for _, p := range repo.productList {
		if p == &product {
			return errors.New("duplicate")
		}
	}

	repo.productList = append(repo.productList, &product)

	return nil
}

func NewFakeProductRepository() *FakeProductRepository {
	return &FakeProductRepository{
		productList: []*Zephyros.Product{},
	}
}
