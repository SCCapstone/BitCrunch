package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "root"
	password = "mypassSQL1@@bit"
	hostname = "127.0.0.1:3306"
)

var err error

type dbase struct {
	sqldb  *sql.DB
	opened bool
}

func New() *dbase {
	return &dbase{
		nil,
		false,
	}
}

func (db *dbase) Open() {
	db.sqldb, err = sql.Open("mysql", dsn(""))
	if err != nil {
		log.Panic("Could not open mysql!")
	}
	db.opened = true
}

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func (db *dbase) CreateOrOpenDB(dbaseName string) {
	if !db.opened {
		db.Open() // will crash if this fails
	}
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.sqldb.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbaseName)
	if err != nil {
		log.Printf("Error %s when creating DB\n", err)
		return
	}
	no, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when fetching rows", err)
		return
	}
	log.Printf("rows affected when creating db %d\n", no)
}
