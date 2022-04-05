package kite

import (
	"github.com/yakiza/Zephyros"
	"github.com/yakiza/Zephyros/repository"
)

type ExistKiteHandler struct {
	exist repository.KiteRepository
}

func (h *ExistKiteHandler) Exist(kite Zephyros.Kite) (bool, error) {
	state, err := h.exist.Exist(kite)
	if err != nil {
		return state, err
	}

	return state, nil
}

func NewExistHandler(exist repository.KiteRepository) ExistKiteHandler {
	return ExistKiteHandler{exist: exist}
}
