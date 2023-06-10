package main

import (
	"database/sql"
    _ "github.com/mattn/go-sqlite3"
	_ "golang.org/x/crypto/bcrypt"
	"os"
	"fmt"
	"log"
)

const defaultPassword = "$2a$10$6OLKtiwKQdk69sHoaRHRJOSZKXSSJlUClSfUDz4j7MdtrBPF3WADi" //holamundo
const defaultUser = "Usuario"
const defaultMail = "user@example.com"

func createDataBase() {
	file, err := os.Create("passwd.db")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}


func conectionToDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "passwd.db")
	if err != nil {
		panic(err)
	}
	return db
}


func createTable() *sql.DB {
	db, err := sql.Open("sqlite3", "passwd.db")
	if err != nil {
		panic(err)
	}
	
	crearTabla := `CREATE TABLE IF NOT EXISTS passwd (id INTEGER PRIMARY KEY AUTOINCREMENT,user TEXT,passwd TEXT,mail TEXT)`
	_, err = db.Exec(crearTabla)
	if err != nil {
		panic(err)
	}

	crearTabla = `CREATE TABLE IF NOT EXISTS coord (id INTEGER PRIMARY KEY AUTOINCREMENT,tipo TEXT,key TEXT,value TEXT)`
	_, err = db.Exec(crearTabla)
	if err != nil {
		panic(err)
	}

	return db
}


func insertDataToUserTable(db *sql.DB){
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


func insertDataToCoord(db *sql.DB, tipo string, key string, value string){
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
}


func getDataFromUserTable2(db *sql.DB, tipo string, key string){
	response, err := db.Query("SELECT KEY, VALUE FROM COORD WHERE TIPO = ? AND KEY = ?", tipo, key)
	if err != nil {
		panic(err)
	}
	for response.Next() {
		var coord string
		var val string
		err = response.Scan(&val, &coord)
		fmt.Println(val, coord)
	}
}


func getDataFromUserTable(){
	fmt.Println("..")
}



/*
func main(){
	fmt.Println("crea DB")
	createDataBase()
	db := createTable()
	*/