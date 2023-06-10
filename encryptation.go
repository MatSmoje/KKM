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
	return "data encrypted"
}


func DecryptData(data string) string  {
	return "data Decrypted"
}


func CompareHash(storedHash string, passwordHash string) bool  {
	err = bcrypt.CompareHashAndPassword(storedHash, []byte(passwordHash))
	if err == nil {
		return true
	} 
	return false
}
