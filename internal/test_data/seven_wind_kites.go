package test_data

import "github.com/yakiza/Zephyros"

func VorasTestKite() Zephyros.Product {
	return Zephyros.Product{
		Category: "Voras",
		Model:    "XX",
		Size:     12,
		Color:    "red",
		Season:   2022,
	}
}

func AeolianTestKite() Zephyros.Product {
	return Zephyros.Product{
		Category: "Aeolian",
		Model:    "FX",
		Size:     12,
		Color:    "yellow",
		Season:   2022,
	}
}

func AllSevenWindKites() []Zephyros.Product {
	return []Zephyros.Product{
		VorasTestKite(),
		AeolianTestKite(),
	}
}
