package api

import "github.com/yakiza/Zephyros"

type KiteJson struct {
	Brand string  `json:"brand"`
	Model string  `json:"model"`
	Size  float32 `json:"size"`
	Color string  `json:"color"`
	Year  int     `json:"year"`
}

func MarshallKite(kite KiteJson) Zephyros.Product {
	return Zephyros.Product{
		Brand: kite.Brand,
		Model: kite.Model,
		Size:  kite.Size,
		Color: kite.Color,
		Year:  kite.Year,
	}
}

func UnmarshalKite(kite Zephyros.Product) KiteJson {
	return KiteJson{
		Brand: kite.Brand,
		Model: kite.Model,
		Size:  kite.Size,
		Color: kite.Color,
		Year:  kite.Year,
	}
}
