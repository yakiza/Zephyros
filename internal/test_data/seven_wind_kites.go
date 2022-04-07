package test_data

import "github.com/yakiza/Zephyros"

func VorasTestKite() Zephyros.Product {
	return Zephyros.Product{
		Brand: "Voras",
		Model: "XX",
		Size:  12,
		Color: "red",
		Year:  2022,
	}
}

func AeolianTestKite() Zephyros.Product {
	return Zephyros.Product{
		Brand: "Aeolian",
		Model: "FX",
		Size:  12,
		Color: "yellow",
		Year:  2022,
	}
}

func AllSevenWindKites() []Zephyros.Product {
	return []Zephyros.Product{
		VorasTestKite(),
		AeolianTestKite(),
	}
}
