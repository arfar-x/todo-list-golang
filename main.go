package main

import (
	"fmt"
	"os"
	"todo_list/cmd"

	"github.com/joho/godotenv"
)

var projectDir, _ = os.Getwd()

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	err := godotenv.Load(projectDir + "/.env")
	if err != nil {
		panic("Env variables could not be injected.")
	}

	cmd.Run(os.Args[1], os.Args[2:])
}
