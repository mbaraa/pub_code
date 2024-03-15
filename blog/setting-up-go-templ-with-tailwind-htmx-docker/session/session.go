package session

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"time"
)

var currentPassword string

func init() {
	currentPassword = generatePassword()
}

func PrintPassword() {
	log.Printf("your dashboard password is: %s", currentPassword)
}

func CheckPassword(pw string) bool {
	return currentPassword == pw
}

func generatePassword() string {
	sha256 := sha256.New()
	sha256.Write([]byte(time.Now().String()))
	return hex.EncodeToString(sha256.Sum(nil))[:36]
}
