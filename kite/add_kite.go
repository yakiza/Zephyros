package kite

import (
	"github.com/yakiza/Zephyros"
	"github.com/yakiza/Zephyros/repository"
)

type AddHandler struct {
	KiteRepository repository.KiteRepository
}

func (h AddHandler) Add(kite Zephyros.Kite) error {
	// TODO: Check that passed in data is valid
	// TODO: Check if it already exists
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
