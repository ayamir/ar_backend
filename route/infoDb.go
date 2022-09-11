package route

import (
	"ar_backend/database"
	"database/sql"
	"log"
)

const (
	table = database.DBName + "." + database.TableName
)

func GetInfos(db *sql.DB) []Info {
	query := "select * from " + table
	results, err := db.Query(query)
	if err != nil {
		log.Printf("Error %s when query all data from %s\n", err, table)
		return nil
	}

	infos := []Info{}
	for results.Next() {
		var info Info

		err = results.Scan(&info.Code, &info.Name, &info.Motto, &info.LastUpdated)
		if err != nil {
			log.Printf("Error %s when read all records\n", err)
			return nil
		}
		infos = append(infos, info)
	}
	return infos
}

func GetInfo(db *sql.DB, code string) *Info {
	query := "select * from " + table + " where code=" + code
	results, err := db.Query(query)
	if err != nil {
		log.Printf("Error %s when query data from %s which code = %s\n", err, table, code)
		return nil
	}

	info := &Info{}
	if results.Next() {
		err = results.Scan(&info.Code, &info.Name, &info.Motto, &info.LastUpdated)
		if err != nil {
			log.Printf("Error %s when read record which code = %s\n", err, code)
			return nil
		}
	} else {
		return nil
	}

	return info
}

func InsertInfo(db *sql.DB, info Info) {
	_, err := db.Query("insert into "+table+" (code,name,motto,last_updated) value (?,?,?,now())",
		info.Code, info.Name, info.Motto)
	if err != nil {
		log.Printf("Error %s when insert into table %s which %s\n", err, table, info.toString())
	}
}
