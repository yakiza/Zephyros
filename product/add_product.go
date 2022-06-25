package product

import (
	"github.com/yakiza/Zephyros"
	"github.com/yakiza/Zephyros/repository"
	"go.uber.org/zap"
)

type AddProductHandler struct {
	Logger            *zap.Logger
	ProductRepository Zephyros.AddProductService
}

func (h AddProductHandler) AddProduct(product Zephyros.Product) error {
	// TODO: Check that passed in data is valid
	h.Logger.Info("Adding Product", zap.Any("Product", product))
	err := h.ProductRepository.AddProduct(product)
	if err != nil {
		return err
	}

	return nil
}

func NewAddProductHandler(logger *zap.Logger, productRepository repository.ProductRepository) AddProductHandler {
	return AddProductHandler{
		Logger:            logger,
		ProductRepository: productRepository,
	}
}
