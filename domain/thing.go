package domain

// Thing - базовый класс вещи
type Thing struct {
	name   string
	number int
}

// NewThing создает вещь
func NewThing(name string, number int) *Thing {
	return &Thing{
		name:   name,
		number: number,
	}
}

func (t *Thing) GetName() string { return t.name }
func (t *Thing) GetNumber() int  { return t.number }
func (t *Thing) SetNumber(n int) { t.number = n }
