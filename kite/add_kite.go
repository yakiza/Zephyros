package kite

import (
	"github.com/yakiza/Zephyros"
	"github.com/yakiza/Zephyros/repository"
)

type AddHandler struct {
	KiteRepository repository.KiteRepository
}

func (h AddHandler) Add(kite Zephyros.Kite) error {
	err := h.KiteRepository.Add(kite)
	if err != nil {
		return err
	}

	return nil
}

func NewAddHandler(kiteRepository repository.KiteRepository) AddHandler {
	return AddHandler{
		KiteRepository: kiteRepository,
	}
}
