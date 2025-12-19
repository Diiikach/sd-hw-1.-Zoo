package domain

import "testing"

func TestNewHerbo(t *testing.T) {
	herbo := NewHerbo("Herbo", 3, 50, true, 7)

	if herbo.GetName() != "Herbo" {
		t.Errorf("Expected name 'Herbo', got '%s'", herbo.GetName())
	}
	if herbo.GetKindnessLevel() != 7 {
		t.Errorf("Expected kindness 7, got %d", herbo.GetKindnessLevel())
	}
}

func TestHerboKindnessLevelBounds(t *testing.T) {
	herbo := NewHerbo("Test", 3, 50, true, -5)
	if herbo.GetKindnessLevel() != 0 {
		t.Errorf("Expected kindness 0, got %d", herbo.GetKindnessLevel())
	}

	herbo2 := NewHerbo("Test2", 3, 50, true, 15)
	if herbo2.GetKindnessLevel() != 10 {
		t.Errorf("Expected kindness 10, got %d", herbo2.GetKindnessLevel())
	}
}

func TestHerboSetKindnessLevel(t *testing.T) {
	herbo := NewHerbo("Test", 3, 50, true, 5)

	herbo.SetKindnessLevel(8)
	if herbo.GetKindnessLevel() != 8 {
		t.Errorf("Expected kindness 8, got %d", herbo.GetKindnessLevel())
	}

	herbo.SetKindnessLevel(-1)
	if herbo.GetKindnessLevel() != 0 {
		t.Errorf("Expected kindness 0, got %d", herbo.GetKindnessLevel())
	}

	herbo.SetKindnessLevel(11)
	if herbo.GetKindnessLevel() != 10 {
		t.Errorf("Expected kindness 10, got %d", herbo.GetKindnessLevel())
	}
}

func TestCanInteractWithVisitors(t *testing.T) {
	herbo1 := NewHerbo("Kind", 3, 50, true, 7)
	if !herbo1.CanInteractWithVisitors() {
		t.Error("Expected CanInteractWithVisitors true for kindness 7")
	}

	herbo2 := NewHerbo("Neutral", 3, 50, true, 5)
	if herbo2.CanInteractWithVisitors() {
		t.Error("Expected CanInteractWithVisitors false for kindness 5")
	}

	herbo3 := NewHerbo("Mean", 3, 50, true, 3)
	if herbo3.CanInteractWithVisitors() {
		t.Error("Expected CanInteractWithVisitors false for kindness 3")
	}
}
