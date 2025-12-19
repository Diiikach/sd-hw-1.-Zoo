package domain

// Animal - базовый класс животного
type Animal struct {
	name    string
	food    int
	number  int
	healthy bool
}

// NewAnimal создает новое животное
func NewAnimal(name string, food int, number int, healthy bool) *Animal {
	return &Animal{
		name:    name,
		food:    food,
		number:  number,
		healthy: healthy,
	}
}

func (a *Animal) GetFood() int      { return a.food }
func (a *Animal) SetFood(food int)  { a.food = food }
func (a *Animal) GetNumber() int    { return a.number }
func (a *Animal) SetNumber(n int)   { a.number = n }
func (a *Animal) GetName() string   { return a.name }
func (a *Animal) IsHealthy() bool   { return a.healthy }
func (a *Animal) SetHealthy(h bool) { a.healthy = h }
