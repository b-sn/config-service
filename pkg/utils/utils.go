package utils

import (
	"crypto/rand"
	"fmt"
	"log"
)

func PrintDebug(variable interface{}, comment string) {
	fmt.Printf("!!! %s => %#v\n", comment, variable)
}

func GenerateRandData(length int) []byte {
	res := make([]byte, length)
	if _, err := rand.Read(res); err != nil {
		log.Fatalf("cannot generate token: %v", err)
	}
	return res
}
