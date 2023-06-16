package main

import (
	"fmt"
	//"golang.org/x/crypto/ssh/terminal"
	"os"
	//"golang.org/x/crypto/bcrypt"
)

/*
func login() {
	

	//var storedPassword2 []byte = encryptPassword("mundo")
	var storedPassword2 []byte = []byte("$2a$10$6OLKtiwKQdk69sHoaRHRJOSZKXSSJlUClSfUDz4j7MdtrBPF3WADi")

	fmt.Print("Ingrese su contraseña: ")

	bytePassword, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Error al leer la contraseña:", err)
		return
	}
	fmt.Println(bytePassword)

	err = bcrypt.CompareHashAndPassword(storedPassword2, []byte(bytePassword))
	if err == nil {
		fmt.Println("Contraseña válida")
	} else {
		fmt.Println("Contraseña incorrecta")
	}
}
*/



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
		case len(arguments) == 8: //Insertar data
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
				/*
				- Conectar database
				- Hashear Input
				- Consultar tabla filtrando por Hashs
				- Imprimit Resultados
				*/
				var i int = 3
				for (i < len(arguments)) {
					fmt.Print("Get data - Hashes: ")
					fmt.Println(hash("santander"), hash(arguments[i]))
					getCoordData(db, hash(arguments[2]), hash(arguments[4]))
					i++
				}
				db.Close()
			}    
		default:
			print_help()
	}
}