package main

import (
	"flag"
	"fmt"
	"math/rand"
)

const (
	letterBytes  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialBytes = "!@#$%^&*()_+-=[]{}\\|;':\",.<>/?`~"
	numBytes     = "0123456789"
)

func generatePassword(length int, useLetters bool, useSpecial bool, useNum bool) string {
	b := make([]byte, length)

	bytesList := []string{}
	weightList := []int{}

	if useLetters {
		bytesList = append(bytesList, letterBytes)
		weightList = append(weightList, 8)
	}
	if useSpecial {
		bytesList = append(bytesList, specialBytes)
		weightList = append(weightList, 1)
	}
	if useNum {
		bytesList = append(bytesList, numBytes)
		weightList = append(weightList, 3)
	}

	for i := range length {
		keyBytes := randomByWeight(bytesList, weightList)
		b[i] = keyBytes[rand.Intn(len(keyBytes))]
	}

	return string(b)
}

var (
	length     int
	useLetters bool
	useSpecial bool
	useNum     bool
	complexity int
)

func main() {
	flag.IntVar(&length, "length", 12, "Length of password")
	flag.BoolVar(&useLetters, "letters", true, "Include Letters")
	flag.BoolVar(&useSpecial, "special", true, "Include Special Characters.")
	flag.BoolVar(&useNum, "numbers", true, "Include numbers in password")
	flag.IntVar(&complexity, "complexity", 0, "Not Currently Used")
	flag.Parse()

	password := generatePassword(length, useLetters, useSpecial, useNum)
	fmt.Println("Random password:", password)

}
