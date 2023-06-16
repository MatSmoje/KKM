package main

import (
	_"fmt"
	"crypto/sha1"
    "encoding/hex"
)


func hash(word string) string {
    h := sha1.New()
    h.Write([]byte(word))
    sha1_hash := hex.EncodeToString(h.Sum(nil))
    return string(sha1_hash)
}

/*
func hash(coord string) string {
	var saltedCoord string = "holaMundo"+coord
	//fmt.Println(bcrypt.DefaultCost) //Averiguar
	hashedValue, err := bcrypt.GenerateFromPassword([]byte(saltedCoord), bcrypt.DefaultCost)
	
	if err != nil {
		fmt.Println("Error al generar el hash de valor:", err)
	}
	//fmt.Println(string(hashedValue))
	return string(hashedValue)
}


func EncryptData(data string) string  {
	return "data encrypted"
}


func DecryptData(data string) string  {
	return "data Decrypted"
}


func CompareHash(storedHash string, passwordHash string) bool  {
	err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(passwordHash))
	if err == nil {
		return true
	} 
	return false
}
*/