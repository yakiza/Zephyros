package product

import (
	"github.com/yakiza/Zephyros"
	"github.com/yakiza/Zephyros/repository"
	"go.uber.org/zap"
)

type AddHandler struct {
	Logger            *zap.Logger
	ProductRepository Zephyros.AddService
}

func (h AddHandler) Add(kite Zephyros.Product) error {
	// TODO: Check that passed in data is valid
	h.Logger.Info("Adding Product", zap.Any("Product", kite))
	err := h.ProductRepository.Add(kite)
	if err != nil {
		return err
	}

	return nil
}

func NewAddHandler(logger *zap.Logger, productRepository repository.ProductRepository) AddHandler {
	return AddHandler{
		Logger:            logger,
		ProductRepository: productRepository,
	}
}
