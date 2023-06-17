package main

import (
	"database/sql"
    _ "github.com/mattn/go-sqlite3"
	"os"
)

const dbName = "kkm_database.db"

func createDataBase() bool {
	_, err := os.Open(dbName)
	if os.IsNotExist(err) {
		file, err := os.Create(dbName)
		if err != nil {
			panic(err)
		}
		file.Close()
		return false
	}
	return true
}


func conectionToDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		panic(err)
	}
	return db
}


func createTable() {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		panic(err)
	}

	crearTabla := `CREATE TABLE IF NOT EXISTS COORD (tipo VARCHAR(20) NOT NULL ,key VARCHAR(2) NOT NULL ,value VARCHAR(2) NOT NULL, PRIMARY KEY (tipo, key))`
	_, err = db.Exec(crearTabla)
	if err != nil {
		panic(err)
	}

	defer db.Close()
}


func insertCoordData(db *sql.DB, tipo string, key string , value string) bool {
	stmt, err := db.Prepare("INSERT INTO coord(tipo, key, value) VALUES(?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(tipo, key, value)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	return true
}


func getCoordData(db *sql.DB, tipo string, key string) string {
	var value string
	_ = db.QueryRow("SELECT VALUE FROM COORD WHERE TIPO = ? AND KEY = ?", tipo, key).Scan(&value); 
	return value
}