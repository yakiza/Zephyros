package fakedb

import (
	"github.com/yakiza/Zephyros"
	"sync"
)

type FakeKiteRepository struct {
	mux      sync.Mutex
	kiteList map[Zephyros.KiteCharacteristics]*Zephyros.Kite
}

func (repo *FakeKiteRepository) Exist(kite Zephyros.Kite) error {
	//TODO implement me
	panic("implement me")
}

func (repo *FakeKiteRepository) Add(kite Zephyros.Kite) error {
	repo.mux.Lock()
	defer repo.mux.Unlock()

	kiteCharacteristics := Zephyros.MakeKiteCharacteristics(kite.Brand, kite.Model, kite.Year)
	repo.kiteList[kiteCharacteristics] = &kite

	return nil
}

func NewFakeKiteRepository() *FakeKiteRepository {
	return &FakeKiteRepository{
		kiteList: make(map[Zephyros.KiteCharacteristics]*Zephyros.Kite),
	}
}
