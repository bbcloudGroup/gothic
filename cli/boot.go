package cli

import (
	"github.com/kataras/iris/v12"
	"github.com/robfig/cron/v3"
	"log"
)

type Schedule func(cron *cron.Cron)

func BootStrap(schedule Schedule) iris.Configurator {

	return func (app *iris.Application) {
		CronLogger := cron.PrintfLogger(log.New(app.Logger().Printer.Output, "cron: ", log.LstdFlags))
		c := cron.New(cron.WithSeconds(), cron.WithLogger(CronLogger))
		schedule(c)
		c.Start()
	}
}