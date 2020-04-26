package logerr

import (
	"github.com/bbcloudGroup/gothic/bootstrap"
	"github.com/bbcloudGroup/gothic/di"
	"github.com/kataras/golog"
)

// "fatal"
// "error"
// "warn"
// "info"
// "debug"


func Log(level golog.Level, v ...interface{}) {
	di.Invoke(func (app bootstrap.Bootstrapper) {
		app.Logger().Log(level, v...)
	})
}


func Error(v ...interface{}) {
	Log(golog.ErrorLevel, v...)
}


func Info(v ...interface{}) {
	Log(golog.InfoLevel, v...)
}


func Debug(v ...interface{}) {
	Log(golog.DebugLevel, v...)
}


func Fatal(v ...interface{}) {
	Log(golog.FatalLevel, v...)
}


func Warn(v ...interface{}) {
	Log(golog.WarnLevel, v...)
}

