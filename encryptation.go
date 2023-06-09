package main

import (
	"golang.org/x/crypto/bcrypt"
	"fmt"
)


func hashPassword(password string) []byte {
	var saltedPasswd string = "hola"+password
	fmt.Println(bcrypt.DefaultCost) //Averiguar
	hashedValue, err := bcrypt.GenerateFromPassword([]byte(saltedPasswd), bcrypt.DefaultCost)
	
	if err != nil {
		fmt.Println("Error al generar el hash de valor:", err)
	}
	fmt.Println(string(hashedValue))
	return hashedValue
}


func EncryptData(data string) string  {
	return fmt.Println("data encrypted")
}


func DecryptData(data string) string  {
	return fmt.Println("data Decrypted")
}


func CompareHash(storedHash string, passwordHash string) bool  {
	return true
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