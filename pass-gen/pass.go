package main

import (
	"math/rand"
	"time"
	"fmt"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialBytes = "!@#$%^&*()_+-=[]{}\\|;':\",.<>/?`~"
	numBytes = "0123456789"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generatePassword(length int, includeNumber bool, includeSpecial bool) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var password []byte
	var charSource string

	if includeNumber {
		charSource += "0123456789"
	}
	if includeSpecial {
		charSource += "!@#$%^&*()_+=-"
	}
	charSource += charset

	for i := 0; i < length; i++ {
		randNum := rand.Intn(len(charSource))
		password = append(password, charSource[randNum])
	}
	return string(password)
}

func main() {
	password := generatePassword(12, true, true)
	fmt.Println("Random password:", password)
}