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

	// Create main database
	ctx, cancelfunc1 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc1()
	_, err := db.sqldb.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS gomap")
	if err != nil {
		log.Panic("Error %s when creating DB\n", err)
	}

	db.sqldb.Close()
	db.sqldb, err = sql.Open("mysql", dsn("gomap"))
	if err != nil {
		log.Panic("Could not open gomap database!")
	}

	// options
	db.sqldb.SetMaxOpenConns(20)
	db.sqldb.SetMaxIdleConns(20)
	db.sqldb.SetConnMaxLifetime(time.Minute * 5)

	// users table creation
	query := "CREATE TABLE IF NOT EXISTS users(username text primary key, password text, email text, admin int)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	_, err = db.sqldb.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating users table", err)
	}

	// floors table creation
	query = "CREATE TABLE IF NOT EXISTS floors(name text primary key, devicelist text)"
	ctx, cancelfunc2 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc2()
	_, err = db.sqldb.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating floors table.", err)
	}

	// devices table creation
	query = "CREATE TABLE IF NOT EXISTS devices(name text primary key"
	ctx, cancelfunc3 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc3()
	_, err = db.sqldb.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating devices table.", err)
	}

	db.opened = true
}

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

// func (db *dbase) CreateOrOpenDB(dbaseName string) {
// 	if !db.opened {
// 		db.Open() // will crash if this fails
// 	}
// 	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancelfunc()
// 	_, err := db.sqldb.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbaseName)
// 	if err != nil {
// 		log.Printf("Error %s when creating DB\n", err)
// 		return
// 	}

// }
