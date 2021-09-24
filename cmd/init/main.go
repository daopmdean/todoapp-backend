package main

import (
	"fmt"
	"os"
	"todoapp/internal/adapter/mongodb"
)

func main() {
	args := os.Args
	if args[1] == "mongodb" {
		mongodb.InitMongoDB()
		fmt.Println("Initialized MongoDB successfully")
	} else {
		panic("command not found")
	}
}
