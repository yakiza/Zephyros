package Zephyros

type ProductCharacteristics struct {
	Brand string
	Model string
	Size  int
}

func MakeProductCharacteristics(brand string, model string, size int) ProductCharacteristics {
	return ProductCharacteristics{
		Brand: brand,
		Model: model,
		Size:  size,
	}
}
