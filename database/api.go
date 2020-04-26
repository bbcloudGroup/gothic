package database

import (
	"errors"
	"fmt"
	"github.com/bbcloudGroup/gothic/bootstrap"
	"github.com/bbcloudGroup/gothic/di"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)


func G(name ...string) (*gorm.DB) {
	db, err := SqlConn(name...)
	if err != nil {
		panic(err)
	}
	return db.(*gorm.DB)
}

func R(name ...string) (*redis.Client) {
	db, err := CacheConn(name...)
	if err != nil {
		panic(err)
	}
	return db.(*redis.Client)
}

func M(name ...string) (*memcache.Client) {
	db, err := CacheConn(name...)
	if err != nil {
		panic(err)
	}
	return db.(*memcache.Client)
}


func connect(database *Databases, name ...string) (Connection, error) {
	conn_name := "default"
	if len(name) > 0 {
		conn_name = name[0]
	}

	conn, ok := (*database)[conn_name]
	if !ok {
		return Connection{}, errors.New(fmt.Sprintf("connection [%s] does not exists", conn_name))
	}

	return conn, nil
}


func SqlConn(name ...string) (interface{}, error) {

	conn, err := connect(&_db, name...)
	if err != nil {
		return nil, err
	}
	if conn.conn != nil {
		return conn.conn, nil
	}

	switch conn.driver {
	case GORM:
		db, err := gorm.Open(conn.database.String(), conn.connString)
		if err != nil {
			return nil, err
		}
		db.SingularTable(true)
		conn.conn = db
		if conn.others["log"].(bool) {
			db.LogMode(true)
		}
		if max_idle := conn.others["max_idle"].(int); max_idle != 2 {
			db.DB().SetMaxIdleConns(max_idle)
		}
		if max_open := conn.others["max_open"].(int); max_open > 0 {
			db.DB().SetMaxOpenConns(max_open)
		}
		di.Invoke(func (app bootstrap.Bootstrapper) {
			db.SetLogger(log.New(app.Logger().Printer.Output, "gorm", 0))
		})

		return db, nil
	default:
		return nil, errors.New("no driver to connection")
	}
}



func CacheConn(name ...string) (interface{}, error) {
	conn, err := connect(&_cache, name...)
	if err != nil {
		return nil, err
	}
	if conn.conn != nil {
		return conn.conn, nil
	}

	switch conn.driver {
	case REDIS:
		db := redis.NewClient(&redis.Options{
			Addr:     conn.others["addr"].(string),
			Password: conn.others["password"].(string),
			DB:       conn.others["db"].(int),
		})
		return db, nil
	case MEMCACHE:
		db := memcache.New(conn.connString)
		return db, nil
	default:
		return nil, errors.New("no driver to connection")
	}
}
