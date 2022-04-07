package kite

import (
	"github.com/yakiza/Zephyros"
	"github.com/yakiza/Zephyros/repository"
)

type AddHandler struct {
	ProductRepository repository.ProductRepository
}

func (h AddHandler) Add(kite Zephyros.Product) error {
	// TODO: Check that passed in data is valid
	// TODO: Check if it already exists
	err := h.ProductRepository.Add(kite)
	if err != nil {
		return err
	}

	return nil
}

func NewAddHandler(productRepository repository.ProductRepository) AddHandler {
	return AddHandler{
		ProductRepository: productRepository,
	}
}
