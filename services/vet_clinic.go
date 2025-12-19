package services

import "moscow-zoo/domain"

// IVetClinic - интерфейс ветеринарной клиники
type IVetClinic interface {
	CheckHealth(animal domain.IAnimal) bool
}

// VetClinic - ветеринарная клиника
type VetClinic struct{}

// NewVetClinic создает ветеринарную клинику
func NewVetClinic() *VetClinic {
	return &VetClinic{}
}

// CheckHealth проверяет здоровье животного
func (v *VetClinic) CheckHealth(animal domain.IAnimal) bool {
	return animal.IsHealthy()
}
