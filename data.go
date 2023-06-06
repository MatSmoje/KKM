package main

import (
	"database/sql"
    _ "github.com/mattn/go-sqlite3"
	"os"
	"fmt"
	"log"
)


//const defaultPassword = []byte("$2a$10$6OLKtiwKQdk69sHoaRHRJOSZKXSSJlUClSfUDz4j7MdtrBPF3WADi") //holamundo


func createDataBase() {
	file, err := os.Create("passwd.db")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}
/*
func createTable(db *sql.DB) {
	

    users_table := `CREATE TABLE users (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        "FirstName" TEXT,
        "LastName" TEXT,
        "Dept" TEXT,
        "Salary" INT);`
        query, err := db.Prepare(users_table)
        if err != nil {
            log.Fatal(err)
        }
        query.Exec()
        fmt.Println("Table created successfully!")
}
*/
func main(){
	fmt.Println("crea DB")
	createDataBase()
	db, err := sql.Open("sqlite3", "passwd.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	
	
	crearTabla := `
    CREATE TABLE IF NOT EXISTS passwd (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user TEXT,
        passwd TEXT,
		mail TEXT
    )
`
_, err = db.Exec(crearTabla)
if err != nil {
    panic(err)
}



/**/

stmt, err := db.Prepare("INSERT INTO passwd(user, passwd, mail) VALUES(?, ?, ?)")
if err != nil {
    panic(err)
}
defer stmt.Close()

_, err = stmt.Exec("John", "$2a$10$6OLKtiwKQdk69sHoaRHRJOSZKXSSJlUClSfUDz4j7MdtrBPF3WADi", "john.doe@example.com")
if err != nil {
    panic(err)
}

}

