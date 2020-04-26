package config

import (
	"github.com/kataras/iris/v12/context"
)

type Property func(env string) map[string]interface{}


func Get(app context.Application, key string) interface{} {
	v, ok := app.ConfigurationReadOnly().GetOther()[key]
	if ok == true {
		return v
	}
	return nil
}


func GetString(app context.Application, key string) string {
	return Get(app, key).(string)
}

func GetInt(app context.Application,key string) int64 {
	return Get(app, key).(int64)
}

func GetBool(app context.Application,key string) bool {
	return Get(app, key).(bool)
}


















