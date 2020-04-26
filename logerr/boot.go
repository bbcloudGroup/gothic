package logerr

import (
	"github.com/bbcloudGroup/gothic/config"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/logger"
	"io"
	"os"
	"path"
	"time"
)

func BootStrap(errorHandler context.Handler, cfgs ...logger.Config) iris.Configurator {
	return func (app *iris.Application) {

		log := config.GetLog(app)

		app.Logger().SetLevel(log.Level)

		if log.Enabled {
			f := newLogFile(log.File)
			app.Logger().SetOutput(io.MultiWriter(f, os.Stdout))
		}

		app.Use(recoverHandler())
		app.Use(logHandler(log, cfgs...))

		app.OnAnyErrorCode(errorHandler)

	}
}

func newLogFile(dir string) *os.File {
	today := time.Now().Format("2006-01-02")
	filename := path.Join(dir, today + ".log")
	f, err := os.OpenFile(filename, os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return f
}


