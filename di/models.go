package di

import "go.uber.org/dig"

type Container struct {
	*dig.Container
}

func newContainer() *Container {
	return &Container{dig.New()}
}

func (c *Container) Register(constructor interface{}) {
	err := c.Provide(constructor)
	if err != nil {
		panic(err)
	}
}
