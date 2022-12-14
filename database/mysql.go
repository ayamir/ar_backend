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
	dbpass    = "ayamir"
	hostname  = "127.0.0.1:3306"
	DBName    = "person"
	TableName = "info"
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

	_, err = db.ExecContext(ctx, "create database if not exists "+DBName)
	if err != nil {
		log.Printf("Error %s when create db %s\n", err, DBName)
		return nil, err
	}

	db.Close()
	db, err = sql.Open("mysql", dsn(DBName))
	if err != nil {
		log.Printf("Error %s when opening DB %s\n", err, DBName)
		return nil, err
	}

	ctx, cancelfunc = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	err = db.PingContext(ctx)
	if err != nil {
		log.Printf("Error %s pinging DB %s\n", err, DBName)
		return nil, err
	}

	return db, nil
}

func CreateTable(db *sql.DB) error {
	query := "create table if not exists `" + TableName + "` (\n" +
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
		log.Printf("Error %s when create table %s\n", err, TableName)
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when getting rows affected\n", err)
		return err
	}
	log.Printf("Rows affected when create table %s: %d", TableName, rows)

	return nil
}
