package di

import "moscow-zoo/services"

// Container - DI контейнер
type Container struct {
	vetClinic services.IVetClinic
	zoo       services.IZoo
}

// NewContainer создает контейнер
func NewContainer() *Container {
	c := &Container{}
	c.init()
	return c
}

func (c *Container) init() {
	c.vetClinic = services.NewVetClinic()
	c.zoo = services.NewZoo(c.vetClinic)
}

// GetVetClinic возвращает ветклинику
func (c *Container) GetVetClinic() services.IVetClinic {
	return c.vetClinic
}

// GetZoo возвращает зоопарк
func (c *Container) GetZoo() services.IZoo {
	return c.zoo
}
