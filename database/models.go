package database

type Driver int32

type Database int32

const (
	GORM Driver = iota
	REDIS
	MEMCACHE

	None Database = iota
	MySQL
	PostgreSQL
	Sqlite3
	Redis
	Memcache
)

func (c Database) String() string {
	switch c {
	case None:
		return "None"
	case MySQL:
		return "mysql"
	case PostgreSQL:
		return "postgres"
	case Sqlite3:
		return "sqlite3"
	case Redis:
		return "redis"
	case Memcache:
		return "memcache"
	}
	return ""
}

type Databases map[string]Connection

type Connection struct {
	database Database
	driver Driver
	conn interface{}
	connString	string
	others map[string]interface{}
}

