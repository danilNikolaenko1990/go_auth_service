package db

import (
	"github.com/go-pg/pg"
	"net"
	"os"
	"sync"
)

func init() {
	//todo remove
	os.Setenv("DB_USER", "admin")
	os.Setenv("DB_PASSWORD", "admin")
	os.Setenv("DB_NAME", "postgres")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
}

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
