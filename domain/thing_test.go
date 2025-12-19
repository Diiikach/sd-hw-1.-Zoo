package domain

import "testing"

func TestNewThing(t *testing.T) {
	thing := NewThing("Item", 500)

	if thing.GetName() != "Item" {
		t.Errorf("Expected name 'Item', got '%s'", thing.GetName())
	}
	if thing.GetNumber() != 500 {
		t.Errorf("Expected number 500, got %d", thing.GetNumber())
	}
}

func TestThingSetNumber(t *testing.T) {
	thing := NewThing("Item", 500)
	thing.SetNumber(600)

	if thing.GetNumber() != 600 {
		t.Errorf("Expected number 600, got %d", thing.GetNumber())
	}
}

func TestNewTable(t *testing.T) {
	table := NewTable(700)

	if table.GetName() != "Стол" {
		t.Errorf("Expected name 'Стол', got '%s'", table.GetName())
	}
	if table.GetNumber() != 700 {
		t.Errorf("Expected number 700, got %d", table.GetNumber())
	}
}

func TestNewComputer(t *testing.T) {
	computer := NewComputer(800)

	if computer.GetName() != "Компьютер" {
		t.Errorf("Expected name 'Компьютер', got '%s'", computer.GetName())
	}
	if computer.GetNumber() != 800 {
		t.Errorf("Expected number 800, got %d", computer.GetNumber())
	}
}
