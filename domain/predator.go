package domain

// Predator - базовый класс хищника
type Predator struct {
	*Animal
}

// NewPredator создает хищника
func NewPredator(name string, food int, number int, healthy bool) *Predator {
	return &Predator{
		Animal: NewAnimal(name, food, number, healthy),
	}
}
