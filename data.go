package main

import (
	"database/sql"
    _ "github.com/mattn/go-sqlite3"
//	_ "golang.org/x/crypto/bcrypt"
	"os"
	"fmt"
	_"reflect"
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

	crearTabla := `CREATE TABLE IF NOT EXISTS COORD (tipo TEXT NOT NULL ,key TEXT NOT NULL ,value TEXT NOT NULL, PRIMARY KEY (tipo, key))`
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
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(tipo, key, value)
	if err != nil {
		panic(err)
		return false
	}
	defer db.Close()
	return true
}


func getCoordData(db *sql.DB, tipo string, key string) /*bool*/ {
	var ra string
	_ = db.QueryRow("SELECT VALUE FROM COORD WHERE TIPO = ? AND KEY = ?", tipo, key).Scan(&ra);  //WHERE TIPO = ? AND KEY = ?", tipo, key)
	//if err != nil { fmt.Println("Error 405051") }
	fmt.Println(ra)
	/*if response {
		fmt.Println(val)
	}*/
	
	//return true
}