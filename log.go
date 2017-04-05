package main

import (
	"os"
	"fmt"
	"log"
)

func NewLogger(name string) *log.Logger {
	return log.New(os.Stdout, fmt.Sprintf("%s ", name), log.LstdFlags)
}
