package domain

// Wolf - волк (хищник)
type Wolf struct {
	*Predator
}

// NewWolf создает волка
func NewWolf(food int, number int, healthy bool) *Wolf {
	return &Wolf{
		Predator: NewPredator("Волк", food, number, healthy),
	}
}
