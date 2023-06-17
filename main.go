package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args

	if ! createDataBase() {
		createTable()
	}

	switch {
		case len(arguments) == 2:
			if arguments[1] == "-h" || arguments[1] == "--help" {
				print_help()
			} 
			if arguments[1] == "--version"{
				print_version()
			}
		case len(arguments) == 5: //Insertar data
			if (arguments[1] == "-a" || arguments[1] == "-add") {
				db := conectionToDatabase()
				fmt.Println(arguments[2], arguments[3], arguments[4])
				fmt.Println(hash(arguments[2]), hash(arguments[3]), hash(arguments[4]))
				val := insertCoordData(db, hash(arguments[2]), hash(arguments[3]), arguments[4])
				fmt.Println(val)
			}
		case len(arguments) == 4:
			/** Consulta **/
			if (arguments[1] == "-g") {       // kkm -g santander a1 d2 j5
				fmt.Println("Consultando registros..")
				db := conectionToDatabase()
				var i int = 3
				for (i < len(arguments)) {
					fmt.Println(i)
					fmt.Println(arguments[i]+": "+getCoordData(db, hash(arguments[2]), hash(arguments[i])))
					i++
				}
				db.Close()
			}    
		default:
			print_help()
	}
}