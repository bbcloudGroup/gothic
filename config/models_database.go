package config

import "github.com/kataras/iris/v12/context"

const (
	DatabaseSQL 		= "DatabaseSQL"
	DatabaseCache 		= "DatabaseCache"
)

type Database struct {
	SQL 	[] SQL 		`yaml:"SQL"`
	Cache 	[] Cache 	`yaml:"Cache"`


	//MySQL		[] MySQL		`yaml:"MySQL"`
	//PostgreSQL	[] PostgreSQL 	`yaml:"PostgreSQL"`
	//Sqlite3		[] Sqlite3		`yaml:"Sqlite3"`
}


type SQL struct {
	Name 		string 	`yaml:"Name"`
	Driver		string	`yaml:"Driver"`
	Host 		string 	`yaml:"Host"`
	Port		int		`yaml:"Port"`
	Dbname 		string 	`yaml:"Dbname"`
	Username 	string 	`yaml:"Username"`
	Password 	string	`yaml:"Password"`
	Sslmode		bool	`yaml:"Sslmode"`
	Charset		string	`yaml:"Charset"`
	MaxIdle		int		`yaml:"MaxIdle"`
	MaxOpen		int		`yaml:"MaxOpen"`
	Log			bool	`yaml:"Log"`
}


type Cache struct {
	Name		string 	`yaml:"Name"`
	Driver		string 	`yaml:"Driver"`
	Host 		string 	`yaml:"Host"`
	Port		int		`yaml:"Port"`
	Password 	string	`yaml:"Password"`
	Db			int		`yaml:"Db"`
	Ssl 		bool	`yaml:"Ssl"`
}

func (db *Database) Register(others *map[string]interface{}) {
	(*others)[DatabaseSQL]		= db.SQL
	(*others)[DatabaseCache]	= db.Cache
}

func GetDatabase(app context.Application) Database {
	return Database{
		SQL:      	Get(app, DatabaseSQL).([] SQL),
		Cache: 		Get(app, DatabaseCache).([] Cache),
	}
}

func DefaultDatabase() Database {
	return Database{
		SQL:	nil,
		Cache: 	nil,
	}
}
