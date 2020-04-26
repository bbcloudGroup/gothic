package database

import (
	"fmt"
	"github.com/bbcloudGroup/gothic/config"
	"github.com/kataras/iris/v12"
)


var _db = make(Databases)
var _cache = make(Databases)

const (
	Default = "Default"
)


func BootStrap() iris.Configurator {

	return func (app *iris.Application) {

		database := config.GetDatabase(app)

		for _, sql := range database.SQL {
			createSQLConnection(sql)
		}

		for _, cache := range database.Cache {
			createCacheConnection(cache)
		}


	}

}



func createCacheConnection(cache config.Cache) {
	var connection Connection
	switch cache.Driver {
	case Memcache.String():
		connection = Connection{
			database:   Memcache,
			driver:     MEMCACHE,
			conn:       nil,
			connString: fmt.Sprintf("%s:%s", cache.Host, cache.Port),
			others:     nil,
		}
		break
	default:
		proto := "redis"
		if cache.Ssl {
			proto = "rediss"
		}
		connection = Connection{
			database:   Redis,
			driver:     REDIS,
			conn:       nil,
			connString: fmt.Sprintf("%s://%s@%s:%d/%d", proto, cache.Password, cache.Host, cache.Port, cache.Db),
			others: 	map[string]interface{}{
				"addr": fmt.Sprintf("%s:%d", cache.Host, cache.Port),
				"password": cache.Password,
				"db": cache.Db,
			},
		}
		break
	}
	setConnection(&_cache, cache.Name, connection)
}

func createSQLConnection(sql config.SQL) {
	var database Database
	var connString string
	switch sql.Driver {
	case PostgreSQL.String():
		ssl := "disable"
		if sql.Sslmode {
			ssl = "enable"
		}
		database = PostgreSQL
		connString = fmt.Sprintf(
					"host=%s user=%s dbname=%s sslmode=%s password=%s",
					sql.Host, sql.Username, sql.Dbname, ssl, sql.Password)
		break
	case Sqlite3.String():
		database = Sqlite3
		connString = sql.Host
		break
	default:
		database = MySQL
		connString = fmt.Sprintf(
					"%s:%s@(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
					sql.Username, sql.Password, sql.Host, sql.Port, sql.Dbname, sql.Charset)
		break
	}

	setConnection(&_db, sql.Name, Connection{
		database: 	database,
		driver:   	GORM,
		conn:		nil,
		connString: connString,
		others: map[string]interface{}{
			"max_idle": sql.MaxIdle,
			"max_open": sql.MaxOpen,
			"log": 		sql.Log,
		},
	})
}


func setConnection(dbs *Databases, name string, connection Connection) {
	if len(name) == 0 {
		name = Default
	}
	(*dbs)[name] = connection
}
