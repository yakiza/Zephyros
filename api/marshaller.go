package api

import "github.com/yakiza/Zephyros"

type KiteJson struct {
	Brand string  `json:"brand"`
	Model string  `json:"model"`
	Size  float32 `json:"size"`
	Color string  `json:"color"`
	Year  int     `json:"year"`
}

func MarshallKite(kite KiteJson) Zephyros.Kite {
	return Zephyros.Kite{
		Brand: kite.Brand,
		Model: kite.Model,
		Size:  kite.Size,
		Color: kite.Color,
		Year:  kite.Year,
	}
}

func UnmarshalKite(kite Zephyros.Kite) KiteJson {
	return KiteJson{
		Brand: kite.Brand,
		Model: kite.Model,
		Size:  kite.Size,
		Color: kite.Color,
		Year:  kite.Year,
	}
}
