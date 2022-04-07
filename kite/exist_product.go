package kite

import (
	"github.com/yakiza/Zephyros"
	"github.com/yakiza/Zephyros/repository"
)

type ExistProductHandler struct {
	exist repository.ProductRepository
}

func (h *ExistProductHandler) Exist(product Zephyros.Product) (bool, error) {
	state, err := h.exist.Exist(product)
	if err != nil {
		return state, err
	}

	return state, nil
}

func NewExistHandler(exist repository.ProductRepository) ExistProductHandler {
	return ExistProductHandler{exist: exist}
}
