package config

import (
	"github.com/kataras/iris/v12/context"
)

const (
	LogLevel	= "LogLevel"
	LogEnabled 	= "LogEnabled"
	LogFile 	= "LogFile"
	LogStatus	= "LogStatus"
	LogIP		= "LogIP"
	LogMethod	= "LogMethod"
	LogPath		= "LogPath"
	LogQuery	= "LogQuery"
	LogColumns	= "LogColumns"
	LogMessageContextKeys	= "LogMessageContextKeys"
	LogMessageHeaderKeys	= "LogMessageHeaderKeys"
)

type Log struct {
	Level	string	`yaml:"Level"`
	Enabled	bool	`yaml:"Enabled"`
	File 	string 	`yaml:"File"`
	Status 	bool	`yaml:"Status"`
	IP 		bool	`yaml:"IP"`
	Method 	bool	`yaml:"Method"`
	Path 	bool	`yaml:"Path"`
	Query 	bool	`yaml:"Query"`
	Columns bool	`yaml:"Columns"`
	MessageContextKeys	[]string `yaml:"MessageContextKeys"`
	MessageHeaderKeys	[]string `yaml:"MessageHeaderKeys"`
}

func (l *Log) Register(others *map[string]interface{}) {
	(*others)[LogLevel] = l.Level
	(*others)[LogEnabled] = l.Enabled
	(*others)[LogFile] = l.File
	(*others)[LogStatus] = l.Status
	(*others)[LogIP] = l.IP
	(*others)[LogMethod] = l.Method
	(*others)[LogPath] = l.Path
	(*others)[LogQuery] = l.Query
	(*others)[LogColumns] = l.Columns
	(*others)[LogMessageContextKeys] = l.MessageContextKeys
	(*others)[LogMessageHeaderKeys] = l.MessageHeaderKeys
}

func GetLog(app context.Application) Log {
	return Log {
		Level:		GetString(app, LogLevel),
		Enabled:	GetBool(app, LogEnabled),
		File:		GetString(app, LogFile),
		Status:  	GetBool(app, LogStatus),
		IP:      	GetBool(app, LogIP),
		Method:  	GetBool(app, LogMethod),
		Path:    	GetBool(app, LogPath),
		Query:   	GetBool(app, LogQuery),
		Columns: 	GetBool(app, LogColumns),
		MessageContextKeys: Get(app, LogMessageContextKeys).([]string),
		MessageHeaderKeys:  Get(app, LogMessageHeaderKeys).([]string),
	}
}

func DefaultLog() Log {
	return Log{
		Level:	 "info",
		Enabled: true,
		File:    "./logs/",
		Status:  true,
		IP:      true,
		Method:  true,
		Path:    true,
		Query:   false,
		Columns: false,
		MessageContextKeys: []string{"logger_message"},
		MessageHeaderKeys:  []string{"User-Agent"},
	}
}
