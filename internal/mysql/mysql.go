package mysql

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql" //Es el conector para mysql

	"github.com/DasJalapa/reportes-salud/internal/lib"
)

var (
	once sync.Once
	db   *sql.DB
	err  error
)

//Connect is a function that permited the connection to mysql
func Connect() *sql.DB {
	c := lib.Config()
	user := c.USERDB
	password := c.PASSWORDDB
	server := c.HOSTDB
	database := c.DATABASE

	once.Do(func() {
		db, err = sql.Open("mysql", user+":"+password+"@tcp("+server+")/"+database)
		if err != nil {
			log.Println(err.Error())
		}
	})

	return db
}
