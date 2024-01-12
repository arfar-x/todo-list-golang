package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"todo_list/cmd"
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

	cmd.Run(os.Args)
}
