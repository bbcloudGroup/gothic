package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func BootStrap(routeMap Map) iris.Configurator {
	return func (app *iris.Application) {

		routes := routeMap(app)
		for path, route := range routes {
			m := mvc.New(app.Party(path))
			m.Router.Use(route.Middleware...)
			if len(route.Terminate) > 0 {
				m.Router.SetExecutionRules(iris.ExecutionRules{
					Done: iris.ExecutionOptions{Force: true},
				})
				m.Router.Done(route.Terminate...)
			}
			//container.DI.Invoke()
			route.Controller(m)
		}
	}
}