package domain

// Monkey - обезьяна (травоядное)
type Monkey struct {
	*Herbo
}

// NewMonkey создает обезьяну
func NewMonkey(food int, number int, healthy bool, kindnessLevel int) *Monkey {
	return &Monkey{
		Herbo: NewHerbo("Обезьяна", food, number, healthy, kindnessLevel),
	}
}
