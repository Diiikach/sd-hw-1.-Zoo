package domain

// Tiger - тигр (хищник)
type Tiger struct {
	*Predator
}

// NewTiger создает тигра
func NewTiger(food int, number int, healthy bool) *Tiger {
	return &Tiger{
		Predator: NewPredator("Тигр", food, number, healthy),
	}
}
