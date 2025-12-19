package domain

// Computer - компьютер
type Computer struct {
	*Thing
}

// NewComputer создает компьютер
func NewComputer(number int) *Computer {
	return &Computer{
		Thing: NewThing("Компьютер", number),
	}
}
