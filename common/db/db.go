package db

import (
	"courses/common/streams"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	"os"
	"time"
)

var (
	dsn string
	db  *dbr.Connection
)

func Init() error {
	var (
		err    error
		stream dbr.EventReceiver = &dbr.NullEventReceiver{}
	)
	dsn = os.Getenv("MYSQL_DSN")
	if dsn == "" {
		dsn = "username:password@tcp(host:port)/dbname"
	}
	if os.Getenv("DB_SHOW_QUERIES") != "no" {
		stream = streams.GetStdOutStream()
	}
	db, err = dbr.Open("mysql", dsn, stream)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	db.SetConnMaxLifetime(time.Minute)
	db.SetMaxOpenConns(30)
	db.SetMaxIdleConns(10)
	return nil
}

func GetDB() *dbr.Connection {
	return db
}

func GetSession() *dbr.Session {
	return GetDB().NewSession(nil)
}
