package config

import "github.com/kataras/iris/v12"

type Configuration interface {
	Register(others *map[string]interface{})
}

type Env struct {
	Iris   		iris.Configuration	`yaml:"Iris"`
	App    		App               	`yaml:"App"`
	Static 		Static            	`yaml:"Static"`
	Log    		Log                	`yaml:"Log"`
	Database	Database			`yaml:"Database"`
}

func (e *Env) Register(others *map[string]interface{}) {
	e.App.Register(others)
	e.Static.Register(others)
	e.Log.Register(others)
	e.Database.Register(others)
}

func DefaultEnv() Env {
	env := Env{
		Iris:	iris.DefaultConfiguration(),
		App:	DefaultApp(),
		Static:	DefaultStatic(),
		Log: 	DefaultLog(),
		Database: DefaultDatabase(),
	}
	return env
}