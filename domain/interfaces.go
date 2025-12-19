package domain

// IAlive - интерфейс для живых существ
type IAlive interface {
	GetFood() int
	SetFood(food int)
}

// IInventory - интерфейс для инвентаризации
type IInventory interface {
	GetNumber() int
	SetNumber(number int)
	GetName() string
}

// IAnimal - интерфейс животного
type IAnimal interface {
	IAlive
	IInventory
	IsHealthy() bool
	SetHealthy(healthy bool)
}

// IHerbivore - интерфейс травоядного
type IHerbivore interface {
	IAnimal
	GetKindnessLevel() int
	SetKindnessLevel(level int)
	CanInteractWithVisitors() bool
}

// IPredator - интерфейс хищника
type IPredator interface {
	IAnimal
}
