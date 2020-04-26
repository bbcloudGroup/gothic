package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/mvc"
)

type Map func(app *iris.Application) map[string]Route

type Route struct {
	Controller func(m *mvc.Application)
	Middleware	[] context.Handler
	Terminate	[] context.Handler
}

func New(controller func(m *mvc.Application)) Route {
	return Route{
		Controller: controller,
		Middleware: [] context.Handler {},
		Terminate: 	[] context.Handler {},
	}
}

func (r Route) WithMiddleware(middleware ...context.Handler) Route {
	r.Middleware = append(r.Middleware, middleware...)
	return r
}

func (r Route) WithTerminate(terminate ...context.Handler) Route {
	r.Terminate = append(r.Terminate, terminate...)
	return r
}



