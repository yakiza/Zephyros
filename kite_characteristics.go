package Zephyros

type KiteCharacteristics struct {
	Brand string
	Model string
	Size  int
}

func MakeKiteCharacteristics(brand string, model string, size int) KiteCharacteristics {
	return KiteCharacteristics{
		Brand: brand,
		Model: model,
		Size:  size,
	}
}
