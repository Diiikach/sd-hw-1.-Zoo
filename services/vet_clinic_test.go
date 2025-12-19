package services

import (
"moscow-zoo/domain"
"testing"
)

func TestVetClinicCheckHealthHealthy(t *testing.T) {
	clinic := NewVetClinic()
	animal := domain.NewMonkey(5, 100, true, 7)

	if !clinic.CheckHealth(animal) {
		t.Error("Expected CheckHealth to return true for healthy animal")
	}
}

func TestVetClinicCheckHealthUnhealthy(t *testing.T) {
	clinic := NewVetClinic()
	animal := domain.NewMonkey(5, 100, false, 7)

	if clinic.CheckHealth(animal) {
		t.Error("Expected CheckHealth to return false for unhealthy animal")
	}
}
