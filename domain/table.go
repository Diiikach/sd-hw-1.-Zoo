package domain

// Table - стол
type Table struct {
	*Thing
}

// NewTable создает стол
func NewTable(number int) *Table {
	return &Table{
		Thing: NewThing("Стол", number),
	}
}
