package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

const (
	dbuser    = "ayamir"
	dbpass    = "ayanamirei"
	hostname  = "127.0.0.1:3306"
	dbname    = "person"
	tablename = "info"
)

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", dbuser, dbpass, hostname, dbName)
}

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return nil, err
	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	res, err := db.ExecContext(ctx, "create database if not exists "+dbname)
	if err != nil {
		log.Printf("Error %s when create db %s\n", err, dbname)
		return nil, err
	}

	no, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when fetching rows\n", err)
		return nil, err
	}
	log.Printf("rows affected %d\n", no)

	db.Close()
	db, err = sql.Open("mysql", dsn(dbname))
	if err != nil {
		log.Printf("Error %s when opening DB %s\n", err, dbname)
		return nil, err
	}

	ctx, cancelfunc = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	err = db.PingContext(ctx)
	if err != nil {
		log.Printf("Error %s pinging DB %s\n", err, dbname)
		return nil, err
	}
	log.Printf("Connected to DB %s successfully\n", dbname)

	return db, nil
}

func CreateTable(db *sql.DB) error {
	query := "create table if not exists `" + tablename + "` (\n" +
		"code varchar(10) not null,\n" +
		"name varchar(20) not null,\n" +
		"motto varchar(150) not null,\n" +
		"last_updated datetime not null,\n" +
		"primary key (code)\n" +
		") engine=InnoDB default charset utf8;"

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	res, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when create table %s\n", err, tablename)
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when getting rows affected\n", err)
		return err
	}
	log.Printf("Rows affected when create table %s: %d", tablename, rows)

	return nil
}
