package config

import "github.com/kataras/iris/v12/context"

const (
	AppName		=	"AppName"
	AppEnv		=	"AppEnv"
	AppUrl		=	"AppUrl"
	AppPort		=	"AppPort"
)

type App struct {
	AppName		string	`yaml:"AppName"`
	AppEnv		string	`yaml:"AppEnv"`
	AppUrl		string	`yaml:"AppUrl"`
	AppPort		string	`yaml:"AppPort"`
}

func (app *App) Register(others *map[string]interface{}) {
	(*others)[AppName] = app.AppName
	(*others)[AppEnv] = app.AppEnv
	(*others)[AppUrl] = app.AppUrl
	(*others)[AppPort] = app.AppPort
}

func GetApp(app context.Application) App {
	return App{
		AppName:	GetString(app, AppName),
		AppEnv: 	GetString(app, AppEnv),
		AppUrl: 	GetString(app, AppUrl),
		AppPort: 	GetString(app, AppPort),
	}
}

func DefaultApp() App {
	return App{
		AppName:  "gothic_app",
		AppEnv:   "production",
		AppUrl:   "http://localhost",
		AppPort:  "8080",
	}
}
