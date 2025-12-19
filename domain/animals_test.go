package domain

import "testing"

func TestNewMonkey(t *testing.T) {
	monkey := NewMonkey(5, 100, true, 6)

	if monkey.GetName() != "Обезьяна" {
		t.Errorf("Expected name 'Обезьяна', got '%s'", monkey.GetName())
	}
	if monkey.GetFood() != 5 {
		t.Errorf("Expected food 5, got %d", monkey.GetFood())
	}
	if monkey.GetNumber() != 100 {
		t.Errorf("Expected number 100, got %d", monkey.GetNumber())
	}
	if !monkey.IsHealthy() {
		t.Error("Expected healthy true")
	}
	if monkey.GetKindnessLevel() != 6 {
		t.Errorf("Expected kindness 6, got %d", monkey.GetKindnessLevel())
	}
}

func TestNewRabbit(t *testing.T) {
	rabbit := NewRabbit(2, 200, true, 8)

	if rabbit.GetName() != "Кролик" {
		t.Errorf("Expected name 'Кролик', got '%s'", rabbit.GetName())
	}
	if rabbit.GetFood() != 2 {
		t.Errorf("Expected food 2, got %d", rabbit.GetFood())
	}
	if rabbit.GetKindnessLevel() != 8 {
		t.Errorf("Expected kindness 8, got %d", rabbit.GetKindnessLevel())
	}
}

func TestNewTiger(t *testing.T) {
	tiger := NewTiger(15, 300, true)

	if tiger.GetName() != "Тигр" {
		t.Errorf("Expected name 'Тигр', got '%s'", tiger.GetName())
	}
	if tiger.GetFood() != 15 {
		t.Errorf("Expected food 15, got %d", tiger.GetFood())
	}
	if tiger.GetNumber() != 300 {
		t.Errorf("Expected number 300, got %d", tiger.GetNumber())
	}
}

func TestNewWolf(t *testing.T) {
	wolf := NewWolf(10, 400, false)

	if wolf.GetName() != "Волк" {
		t.Errorf("Expected name 'Волк', got '%s'", wolf.GetName())
	}
	if wolf.GetFood() != 10 {
		t.Errorf("Expected food 10, got %d", wolf.GetFood())
	}
	if wolf.IsHealthy() {
		t.Error("Expected healthy false")
	}
}
