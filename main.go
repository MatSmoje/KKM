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
	//createDataBase()
	//createTable()
	argumentsInterpreter(os.Args)
	fmt.Println("crea DB")
	
	//login()
}