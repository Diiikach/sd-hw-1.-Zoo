package domain

// Herbo - базовый класс травоядного
type Herbo struct {
	*Animal
	kindnessLevel int
}

// NewHerbo создает травоядное
func NewHerbo(name string, food int, number int, healthy bool, kindnessLevel int) *Herbo {
	if kindnessLevel < 0 {
		kindnessLevel = 0
	}
	if kindnessLevel > 10 {
		kindnessLevel = 10
	}
	return &Herbo{
		Animal:        NewAnimal(name, food, number, healthy),
		kindnessLevel: kindnessLevel,
	}
}

func (h *Herbo) GetKindnessLevel() int { return h.kindnessLevel }

func (h *Herbo) SetKindnessLevel(level int) {
	if level < 0 {
		level = 0
	}
	if level > 10 {
		level = 10
	}
	h.kindnessLevel = level
}

func (h *Herbo) CanInteractWithVisitors() bool {
	return h.kindnessLevel > 5
}
