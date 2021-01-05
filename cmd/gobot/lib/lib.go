package lib

import (
	"log"
	"os"
)

func Getenv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		log.Printf("[ERROR] Missing required environment variable " + name)
	}
	return v
}
