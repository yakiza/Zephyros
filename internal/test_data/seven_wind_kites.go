package test_data

import "github.com/yakiza/Zephyros"

func VorasTestKite() Zephyros.Product {
	return Zephyros.Product{
		Name:        "Kite",
		Description: "Lorem Ipsum",
		Brand:       "Voras",
		Price:       1280,
		Currency:    "GBP",
	}
}

func AeolianTestKite() Zephyros.Product {
	return Zephyros.Product{
		Name:        "Kite",
		Description: "Lorem Ipsum",
		Brand:       "Aeolian",
		Price:       1480,
		Currency:    "EURO",
	}
}

func AllSevenWindKites() []Zephyros.Product {
	return []Zephyros.Product{
		VorasTestKite(),
		AeolianTestKite(),
	}
}
