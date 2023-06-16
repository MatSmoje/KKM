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
		case len(arguments) >= 3: //Insertar data
			if (arguments[1] == "-a" || arguments[1] == "-add")  && (arguments[2] == "-t") && (arguments[4] == "-k") && (arguments[6] == "-v"){
				db := conectionToDatabase()
				fmt.Println(arguments[3], arguments[5], arguments[7])
				fmt.Println(hash(arguments[3]), hash(arguments[5]), hash(arguments[7]))
				val := insertCoordData(db, hash(arguments[3]), hash(arguments[5]), arguments[7])
				fmt.Println(val)
			}
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