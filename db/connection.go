package db

import (
	"github.com/go-pg/pg"
	"net"
	"sync"
)

var db *pg.DB
var once sync.Once

type ConnectionSettings struct {
	DbUser     string
	DbPassword string
	DbHost     string
	DbPort     string
	DbName     string
}

func GetConnectionInstance(cs ConnectionSettings) *pg.DB {
	once.Do(func() {
		db = pg.Connect(&pg.Options{
			User:     cs.DbUser,
			Password: cs.DbPassword,
			Addr:     net.JoinHostPort(cs.DbHost, cs.DbPort),
			Database: cs.DbName,
		})
		_, err := db.Exec("SELECT 1;")

		if err != nil {
			panic(err)
		}
	})
	return db
}
