package static

import (
	"github.com/bbcloudGroup/gothic/config"
	"github.com/kataras/iris/v12"
)


func BootStrap(opts ...iris.DirOptions) iris.Configurator {

	return func(app *iris.Application) {

		static := config.GetStatic(app)
		app.Favicon(static.Assets + static.Favicon)

		if len(opts) == 0 {
			opt := iris.DirOptions{
				IndexName: static.IndexName,
				Gzip:	static.Gzip,
				ShowList:	static.ShowList,
			}
			opts = append(opts, opt)
		}

		app.HandleDir(static.RequestPath, static.Assets, opts...)

	}
}