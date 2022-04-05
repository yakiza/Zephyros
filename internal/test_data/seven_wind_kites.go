package test_data

import "github.com/yakiza/Zephyros"

func VorasTestKite() Zephyros.Kite {
	return Zephyros.Kite{
		Brand: "Voras",
		Model: "XX",
		Size:  12,
		Color: "red",
		Year:  2022,
	}
}

func AeolianTestKite() Zephyros.Kite {
	return Zephyros.Kite{
		Brand: "Aeolian",
		Model: "FX",
		Size:  12,
		Color: "yellow",
		Year:  2022,
	}
}

func AllSevenWindKites() []Zephyros.Kite {
	return []Zephyros.Kite{
		VorasTestKite(),
		AeolianTestKite(),
	}
}