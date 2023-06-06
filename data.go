// +build ignore

package main

import (
	"database/sql"
    _ "github.com/mattn/go-sqlite3"  // _ para que no tire error por no usar el package
)


const defaultPassword = []byte("$2a$10$6OLKtiwKQdk69sHoaRHRJOSZKXSSJlUClSfUDz4j7MdtrBPF3WADi") //holamundo


func createDataBase() {
	db, err := sql.Open("sqlite3", "/home/manjaro/.kpm/passwd.db")
	if err != nil {
		panic(err)
	}
	//defer db.Close()


	crearTabla := `
    CREATE TABLE IF NOT EXISTS UserData (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user TEXT,
        passwd TEXT,
		email TEXT
    )`
	_, err = db.Exec(crearTabla)
	if err != nil {
		panic(err)
	}


	stmt, err := db.Prepare("INSERT INTO UserData(user, passwd, email) VALUES(?, ?, ?)")
	if err != nil {
		panic(err)
	}
	//defer stmt.Close()


	_, err = stmt.Exec("admin", defaultPassword, "example@example.com")
	if err != nil {
		panic(err)
	}
}


