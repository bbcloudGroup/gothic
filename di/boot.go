package di

import (
	"github.com/bbcloudGroup/gothic/bootstrap"
	"github.com/kataras/iris/v12"
)

var di = newContainer()

type Provider func (container Container)

func BootStrap(providers ...Provider) iris.Configurator {
	return func (app *iris.Application) {
		for _, provider := range providers {
			provider(*di)
		}
		di.Register(bootstrap.App)
	}
}
