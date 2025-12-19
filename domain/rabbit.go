package domain

// Rabbit - кролик (травоядное)
type Rabbit struct {
	*Herbo
}

// NewRabbit создает кролика
func NewRabbit(food int, number int, healthy bool, kindnessLevel int) *Rabbit {
	return &Rabbit{
		Herbo: NewHerbo("Кролик", food, number, healthy, kindnessLevel),
	}
}
