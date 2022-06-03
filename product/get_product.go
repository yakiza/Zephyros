package product

import (
	"github.com/yakiza/Zephyros"
	"go.uber.org/zap"
)

type GetProductHandler struct {
	logger  *zap.Logger
	service Zephyros.GetService
}

func (h *GetProductHandler) Get(productCharacteristics Zephyros.ProductCharacteristics) (Zephyros.Product, error) {
	h.logger.Info("Get", zap.Any("ProductCharacteristics", productCharacteristics))
	product, err := h.service.Get(productCharacteristics)
	if err != nil {
		return Zephyros.Product{}, err
	}
	return *product, nil
}

func NewGetProductHandler() GetProductHandler {
	return GetProductHandler{}
}
