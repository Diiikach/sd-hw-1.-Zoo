package di

import "testing"

func TestNewContainer(t *testing.T) {
	container := NewContainer()

	if container == nil {
		t.Error("Expected container to be not nil")
	}
}

func TestContainerGetVetClinic(t *testing.T) {
	container := NewContainer()
	clinic := container.GetVetClinic()

	if clinic == nil {
		t.Error("Expected VetClinic to be not nil")
	}
}

func TestContainerGetZoo(t *testing.T) {
	container := NewContainer()
	zoo := container.GetZoo()

	if zoo == nil {
		t.Error("Expected Zoo to be not nil")
	}
}
