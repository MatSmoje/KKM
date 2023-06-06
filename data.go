package main

import (
	"database/sql"
    _ "github.com/mattn/go-sqlite3"
	"os"
	"fmt"
	"log"
)


const defaultPassword = "$2a$10$6OLKtiwKQdk69sHoaRHRJOSZKXSSJlUClSfUDz4j7MdtrBPF3WADi" //holamundo
const defaultUser = "JohnDoe" //holamundo
const defaultMail = "john.doe@example.com" //holamundo


func createDataBase() {
	file, err := os.Create("passwd.db")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}


func createTable() *sql.DB {
	db, err := sql.Open("sqlite3", "passwd.db")
	if err != nil {
		panic(err)
	}
	//defer db.Close()
	
	crearTabla := `
    CREATE TABLE IF NOT EXISTS passwd (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user TEXT,
        passwd TEXT,
		mail TEXT
    )`
	_, err = db.Exec(crearTabla)
	if err != nil {
		panic(err)
	}
	return db
}

func insertData(db *sql.DB){
	stmt, err := db.Prepare("INSERT INTO passwd(user, passwd, mail) VALUES(?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(defaultUser, defaultPassword, defaultMail)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

func main(){
	fmt.Println("crea DB")
	createDataBase()
	db := createTable()
	insertData(db)
}

