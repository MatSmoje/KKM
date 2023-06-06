package main

import (
	"golang.org/x/crypto/bcrypt"
	"fmt"
)


func encryptPassword(password string) []byte {
	var saltedPwd string = "hola"+password
	fmt.Println(bcrypt.DefaultCost)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(saltedPwd), bcrypt.DefaultCost)
	
	if err != nil {
		fmt.Println("Error al generar el hash de contraseña:", err)
	}
	fmt.Println(string(hashedPassword))
	return hashedPassword
}

func encryptCoord(coord string) string {
	var salt string = "1990"+coord
	return salt
}


/*

	// Verificar contraseña
	enteredPassword := "contrasena_secreta" // Contraseña ingresada por el usuario
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(enteredPassword))
	if err == nil {
		fmt.Println("Contraseña válida")
	} else {
		fmt.Println("Contraseña incorrecta")
	}
*/