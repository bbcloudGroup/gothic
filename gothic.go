package gothic

import (

	"github.com/bbcloudGroup/gothic/bootstrap"
	"github.com/bbcloudGroup/gothic/config"
	"github.com/bbcloudGroup/gothic/di"
)

func Config(key string) interface{} {
	var value interface{}
	di.Invoke(func (app bootstrap.Bootstrapper) {
		value = config.Get(&app, key)
	})
	return value
}

