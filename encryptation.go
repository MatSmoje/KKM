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