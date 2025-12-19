package services

import "moscow-zoo/domain"

// IZoo - интерфейс зоопарка
type IZoo interface {
	AddAnimal(animal domain.IAnimal) bool
	GetAnimals() []domain.IAnimal
	GetTotalFoodConsumption() int
	GetAnimalsCount() int
	GetContactZooAnimals() []domain.IHerbivore
	AddThing(thing domain.IInventory)
	GetThings() []domain.IInventory
	GetAllInventory() []domain.IInventory
}

// Zoo - зоопарк
type Zoo struct {
	animals   []domain.IAnimal
	things    []domain.IInventory
	vetClinic IVetClinic
}

// NewZoo создает зоопарк
func NewZoo(vetClinic IVetClinic) *Zoo {
	return &Zoo{
		animals:   make([]domain.IAnimal, 0),
		things:    make([]domain.IInventory, 0),
		vetClinic: vetClinic,
	}
}

// AddAnimal добавляет животное после проверки здоровья
func (z *Zoo) AddAnimal(animal domain.IAnimal) bool {
	if z.vetClinic.CheckHealth(animal) {
		z.animals = append(z.animals, animal)
		return true
	}
	return false
}

// GetAnimals возвращает всех животных
func (z *Zoo) GetAnimals() []domain.IAnimal {
	return z.animals
}

// GetTotalFoodConsumption возвращает общее потребление еды
func (z *Zoo) GetTotalFoodConsumption() int {
	total := 0
	for _, animal := range z.animals {
		total += animal.GetFood()
	}
	return total
}

// GetAnimalsCount возвращает количество животных
func (z *Zoo) GetAnimalsCount() int {
	return len(z.animals)
}

// GetContactZooAnimals возвращает животных для контактного зоопарка
func (z *Zoo) GetContactZooAnimals() []domain.IHerbivore {
	result := make([]domain.IHerbivore, 0)
	for _, animal := range z.animals {
		if herbivore, ok := animal.(domain.IHerbivore); ok {
			if herbivore.CanInteractWithVisitors() {
				result = append(result, herbivore)
			}
		}
	}
	return result
}

// AddThing добавляет вещь
func (z *Zoo) AddThing(thing domain.IInventory) {
	z.things = append(z.things, thing)
}

// GetThings возвращает вещи
func (z *Zoo) GetThings() []domain.IInventory {
	return z.things
}

// GetAllInventory возвращает все инвентаризируемые объекты
func (z *Zoo) GetAllInventory() []domain.IInventory {
	result := make([]domain.IInventory, 0, len(z.animals)+len(z.things))
	for _, animal := range z.animals {
		result = append(result, animal)
	}
	result = append(result, z.things...)
	return result
}
