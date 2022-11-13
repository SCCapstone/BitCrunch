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

	db.sqldb.SetMaxOpenConns(20)
	db.sqldb.SetMaxIdleConns(20)
	db.sqldb.SetConnMaxLifetime(time.Minute * 5)

	query := "CREATE TABLE IF NOT EXISTS users(userid int primary key auto_increment, username text, password text, email text, admin int)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	_, err := db.sqldb.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating product table", err)
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
	_, err := db.sqldb.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbaseName)
	if err != nil {
		log.Printf("Error %s when creating DB\n", err)
		return
	}

}

func (db *dbase) AddUser(username, password, email string, admin int) error {
	if !db.opened {
		db.Open() // will crash if this fails
	}
	query := "INSERT INTO users(username, password, email) VALUES (?, ?, ?, ?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.sqldb.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, username, password, email, admin)
	if err != nil {
		return err
	}

	return nil
}
