package domain

import "testing"

func TestNewAnimal(t *testing.T) {
	animal := NewAnimal("Test", 5, 100, true)

	if animal.GetName() != "Test" {
		t.Errorf("Expected name 'Test', got '%s'", animal.GetName())
	}
	if animal.GetFood() != 5 {
		t.Errorf("Expected food 5, got %d", animal.GetFood())
	}
	if animal.GetNumber() != 100 {
		t.Errorf("Expected number 100, got %d", animal.GetNumber())
	}
	if !animal.IsHealthy() {
		t.Error("Expected healthy true")
	}
}

func TestAnimalSetters(t *testing.T) {
	animal := NewAnimal("Test", 5, 100, true)

	animal.SetFood(10)
	if animal.GetFood() != 10 {
		t.Errorf("Expected food 10, got %d", animal.GetFood())
	}

	animal.SetNumber(200)
	if animal.GetNumber() != 200 {
		t.Errorf("Expected number 200, got %d", animal.GetNumber())
	}

	animal.SetHealthy(false)
	if animal.IsHealthy() {
		t.Error("Expected healthy false")
	}
}
