package services

import (
"moscow-zoo/domain"
"testing"
)

func TestZooAddAnimalHealthy(t *testing.T) {
	clinic := NewVetClinic()
	zoo := NewZoo(clinic)

	monkey := domain.NewMonkey(5, 100, true, 7)
	result := zoo.AddAnimal(monkey)

	if !result {
		t.Error("Expected AddAnimal to return true for healthy animal")
	}
	if zoo.GetAnimalsCount() != 1 {
		t.Errorf("Expected 1 animal, got %d", zoo.GetAnimalsCount())
	}
}

func TestZooAddAnimalUnhealthy(t *testing.T) {
	clinic := NewVetClinic()
	zoo := NewZoo(clinic)

	monkey := domain.NewMonkey(5, 100, false, 7)
	result := zoo.AddAnimal(monkey)

	if result {
		t.Error("Expected AddAnimal to return false for unhealthy animal")
	}
	if zoo.GetAnimalsCount() != 0 {
		t.Errorf("Expected 0 animals, got %d", zoo.GetAnimalsCount())
	}
}

func TestZooGetTotalFoodConsumption(t *testing.T) {
	clinic := NewVetClinic()
	zoo := NewZoo(clinic)

	zoo.AddAnimal(domain.NewMonkey(5, 100, true, 7))
	zoo.AddAnimal(domain.NewRabbit(3, 101, true, 8))
	zoo.AddAnimal(domain.NewTiger(15, 102, true))

	total := zoo.GetTotalFoodConsumption()
	expected := 5 + 3 + 15

	if total != expected {
		t.Errorf("Expected total food %d, got %d", expected, total)
	}
}

func TestZooGetContactZooAnimals(t *testing.T) {
	clinic := NewVetClinic()
	zoo := NewZoo(clinic)

	zoo.AddAnimal(domain.NewMonkey(5, 100, true, 7))
	zoo.AddAnimal(domain.NewRabbit(3, 101, true, 8))
	zoo.AddAnimal(domain.NewMonkey(5, 102, true, 3))
	zoo.AddAnimal(domain.NewTiger(15, 103, true))

	contactAnimals := zoo.GetContactZooAnimals()

	if len(contactAnimals) != 2 {
		t.Errorf("Expected 2 contact zoo animals, got %d", len(contactAnimals))
	}
}

func TestZooAddThing(t *testing.T) {
	clinic := NewVetClinic()
	zoo := NewZoo(clinic)

	table := domain.NewTable(500)
	zoo.AddThing(table)

	things := zoo.GetThings()
	if len(things) != 1 {
		t.Errorf("Expected 1 thing, got %d", len(things))
	}
}

func TestZooGetAllInventory(t *testing.T) {
	clinic := NewVetClinic()
	zoo := NewZoo(clinic)

	zoo.AddAnimal(domain.NewMonkey(5, 100, true, 7))
	zoo.AddAnimal(domain.NewRabbit(3, 101, true, 8))
	zoo.AddThing(domain.NewTable(500))
	zoo.AddThing(domain.NewComputer(501))

	inventory := zoo.GetAllInventory()

	if len(inventory) != 4 {
		t.Errorf("Expected 4 inventory items, got %d", len(inventory))
	}
}

func TestZooGetAnimals(t *testing.T) {
	clinic := NewVetClinic()
	zoo := NewZoo(clinic)

	zoo.AddAnimal(domain.NewMonkey(5, 100, true, 7))
	zoo.AddAnimal(domain.NewTiger(15, 101, true))

	animals := zoo.GetAnimals()

	if len(animals) != 2 {
		t.Errorf("Expected 2 animals, got %d", len(animals))
	}
}

func TestZooEmptyContactZoo(t *testing.T) {
	clinic := NewVetClinic()
	zoo := NewZoo(clinic)

	zoo.AddAnimal(domain.NewTiger(15, 100, true))
	zoo.AddAnimal(domain.NewMonkey(5, 101, true, 3))

	contactAnimals := zoo.GetContactZooAnimals()

	if len(contactAnimals) != 0 {
		t.Errorf("Expected 0 contact zoo animals, got %d", len(contactAnimals))
	}
}
