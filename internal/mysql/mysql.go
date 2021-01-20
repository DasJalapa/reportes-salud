package mysql

import (
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"io/ioutil"
	"log"
	"sync"

	"github.com/go-sql-driver/mysql"
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

	rootCert := x509.NewCertPool()
	pem, err := ioutil.ReadFile("./internal/certificates/mysqlCertified.pem")
	if err != nil {
		log.Fatal(err)
	}

	if ok := rootCert.AppendCertsFromPEM(pem); !ok {
		log.Fatal(err)
	}

	mysql.RegisterTLSConfig("custom", &tls.Config{
		ServerName: server,
		RootCAs:    rootCert,
	})

	once.Do(func() {
		db, err = sql.Open("mysql", user+":"+password+"@tcp("+server+")/"+database+"?tls=custom")
		if err != nil {
			log.Println(err.Error())
		}
	})

	return db
}
